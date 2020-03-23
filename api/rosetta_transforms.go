package api

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/p2p"
)

func BlockIdentifierFromHeader(header *types.Header) BlockIdentifier {
	return BlockIdentifier{
		Index: header.Number.Int64(),
		Hash:  header.Hash().Hex(),
	}
}

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
