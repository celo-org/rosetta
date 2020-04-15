package db

import (
	"context"
	"database/sql"
	"math/big"

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

func initDatabase(db *sql.DB) error {
	schema := []string{
		"CREATE table IF NOT EXISTS registry (contract text, fromBlock integer, fromTx integer, address blob)",
		"CREATE table IF NOT EXISTS gasPriceMinimum (fromBlock integer, val blob)",
		"CREATE table IF NOT EXISTS stats (lastBlock integer not null DEFAULT 0)",
	}

	for _, sqlString := range schema {
		if _, err := db.Exec(sqlString); err != nil {
			return err
		}
	}

	// Insert an initial lastBlock if none found
	var count uint
	if err := db.QueryRow("SELECT count(lastBlock) FROM stats").Scan(&count); err != nil {
		return err
	}

	if count == 0 {
		_, err := db.Exec(`INSERT INTO stats (lastBlock) VALUES(?)`, 0)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewSqliteDb(dbpath string) (*rosettaSqlDb, error) {
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		return nil, err
	}

	if err := initDatabase(db); err != nil {
		return nil, err
	}

	updateLastBlockStmt, err := db.Prepare("UPDATE stats SET lastBlock = $1")
	if err != nil {
		return nil, err
	}

	insertGasPriceMinimumStmt, err := db.Prepare("INSERT INTO gasPriceMinimum (fromBlock, val) VALUES (?, ?)")
	if err != nil {
		return nil, err
	}

	insertRegistryAddressStmt, err := db.Prepare("INSERT INTO registry (contract, fromBlock, fromTx, address) VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	getLastBlockStmt, err := db.Prepare("SELECT lastBlock FROM stats")
	if err != nil {
		return nil, err
	}

	getRegistryAddressStmt, err := db.Prepare(`
		SELECT address 
			FROM registry 
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
	var block int64

	if err := cs.getLastBlockStmt.QueryRowContext(ctx).Scan(&block); err != nil {
		if err == sql.ErrNoRows {
			return big.NewInt(block), nil
		}
		return nil, err
	}

	return big.NewInt(block), nil
}

func (cs *rosettaSqlDb) CheckBlockNumber(ctx context.Context, block *big.Int) error {
	lastBlock, err := cs.LastPersistedBlock(ctx)
	if err != nil {
		return err
	}
	if block.Cmp(lastBlock) > 0 {
		return ErrFutureBlock
	}
	return nil
}

// GasPriceMinimumFor reads the gpm that was used FOR the specified block.
// That is, it gets the gpm that was set at the end of the preceding block.
// The gpm value is updated at the end of each block, and applied to txs
// in the following block.
func (cs *rosettaSqlDb) GasPriceMinimumFor(ctx context.Context, block *big.Int) (*big.Int, error) {
	prev := new(big.Int).Sub(block, big.NewInt(1))
	if err := cs.CheckBlockNumber(ctx, prev); err != nil {
		return nil, err
	}

	var gpmBytes []byte

	if err := cs.getGasPriceMinimumStmt.QueryRowContext(ctx, prev.Uint64()).Scan(&gpmBytes); err != nil {
		if err == sql.ErrNoRows {
			return big.NewInt(0), nil
		}
		return nil, err
	}

	return new(big.Int).SetBytes(gpmBytes), nil
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
		_, err = tx.StmtContext(ctx, cs.insertGasPriceMinimumStmt).ExecContext(ctx, changeSet.BlockNumber.Int64(), changeSet.GasPriceMinimum.Bytes())
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
