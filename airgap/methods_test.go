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

package airgap

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/gomega"
)

func TestMethodArgsSerializing(t *testing.T) {
	testCases := []struct {
		method *CeloMethod
		args   []interface{}
	}{
		{
			method: CreateAccount,
			args:   []interface{}{},
		},
		{
			method: AuthorizeVoteSigner,
			args: []interface{}{
				common.HexToAddress("0x1111"),
				[]byte{1, 2, 3, 4, 5, 6},
			},
		},
		{
			method: AuthorizeAttestationSigner,
			args: []interface{}{
				common.HexToAddress("0x1111"),
				[]byte{1, 2, 3, 4, 5, 6},
			},
		},
		{
			method: AuthorizeValidatorSigner,
			args: []interface{}{
				common.HexToAddress("0x1111"),
				[]byte{1, 2, 3, 4, 5, 6},
			},
		},
		{
			method: LockGold,
			args:   []interface{}{},
		},
		{
			method: UnlockGold,
			args:   []interface{}{big.NewInt(10000)},
		},
		{
			method: RelockGold,
			args: []interface{}{
				big.NewInt(1),
				big.NewInt(1000),
			},
		},
		{
			method: WithdrawGold,
			args:   []interface{}{big.NewInt(1)},
		},
		{
			method: Vote,
			args: []interface{}{
				common.HexToAddress("0x1111"),
				big.NewInt(1000),
			},
		},
		{
			method: ActivateVotes,
			args: []interface{}{
				common.HexToAddress("0x1111"),
			},
		},
		{
			method: RevokePendingVotes,
			args: []interface{}{
				common.HexToAddress("0x1111"),
				common.HexToAddress("0x2222"),
				big.NewInt(1000),
			},
		},
		{
			method: RevokeActiveVotes,
			args: []interface{}{
				common.HexToAddress("0x1111"),
				common.HexToAddress("0x2222"),
				big.NewInt(1000),
			},
		},
		{
			method: ReleaseGoldCreateAccount,
			args:   []interface{}{},
		},
		{
			method: ReleaseGoldAuthorizeVoteSigner,
			args: []interface{}{
				common.HexToAddress("0x1111"),
				[]byte{1, 2, 3, 4, 5, 6},
			},
		},
		{
			method: ReleaseGoldAuthorizeAttestationSigner,
			args: []interface{}{
				common.HexToAddress("0x1111"),
				[]byte{1, 2, 3, 4, 5, 6},
			},
		},
		{
			method: ReleaseGoldAuthorizeValidatorSigner,
			args: []interface{}{
				common.HexToAddress("0x1111"),
				[]byte{1, 2, 3, 4, 5, 6},
			},
		},
		{
			method: ReleaseGoldLockGold,
			args:   []interface{}{big.NewInt(10000)},
		},
		{
			method: ReleaseGoldUnlockGold,
			args:   []interface{}{big.NewInt(10000)},
		},
		{
			method: ReleaseGoldRelockGold,
			args: []interface{}{
				big.NewInt(1),
				big.NewInt(1000),
			},
		},
		{
			method: ReleaseGoldWithdrawGold,
			args:   []interface{}{big.NewInt(1)},
		},
		{
			method: RevokePendingVotes,
			args: []interface{}{
				common.HexToAddress("0x1111"),
				common.HexToAddress("0x2222"),
				big.NewInt(1000),
			},
		},
		{
			method: RevokeActiveVotes,
			args: []interface{}{
				common.HexToAddress("0x1111"),
				common.HexToAddress("0x2222"),
				big.NewInt(1000),
			},
		},
	}

	for _, _testCase := range testCases {
		testCase := _testCase
		t.Run(testCase.method.String(), func(t *testing.T) {
			RegisterTestingT(t)

			parsedArgs, err := testCase.method.SerializeArguments(testCase.args...)
			Ω(err).ShouldNot(HaveOccurred())

			desearlizedArgs, err := testCase.method.DeserializeArguments(parsedArgs...)
			Ω(err).ShouldNot(HaveOccurred())

			Ω(desearlizedArgs).Should(Equal(testCase.args))
		})
	}
}
