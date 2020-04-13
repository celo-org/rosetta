package api

import (
	"math/big"

	"github.com/celo-org/rosetta/celo"
	"github.com/celo-org/rosetta/celo/client/txpool"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/p2p"
)

func PeersFromInfo(peersInfo *[]p2p.PeerInfo) []Peer {
	peers := make([]Peer, len(*peersInfo))
	for i, peerInfo := range *peersInfo {
		peers[i].PeerId = peerInfo.ID
	}
	return peers
}

func BuildErrorResponse(code int32, err error) Error {
	return Error{
		Code:    code,
		Message: err.Error(),
	}
}

func TxIdsFromTxAccountMap(txAccountMap txpool.TxAccountMap) []TransactionIdentifier {
	identifiers := []TransactionIdentifier{}
	for _, txNonceMap := range txAccountMap {
		for _, tx := range txNonceMap {
			identifiers = append(identifiers, TransactionIdentifier{
				Hash: tx.Hash.String(),
			})
		}
	}
	return identifiers
}

func HeaderContainsTx(header *ethclient.HeaderAndTxnHashes, txHash common.Hash) bool {
	for _, tx := range header.Transactions {
		if tx == txHash {
			return true
		}
	}
	return false
}

func FullToPartialBlockIdentifier(blockIdentifer BlockIdentifier) PartialBlockIdentifier {
	return PartialBlockIdentifier{
		Index: &blockIdentifer.Index,
		Hash:  &blockIdentifer.Hash,
	}
}

func HeaderToBlockIdentifier(header *types.Header) *BlockIdentifier {
	return &BlockIdentifier{
		Hash:  header.Hash().Hex(),
		Index: header.Number.Int64(),
	}
}

func HeaderToParentBlockIdentifier(header *types.Header) *BlockIdentifier {
	if header.Number.Cmp(big.NewInt(0)) == 0 {
		return HeaderToBlockIdentifier(header)
	}
	return &BlockIdentifier{
		Hash:  header.ParentHash.Hex(),
		Index: new(big.Int).Sub(header.Number, big.NewInt(1)).Int64(),
	}
}

func MapTxHashesToTransaction(txHashes []common.Hash) []TransactionIdentifier {
	transactions := make([]TransactionIdentifier, len(txHashes))
	for i, tx := range txHashes {
		transactions[i] = TransactionIdentifier{Hash: tx.Hex()}
	}
	return transactions
}

func NewAccountIdentifier(addr common.Address, subAccount *SubAccountIdentifier) AccountIdentifier {
	return AccountIdentifier{
		Address:    addr.Hex(),
		SubAccount: subAccount,
	}
}

func NewSubAccountIdentifier(name string, id string, value string) *SubAccountIdentifier {
	return &SubAccountIdentifier{
		SubAccount: name,
		Metadata:   map[string]interface{}{id: value},
	}
}

func NewOperationIdentifier(index int64) OperationIdentifier {
	return OperationIdentifier{
		Index: index,
	}
}

func NewAmount(value *big.Int, currency Currency) Amount {
	return Amount{
		Value:    value.String(),
		Currency: currency,
	}
}

func NewBalance(account common.Address, subAccount *SubAccountIdentifier, amount Amount) *Balance {
	return &Balance{
		AccountIdentifier: NewAccountIdentifier(account, subAccount),
		Amounts:           []Amount{amount},
	}
}

func NewCeloGoldBalance(account common.Address, value *big.Int, subAccount *SubAccountIdentifier) *Balance {
	return NewBalance(account, subAccount, NewAmount(value, CeloGold))
}

func RewardsToOperations(rewards map[common.Address]*big.Int) []Operation {
	operations := make([]Operation, 0, len(rewards))
	opIndex := int64(0)
	for address, value := range rewards {
		operations = append(operations, Operation{
			OperationIdentifier: NewOperationIdentifier(opIndex),
			Account:             NewAccountIdentifier(address, nil),
			Amount:              NewAmount(value, CeloGold),
			Status:              OperationSuccess.String(),
			Type:                OpKindMint.String(),
		})
		opIndex++
	}
	return operations
}

func GasDetailsToOperations(gasDetails map[common.Address]*big.Int) []Operation {
	var operations []Operation
	opIndex := int64(0)
	for address, value := range gasDetails {
		var relatedOps []OperationIdentifier
		if opIndex > 0 {
			relatedOps = make([]OperationIdentifier, opIndex)
			for i := int64(0); i < opIndex; i++ {
				relatedOps[i] = NewOperationIdentifier(i)
			}
		}

		operations = append(operations, Operation{
			OperationIdentifier: NewOperationIdentifier(opIndex),
			Account:             NewAccountIdentifier(address, nil),
			Amount:              NewAmount(value, CeloGold),
			Status:              OperationSuccess.String(),
			Type:                OpKindFee.String(),
			RelatedOperations:   relatedOps,
		})
		opIndex += 1
	}

	return operations
}

func TransferToOperations(baseIndex int64, transfer *celo.Transfer) []Operation {
	return []Operation{
		{
			OperationIdentifier: NewOperationIdentifier(baseIndex),
			Account:             NewAccountIdentifier(transfer.From.Address, nil),
			Amount:              NewAmount(new(big.Int).Neg(transfer.Value), CeloGold),
			Status:              GetOperationStatus(transfer.Status).String(),
			Type:                OpKindTransfer.String(),
		},
		{
			OperationIdentifier: NewOperationIdentifier(baseIndex + 1),
			Account:             NewAccountIdentifier(transfer.To.Address, nil),
			Amount:              NewAmount(transfer.Value, CeloGold),
			Status:              GetOperationStatus(transfer.Status).String(),
			Type:                OpKindTransfer.String(),
			RelatedOperations:   []OperationIdentifier{{Index: baseIndex}},
		},
	}
}
