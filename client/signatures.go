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

package client

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
)

func Sign(message []byte, privateKey *ecdsa.PrivateKey) (*[]byte, error) {
	digest := crypto.Keccak256(message)
	sig, err := crypto.Sign(digest, privateKey)
	return &sig, err
}

func Verify(message []byte, publicKey *ecdsa.PublicKey, signature *[]byte) bool {
	digest := crypto.Keccak256(message)
	publicKeyBytes := crypto.FromECDSAPub(publicKey)
	return crypto.VerifySignature(publicKeyBytes, digest, *signature)
}
