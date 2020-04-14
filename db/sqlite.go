package db

import (
	"context"
	"database/sql"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	_ "github.com/mattn/go-sqlite3"
)

type rosettaSqlDb struct {
	db *sql.DB

	getLastBlockStmt          *sql.Stmt
	updateLastBlockStmt       *sql.Stmt
	getRegistryAddressStmt    *sql.Stmt
	getGasPriceMinimumStmt    *sql.Stmt
	insertGasPriceMinimumStmt *sql.Stmt
	insertRegistryAddressStmt *sql.Stmt
}

const (
	setGasPriceMinimumOnStmt   = "INSERT INTO gasPriceMinimum (fromBlock, val) VALUES (?, ?)"
	setRegisteredAddressOnStmt = "INSERT INTO registryAddresses (contract, fromBlock, fromTx, address) VALUES (?, ?, ?, ?)"
)

func initDatabase(db *sql.DB) error {
	if _, err := db.Exec("CREATE table IF NOT EXISTS registryAddresses (contract text, fromBlock integer, fromTx integer, address blob)"); err != nil {
		return err
	}

	if _, err := db.Exec("CREATE table IF NOT EXISTS gasPriceMinimum (fromBlock integer, val integer)"); err != nil {
		return err
	}

	_, err := db.Exec(`CREATE table IF NOT EXISTS stats (lastPersistedBlock integer not null DEFAULT 0)`)
	if err != nil {
		return err
	}

	// Insert an initial lastPersistedBlock if none found
	var count uint
	if err := db.QueryRow("SELECT count(lastPersistedBlock) FROM stats").Scan(&count); err != nil {
		return err
	}

	if count == 0 {
		_, err := db.Exec(`INSERT INTO stats (lastPersistedBlock) VALUES(?)`, 0)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewSqliteDb(dbpath string) (*rosettaSqlDb, error) {
	os.Remove(dbpath)

	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		return nil, err
	}

	if err := initDatabase(db); err != nil {
		return nil, err
	}

	updateLastBlockStmt, err := db.Prepare("UPDATE stats SET lastPersistedBlock = $1")
	if err != nil {
		return nil, err
	}

	insertGasPriceMinimumStmt, err := db.Prepare("INSERT INTO gasPriceMinimum (fromBlock, val) VALUES (?, ?)")
	if err != nil {
		return nil, err
	}

	insertRegistryAddressStmt, err := db.Prepare("INSERT INTO registryAddresses (contract, fromBlock, fromTx, address) VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	getLastBlockStmt, err := db.Prepare("SELECT lastPersistedBlock FROM stats")
	if err != nil {
		return nil, err
	}

	getRegistryAddressStmt, err := db.Prepare(`
		SELECT address 
			FROM registryAddresses 
			WHERE contract == $1 and (fromBlock < $2 OR (fromBlock = $2 AND fromTx <= $3)) 
			ORDER BY fromblock DESC, fromTx DESC 
			LIMIT 1
	`)
	if err != nil {
		return nil, err
	}

	getGasPriceMinimumStmt, err := db.Prepare(`
		SELECT val 
			FROM gasPriceMinimum 
			WHERE fromBlock <= $1 
			ORDER BY fromblock DESC
			LIMIT 1
	`)
	if err != nil {
		return nil, err
	}

	return &rosettaSqlDb{
		db:                        db,
		getLastBlockStmt:          getLastBlockStmt,
		updateLastBlockStmt:       updateLastBlockStmt,
		getRegistryAddressStmt:    getRegistryAddressStmt,
		getGasPriceMinimumStmt:    getGasPriceMinimumStmt,
		insertGasPriceMinimumStmt: insertGasPriceMinimumStmt,
		insertRegistryAddressStmt: insertRegistryAddressStmt,
	}, nil
}

func (cs *rosettaSqlDb) LastPersistedBlock(ctx context.Context) (*big.Int, error) {
	var block int64 // TODO(Alec): Figure out better (safer) way of storing bigints

	if err := cs.getLastBlockStmt.QueryRowContext(ctx).Scan(&block); err != nil {
		if err == sql.ErrNoRows {
			return big.NewInt(block), nil
		}
		return nil, err
	}

	return big.NewInt(block), nil
}

func (cs *rosettaSqlDb) CheckBlockNumber(ctx context.Context, block *big.Int) error {
	lastPersistedBlock, err := cs.LastPersistedBlock(ctx)
	if err != nil {
		return err
	}
	if block.Cmp(lastPersistedBlock) > 0 {
		return ErrFutureBlock
	}
	return nil
}

func (cs *rosettaSqlDb) GasPriceMinimumOn(ctx context.Context, block *big.Int) (*big.Int, error) {
	if err := cs.CheckBlockNumber(ctx, block); err != nil {
		return nil, err
	}

	var gpm int64

	if err := cs.getGasPriceMinimumStmt.QueryRowContext(ctx, block.Uint64()).Scan(&gpm); err != nil {
		if err == sql.ErrNoRows {
			return big.NewInt(0), nil
		}
		return nil, err
	}

	return big.NewInt(gpm), nil
}

func (cs *rosettaSqlDb) RegistryAddressOn(ctx context.Context, block *big.Int, txIndex uint, contractName string) (common.Address, error) {
	if err := cs.CheckBlockNumber(ctx, block); err != nil {
		return common.ZeroAddress, err
	}

	var addr common.Address

	if err := cs.getRegistryAddressStmt.QueryRowContext(ctx, contractName, block.Uint64(), txIndex).Scan(&addr); err != nil {
		if err == sql.ErrNoRows {
			return common.ZeroAddress, ErrContractNotFound
		}
		return common.ZeroAddress, err
	}

	return addr, nil
}

func (cs *rosettaSqlDb) RegistryAddressesOn(ctx context.Context, block *big.Int, txIndex uint, contractNames ...string) (map[string]common.Address, error) {
	if err := cs.CheckBlockNumber(ctx, block); err != nil {
		return nil, err
	}

	addresses := make(map[string]common.Address)
	// TODO(Alec): Could this be done more efficiently?
	for _, name := range contractNames {
		address, err := cs.RegistryAddressOn(ctx, block, txIndex, name)
		if err == ErrContractNotFound {
			continue
		} else if err != nil {
			return nil, err
		}
		addresses[name] = address
	}
	return addresses, nil
}

func (cs *rosettaSqlDb) ApplyChanges(ctx context.Context, changeSet *BlockChangeSet) error {

	tx, err := cs.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.StmtContext(ctx, cs.updateLastBlockStmt).ExecContext(ctx, changeSet.BlockNumber.Int64())
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return rollbackErr
		}
		return err
	}

	if changeSet.GasPriceMinimum != nil {
		_, err = tx.StmtContext(ctx, cs.insertGasPriceMinimumStmt).ExecContext(ctx, changeSet.BlockNumber.Int64(), changeSet.GasPriceMinimum.Int64())
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				return rollbackErr
			}
			return err
		}
	}

	setRegisteredAddressOnStmtPrep := tx.StmtContext(ctx, cs.insertRegistryAddressStmt)

	for _, rc := range changeSet.RegistryChanges {
		if _, err := setRegisteredAddressOnStmtPrep.ExecContext(ctx, rc.Contract, changeSet.BlockNumber.Int64(), int64(rc.TxIndex), rc.NewAddress); err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				return rollbackErr
			}
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
