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

package server

import (
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/celo-org/rosetta/celo/airgap"
	"github.com/celo-org/rosetta/celo/wrapper"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/gomega"
	"golang.org/x/net/context"
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

func TestABIPacking(t *testing.T) {
	type testCase struct {
		method *airgap.CeloMethod
		args   []interface{}
	}
	tests := []testCase{
		{airgap.CreateAccount, nil},
		{airgap.AuthorizeVoteSigner, []interface{}{common.HexToAddress("0x01234"), []byte{1, 2, 3, 4, 5}}},

		{airgap.Vote, []interface{}{common.HexToAddress("0x01234"), big.NewInt(1000)}},
		{airgap.RevokeActiveVotes, []interface{}{common.HexToAddress("0x01234"), common.HexToAddress("0x45"), big.NewInt(1000)}},
		{airgap.RevokePendingVotes, []interface{}{common.HexToAddress("0x01234"), common.HexToAddress("0x45"), big.NewInt(1000)}},
		{airgap.ActivateVotes, []interface{}{common.HexToAddress("0x01234")}},

		{airgap.LockGold, nil},
		{airgap.UnlockGold, []interface{}{big.NewInt(100)}},
		{airgap.RelockGold, []interface{}{big.NewInt(1), big.NewInt(100)}},
		{airgap.WithdrawGold, []interface{}{big.NewInt(1)}},
	}

	srvCtx := &serverContextStub{}

	getCtx := func(method *airgap.CeloMethod) (*abi.ABI, argumentsParser, error) {
		for _, cm := range serverMethods {
			for m, parser := range cm.methods {
				if m == method {
					abi, err := cm.abiFactory()
					if err != nil {
						return nil, nil, err
					}
					return abi, parser, nil
				}
			}
		}
		return nil, nil, fmt.Errorf("method %s not defined in contractMethods", method)
	}

	for _, test := range tests {
		method := test.method
		args := test.args

		t.Run(
			fmt.Sprintf("Not serialized call: %s", test.method),
			func(t *testing.T) {
				RegisterTestingT(t)

				abi, parseArgs, err := getCtx(method)
				Ω(err).ShouldNot(HaveOccurred())

				parsedArgs, err := parseArgs(context.Background(), srvCtx, args)
				Ω(err).ShouldNot(HaveOccurred())

				_, err = abi.Pack(method.Name, parsedArgs...)
				Ω(err).ShouldNot(HaveOccurred())
			},
		)

		t.Run(
			fmt.Sprintf("serialized call: %s", test.method),
			func(t *testing.T) {
				RegisterTestingT(t)

				abi, parseArgs, err := getCtx(method)
				Ω(err).ShouldNot(HaveOccurred())

				serialized, err := json.Marshal(args)
				Ω(err).ShouldNot(HaveOccurred())

				var newArgs []interface{}
				err = json.Unmarshal(serialized, &newArgs)
				Ω(err).ShouldNot(HaveOccurred())

				parsedArgs, err := parseArgs(context.Background(), srvCtx, newArgs)
				Ω(err).ShouldNot(HaveOccurred())

				_, err = abi.Pack(method.Name, parsedArgs...)
				Ω(err).ShouldNot(HaveOccurred())
			},
		)
	}

}

// func TestStakingOps(t *testing.T) {
// 	g := NewGomegaWithT(t)

// 	builder, err := NewAbiBuilder()
// 	if err != nil {
// 		t.Fatal("can't create abiBuilder")
// 	}

// 	t.Run("RelockGold", func(t *testing.T) {
// 		data, err := builder.GetData(RelockGold, big.NewInt(1), big.NewInt(100))
// 		g.Expect(err).ToNot(HaveOccurred())
// 		t.Log(data)
// 	})

// 	t.Run("WithdrawGold", func(t *testing.T) {
// 		data, err := builder.GetData(WithdrawGold, big.NewInt(1))
// 		g.Expect(err).ToNot(HaveOccurred())
// 		t.Log(data)
// 	})

// 	t.Run("Vote", func(t *testing.T) {
// 		data, err := builder.GetData(Vote, common.ZeroAddress, big.NewInt(1), common.ZeroAddress, common.ZeroAddress)
// 		g.Expect(err).ToNot(HaveOccurred())
// 		t.Log(data)

// 		t.Run("TooFewArgs", func(t *testing.T) {
// 			data, err := builder.GetData(Vote, common.ZeroAddress, big.NewInt(1), common.ZeroAddress)
// 			g.Expect(err).To(HaveOccurred())
// 			g.Expect(err.Error()).To(Equal("argument count mismatch: 3 for 4"))
// 			t.Log(data)
// 		})

// 		t.Run("TooManyArgs", func(t *testing.T) {
// 			data, err := builder.GetData(Vote, common.ZeroAddress, big.NewInt(1), common.ZeroAddress, common.ZeroAddress, common.ZeroAddress)
// 			g.Expect(err).To(HaveOccurred())
// 			g.Expect(err.Error()).To(Equal("argument count mismatch: 5 for 4"))
// 			t.Log(data)
// 		})

// 		t.Run("WrongTypes", func(t *testing.T) {
// 			data, err := builder.GetData(Vote, big.NewInt(1), big.NewInt(1), common.ZeroAddress, common.ZeroAddress)
// 			g.Expect(err).To(HaveOccurred())
// 			g.Expect(err.Error()).To(Equal("abi: cannot use ptr as type array as argument"))
// 			t.Log(data)
// 		})
// 	})

// 	t.Run("ActivateVotes", func(t *testing.T) {
// 		data, err := builder.GetData(ActivateVotes, common.ZeroAddress)
// 		g.Expect(err).ToNot(HaveOccurred())
// 		t.Log(data)
// 	})

// 	t.Run("RevokePendingVotes", func(t *testing.T) {
// 		data, err := builder.GetData(RevokePendingVotes, common.ZeroAddress, big.NewInt(1), common.ZeroAddress, common.ZeroAddress, big.NewInt(1))
// 		g.Expect(err).ToNot(HaveOccurred())
// 		t.Log(data)
// 	})

// 	t.Run("RevokeActiveVotes", func(t *testing.T) {
// 		data, err := builder.GetData(RevokeActiveVotes, common.ZeroAddress, big.NewInt(1), common.ZeroAddress, common.ZeroAddress, big.NewInt(1))
// 		g.Expect(err).ToNot(HaveOccurred())
// 		t.Log(data)
// 	})
// }
