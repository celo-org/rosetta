package api

import (
	"fmt"
	"math/big"

	"github.com/celo-org/rosetta/celo/client/txpool"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/p2p"
)

func NetworkNameFromId(id *big.Int) string {
	uid := id.Uint64()
	if name, ok := ChainIdToNetwork[uid]; ok {
		return name
	}
	return fmt.Sprintf("unknown %d", uid)
}

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

func HeaderContainsTx(header *ethclient.ExtendedHeader, txHash common.Hash) bool {
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
	return &BlockIdentifier{
		Hash:  header.Hash().Hex(),
		Index: new(big.Int).Sub(header.Number, big.NewInt(1)).Int64(),
	}
}

func MapTxHashesToTransaction(txHashes []common.Hash) []Transaction {
	transactions := make([]Transaction, len(txHashes))
	for i, tx := range txHashes {
		transactions[i] = Transaction{
			TransactionIdentifier: TransactionIdentifier{
				Hash: tx.Hex(),
			},
		}
	}
	return transactions
}
