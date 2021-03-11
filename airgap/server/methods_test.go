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
	"math/big"
	"testing"

	"github.com/celo-org/celo-blockchain/common"
	"github.com/celo-org/rosetta/airgap"
	. "github.com/onsi/gomega"
	"golang.org/x/net/context"
)

func TestMethodArgumentParsing(t *testing.T) {
	RegisterTestingT(t)
	type testCase struct {
		method *airgap.CeloMethod
		args   []interface{}
	}
	tests := []testCase{
		{airgap.CreateAccount, nil},
		{airgap.AuthorizeVoteSigner, []interface{}{common.HexToAddress("0x1234"), []byte{1, 2, 3, 4, 5}}},

		{airgap.Vote, []interface{}{common.HexToAddress("0x1234"), big.NewInt(1000)}},
		{airgap.ActivateVotes, []interface{}{common.HexToAddress("0x1234")}},
		{airgap.RevokeActiveVotes, []interface{}{common.HexToAddress("0x1234"), common.HexToAddress("0x45"), big.NewInt(1000)}},
		{airgap.RevokePendingVotes, []interface{}{common.HexToAddress("0x1234"), common.HexToAddress("0x45"), big.NewInt(1000)}},

		{airgap.LockGold, nil},
		{airgap.UnlockGold, []interface{}{big.NewInt(100)}},
		{airgap.RelockGold, []interface{}{big.NewInt(1), big.NewInt(100)}},
		{airgap.WithdrawGold, []interface{}{big.NewInt(1)}},

		{airgap.StableTokenTransfer, []interface{}{common.HexToAddress("0x1234"), big.NewInt(1)}},
	}

	ctx := context.Background()
	srvCtx := &serverContextStub{}

	contractMethods, err := hydrateMethods(srvCtx, serverTransactionMethodDefinitions)
	Ω(err).ShouldNot(HaveOccurred())

	for _, test := range tests {
		method := contractMethods[test.method]
		args := test.args
		t.Run(test.method.String(), func(t *testing.T) {
			RegisterTestingT(t)

			_, err := method(ctx, args)
			Ω(err).ShouldNot(HaveOccurred())
		})
	}
}
