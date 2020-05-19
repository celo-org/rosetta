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

	"github.com/celo-org/rosetta/airgap"
	"github.com/ethereum/go-ethereum/common"
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
}
