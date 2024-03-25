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
	"context"
	"encoding/json"
	"math/big"
	"testing"

	"github.com/celo-org/celo-blockchain/common"
	"github.com/celo-org/celo-blockchain/common/hexutil"
	"github.com/celo-org/rosetta/airgap"
	. "github.com/onsi/gomega"
)

func simulateWire(txArgs *airgap.TxArgs) (*airgap.TxArgs, error) {
	data, err := json.Marshal(txArgs)
	if err != nil {
		return nil, err
	}

	var parsed airgap.TxArgs
	err = json.Unmarshal(data, &parsed)
	if err != nil {
		return nil, err
	}

	if parsed.Method != nil {
		parsed.Method, err = airgap.MethodFromString(parsed.Method.String())
		if err != nil {
			return nil, err
		}
	}

	return &parsed, nil
}

func TestClientServer(t *testing.T) {
	RegisterTestingT(t)
	ctx := context.Background()

	client := airgap.NewClient()
	argBuilder := airgap.NewArgBuilder()
	server, err := NewAirgapServer(big.NewInt(400), &serverContextStub{})
	Ω(err).ShouldNot(HaveOccurred())

	t.Run("Vote()", func(t *testing.T) {
		RegisterTestingT(t)

		txArgs, err := argBuilder.Vote(common.HexToAddress("A"), common.HexToAddress("C"), big.NewInt(4))
		Ω(err).ShouldNot(HaveOccurred())

		txArgs, err = simulateWire(txArgs)
		Ω(err).ShouldNot(HaveOccurred())

		txMetadata, err := server.ObtainMetadata(ctx, txArgs)
		Ω(err).ShouldNot(HaveOccurred())

		tx, err := client.ConstructTxFromMetadata(txMetadata)
		Ω(err).ShouldNot(HaveOccurred())

		_ = tx
	})

	t.Run("ReleaseGoldCreateAccount()", func(t *testing.T) {
		RegisterTestingT(t)

		txArgs, err := argBuilder.ReleaseGoldCreateAccount(common.HexToAddress("A"), common.HexToAddress("B"))
		Ω(err).ShouldNot(HaveOccurred())

		txArgs, err = simulateWire(txArgs)
		Ω(err).ShouldNot(HaveOccurred())

		txMetadata, err := server.ObtainMetadata(ctx, txArgs)
		Ω(err).ShouldNot(HaveOccurred())

		tx, err := client.ConstructTxFromMetadata(txMetadata)
		Ω(err).ShouldNot(HaveOccurred())

		_ = tx
	})

	t.Run("contract call with no arguments", func(t *testing.T) {
		RegisterTestingT(t)

		txArgs := buildContractCallTxArgs(common.HexToAddress("A"), common.HexToAddress("B"), "noArgs()", []interface{}{}, false)
		Ω(err).ShouldNot(HaveOccurred())

		txArgs, err = simulateWire(txArgs)
		Ω(err).ShouldNot(HaveOccurred())

		txMetadata, err := server.ObtainMetadata(ctx, txArgs)
		Ω(err).ShouldNot(HaveOccurred())

		hexEncodedData := hexutil.Encode(txMetadata.Data)
		expectedData := "0x83c962bb" // The data is just the function selector
		Ω(hexEncodedData).Should(Equal(expectedData))

		tx, err := client.ConstructTxFromMetadata(txMetadata)
		Ω(err).ShouldNot(HaveOccurred())

		_ = tx
	})

	t.Run("contract call with unencoded primitive-typed arguments", func(t *testing.T) {
		RegisterTestingT(t)

		methodSig := "withdraw(address,uint256,uint32,bytes)"
		methodArgs := []interface{}{"0xDeadDeAddeAddEAddeadDEaDDEAdDeaDDeAD0000", "23535", "0", "0x"}
		txArgs := buildContractCallTxArgs(common.HexToAddress("A"), common.HexToAddress("B"), methodSig, methodArgs, false)
		Ω(err).ShouldNot(HaveOccurred())

		txArgs, err = simulateWire(txArgs)
		Ω(err).ShouldNot(HaveOccurred())

		txMetadata, err := server.ObtainMetadata(ctx, txArgs)
		Ω(err).ShouldNot(HaveOccurred())

		hexEncodedData := hexutil.Encode(txMetadata.Data)
		expectedData := "0x32b7006d000000000000000000000000deaddeaddeaddeaddeaddeaddeaddeaddead00000000000000000000000000000000000000000000000000000000000000005bef000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000000"
		Ω(hexEncodedData).Should(Equal(expectedData))

		tx, err := client.ConstructTxFromMetadata(txMetadata)
		Ω(err).ShouldNot(HaveOccurred())

		_ = tx
	})

	t.Run("contract call with fixed-length byte argument", func(t *testing.T) {
		RegisterTestingT(t)

		methodSig := "deploy(bytes32,address,address)"
		methodArgs := []interface{}{"0x0000000000000000000000000000000000000000000000000000000000000000", "0xDeadDeAddeAddEAddeadDEaDDEAdDeaDDeAD0000", "0xb0935a466e6Fa8FDa8143C7f4a8c149CA56D06FE"}
		txArgs := buildContractCallTxArgs(common.HexToAddress("A"), common.HexToAddress("B"), methodSig, methodArgs, false)
		Ω(err).ShouldNot(HaveOccurred())

		txArgs, err = simulateWire(txArgs)
		Ω(err).ShouldNot(HaveOccurred())

		txMetadata, err := server.ObtainMetadata(ctx, txArgs)
		Ω(err).ShouldNot(HaveOccurred())

		hexEncodedData := hexutil.Encode(txMetadata.Data)
		expectedData := "0xcf9d137c0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000deaddeaddeaddeaddeaddeaddeaddeaddead0000000000000000000000000000b0935a466e6fa8fda8143c7f4a8c149ca56d06fe"
		Ω(hexEncodedData).Should(Equal(expectedData))

		tx, err := client.ConstructTxFromMetadata(txMetadata)
		Ω(err).ShouldNot(HaveOccurred())

		_ = tx
	})

	t.Run("contract call with already encoded arg data", func(t *testing.T) {
		RegisterTestingT(t)

		methodSig := "deploy(bytes32,address,address)"
		methodArgs := []interface{}{"0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000deaddeaddeaddeaddeaddeaddeaddeaddead0000000000000000000000000000b0935a466e6fa8fda8143c7f4a8c149ca56d06fe"}
		txArgs := buildContractCallTxArgs(common.HexToAddress("A"), common.HexToAddress("B"), methodSig, methodArgs, true)
		Ω(err).ShouldNot(HaveOccurred())

		txArgs, err = simulateWire(txArgs)
		Ω(err).ShouldNot(HaveOccurred())

		txMetadata, err := server.ObtainMetadata(ctx, txArgs)
		Ω(err).ShouldNot(HaveOccurred())

		hexEncodedData := hexutil.Encode(txMetadata.Data)
		expectedData := "0xcf9d137c0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000deaddeaddeaddeaddeaddeaddeaddeaddead0000000000000000000000000000b0935a466e6fa8fda8143c7f4a8c149ca56d06fe"
		Ω(hexEncodedData).Should(Equal(expectedData))

		tx, err := client.ConstructTxFromMetadata(txMetadata)
		Ω(err).ShouldNot(HaveOccurred())

		_ = tx
	})
}

func buildContractCallTxArgs(from common.Address, to common.Address, methodSig string, args []interface{}, argsEncoded bool) *airgap.TxArgs {
	bigIntValue := big.NewInt(0)
	return &airgap.TxArgs{
		From:        from,
		Value:       bigIntValue,
		To:          &to,
		Method:      &airgap.CeloMethod{Name: methodSig},
		Args:        args,
		ArgsEncoded: argsEncoded,
	}
}
