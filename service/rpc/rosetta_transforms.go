// Copyright 2020 Celo Org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rpc

import (
	"math/big"

	"github.com/celo-org/kliento/client/txpool"
	"github.com/celo-org/rosetta/analyzer"

	rosettaTypes "github.com/coinbase/rosetta-sdk-go/types"

	"github.com/celo-org/celo-blockchain/common"
	gethTypes "github.com/celo-org/celo-blockchain/core/types"
	"github.com/celo-org/celo-blockchain/ethclient"
	"github.com/celo-org/celo-blockchain/p2p"
)

func PeersFromInfo(peersInfo []p2p.PeerInfo) []*rosettaTypes.Peer {
	peers := make([]*rosettaTypes.Peer, len(peersInfo))
	for i, peerInfo := range peersInfo {
		peers[i] = &rosettaTypes.Peer{PeerID: peerInfo.ID}
	}
	return peers
}

func TxIdsFromTxAccountMap(txAccountMap txpool.TxAccountMap) []*rosettaTypes.TransactionIdentifier {
	identifiers := []*rosettaTypes.TransactionIdentifier{}
	for _, txNonceMap := range txAccountMap {
		for _, tx := range txNonceMap {
			identifiers = append(identifiers, &rosettaTypes.TransactionIdentifier{
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

func FullToPartialBlockIdentifier(blockIdentifer *rosettaTypes.BlockIdentifier) *rosettaTypes.PartialBlockIdentifier {
	return &rosettaTypes.PartialBlockIdentifier{
		Index: &blockIdentifer.Index,
		Hash:  &blockIdentifer.Hash,
	}
}

func HeaderToBlockIdentifier(header *gethTypes.Header) *rosettaTypes.BlockIdentifier {
	return &rosettaTypes.BlockIdentifier{
		Hash:  header.Hash().Hex(),
		Index: header.Number.Int64(),
	}
}

func HeaderToParentBlockIdentifier(header *gethTypes.Header) *rosettaTypes.BlockIdentifier {
	if header.Number.Cmp(big.NewInt(0)) == 0 {
		return HeaderToBlockIdentifier(header)
	}
	return &rosettaTypes.BlockIdentifier{
		Hash:  header.ParentHash.Hex(),
		Index: new(big.Int).Sub(header.Number, big.NewInt(1)).Int64(),
	}
}

func MapTxHashesToTransaction(txHashes []common.Hash) []*rosettaTypes.TransactionIdentifier {
	transactions := make([]*rosettaTypes.TransactionIdentifier, len(txHashes))
	for i, tx := range txHashes {
		transactions[i] = &rosettaTypes.TransactionIdentifier{Hash: tx.Hex()}
	}
	return transactions
}

func NewAccountIdentifier(addr common.Address, subAccount *rosettaTypes.SubAccountIdentifier) rosettaTypes.AccountIdentifier {
	return rosettaTypes.AccountIdentifier{
		Address:    addr.Hex(),
		SubAccount: subAccount,
	}
}

func NewOperationIdentifier(index int64) *rosettaTypes.OperationIdentifier {
	return &rosettaTypes.OperationIdentifier{
		Index: index,
	}
}

func NewAmount(value *big.Int, currency *rosettaTypes.Currency) *rosettaTypes.Amount {
	if value == nil {
		return nil
	}
	return &rosettaTypes.Amount{
		Value:    value.String(),
		Currency: currency,
	}
}

func AccountFromAnalyzer(acc analyzer.Account) *rosettaTypes.AccountIdentifier {
	if acc.SubAccount.Identifier == analyzer.AccMain {
		return &rosettaTypes.AccountIdentifier{
			Address: acc.Address.Hex(),
		}
	}

	return &rosettaTypes.AccountIdentifier{
		Address: acc.Address.Hex(),
		SubAccount: &rosettaTypes.SubAccountIdentifier{
			Address:  string(acc.SubAccount.Identifier),
			Metadata: acc.SubAccount.Metadata,
		},
	}
}

func OperationsFromAnalyzer(iop *analyzer.Operation, baseIndex int64) []*rosettaTypes.Operation {
	opIndex := baseIndex
	operations := make([]*rosettaTypes.Operation, len(iop.Changes))
	for i, change := range iop.Changes {
		var relatedOps []*rosettaTypes.OperationIdentifier
		for i := baseIndex; i < opIndex; i++ {
			relatedOps = append(relatedOps, NewOperationIdentifier(i))
		}

		operations[i] = &rosettaTypes.Operation{
			OperationIdentifier: NewOperationIdentifier(opIndex),
			Account:             AccountFromAnalyzer(change.Account),
			Amount:              NewAmount(change.Amount, CeloGold),
			Status:              GetOperationStatus(iop.Successful).String(),
			Type:                string(iop.Type),
			RelatedOperations:   relatedOps,
		}
		opIndex++
	}
	return operations
}
