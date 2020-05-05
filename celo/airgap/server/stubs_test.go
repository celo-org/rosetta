package server

import (
	"context"
	"math/big"

	"github.com/celo-org/rosetta/celo/wrapper"
	"github.com/ethereum/go-ethereum/common"
)

type serverContextStub struct{}

func (sc *serverContextStub) addressFor(ctx context.Context, identifier wrapper.RegistryKey) (common.Address, error) {
	return common.ZeroAddress, nil
}

func (sc *serverContextStub) authorizeMetadata(ctx context.Context, popSignature []byte) (*wrapper.EncodedSignature, error) {
	return &wrapper.EncodedSignature{
		R: [32]byte{},
		S: [32]byte{},
		V: 31,
	}, nil
}

func (sc *serverContextStub) voteMetadata(ctx context.Context, group common.Address, value *big.Int) (*wrapper.AddressLesserGreater, error) {
	return &wrapper.AddressLesserGreater{
		Greater: common.HexToAddress("0x01"),
		Lesser:  common.HexToAddress("0x02"),
	}, nil
}

func (sc *serverContextStub) revokeMetadata(ctx context.Context, account common.Address, group common.Address, value *big.Int) (*wrapper.RevokeMetadata, error) {
	return &wrapper.RevokeMetadata{
		Index: big.NewInt(1),
		AddressLesserGreater: &wrapper.AddressLesserGreater{
			Greater: common.HexToAddress("0x01"),
			Lesser:  common.HexToAddress("0x02"),
		},
		Value: big.NewInt(4),
	}, nil
}
