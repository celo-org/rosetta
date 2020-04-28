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

package transaction

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/gomega"
)

func TestStakingOps(t *testing.T) {
	g := NewGomegaWithT(t)

	builder := NewAbiBuilder()

	t.Run("CreateAccount", func(t *testing.T) {
		data, err := builder.GetData(&CreateAccount)
		g.Expect(err).ToNot(HaveOccurred())
		t.Log(data)
	})

	t.Run("LockGold", func(t *testing.T) {
		data, err := builder.GetData(&LockGold)
		g.Expect(err).ToNot(HaveOccurred())
		t.Log(data)
	})

	t.Run("UnlockGold", func(t *testing.T) {
		data, err := builder.GetData(&UnlockGold, big.NewInt(100))
		g.Expect(err).ToNot(HaveOccurred())
		t.Log(data)
	})

	t.Run("RelockGold", func(t *testing.T) {
		data, err := builder.GetData(&RelockGold, big.NewInt(1), big.NewInt(100))
		g.Expect(err).ToNot(HaveOccurred())
		t.Log(data)
	})

	t.Run("WithdrawGold", func(t *testing.T) {
		data, err := builder.GetData(&WithdrawGold, big.NewInt(1))
		g.Expect(err).ToNot(HaveOccurred())
		t.Log(data)
	})

	t.Run("Vote", func(t *testing.T) {
		data, err := builder.GetData(&Vote, common.ZeroAddress, big.NewInt(1), common.ZeroAddress, common.ZeroAddress)
		g.Expect(err).ToNot(HaveOccurred())
		t.Log(data)

		t.Run("TooFewArgs", func(t *testing.T) {
			data, err := builder.GetData(&Vote, common.ZeroAddress, big.NewInt(1), common.ZeroAddress)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("argument count mismatch: 3 for 4"))
			t.Log(data)
		})

		t.Run("TooManyArgs", func(t *testing.T) {
			data, err := builder.GetData(&Vote, common.ZeroAddress, big.NewInt(1), common.ZeroAddress, common.ZeroAddress, common.ZeroAddress)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("argument count mismatch: 5 for 4"))
			t.Log(data)
		})

		t.Run("WrongTypes", func(t *testing.T) {
			data, err := builder.GetData(&Vote, big.NewInt(1), big.NewInt(1), common.ZeroAddress, common.ZeroAddress)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("abi: cannot use ptr as type array as argument"))
			t.Log(data)
		})
	})

	t.Run("ActivateVotes", func(t *testing.T) {
		data, err := builder.GetData(&ActivateVotes, common.ZeroAddress)
		g.Expect(err).ToNot(HaveOccurred())
		t.Log(data)
	})

	t.Run("RevokePendingVotes", func(t *testing.T) {
		data, err := builder.GetData(&RevokePendingVotes, common.ZeroAddress, big.NewInt(1), common.ZeroAddress, common.ZeroAddress, big.NewInt(1))
		g.Expect(err).ToNot(HaveOccurred())
		t.Log(data)
	})

	t.Run("RevokeActiveVotes", func(t *testing.T) {
		data, err := builder.GetData(&RevokeActiveVotes, common.ZeroAddress, big.NewInt(1), common.ZeroAddress, common.ZeroAddress, big.NewInt(1))
		g.Expect(err).ToNot(HaveOccurred())
		t.Log(data)
	})
}
