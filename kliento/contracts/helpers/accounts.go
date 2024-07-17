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

package helpers

import (
	"fmt"

	"github.com/celo-org/celo-blockchain/crypto"
)

type SignatureValues struct {
	R [32]byte
	S [32]byte
	V uint8
}

func GetSignatureValues(signature []byte) (*SignatureValues, error) {
	if len(signature) != crypto.SignatureLength {
		return nil, fmt.Errorf("Invalid signature")
	}

	var r [32]byte
	copy(r[:], signature[:32])
	var s [32]byte
	copy(s[:], signature[32:64])
	v := signature[64]

	encodedSig := SignatureValues{
		R: r,
		S: s,
		V: v,
	}
	return &encodedSig, nil
}
