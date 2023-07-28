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

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/celo-org/celo-blockchain/core"
	"github.com/coinbase/rosetta-sdk-go/types"
)

type BootstrapBalance struct {
	Account  *types.AccountIdentifier `json:"account_identifier,omitempty"`
	Currency *types.Currency          `json:"currency,omitempty"`
	Value    string                   `json:"value,omitempty"`
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: generate_balances <genesisURL> <outputFile>")
		os.Exit(1)
	}
	genesisURL := os.Args[1]
	bootstrapBalancesFile := os.Args[2]
	// Get Genesis File
	resp, err := http.Get(genesisURL)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	genesis := &core.Genesis{}
	if err := json.Unmarshal(body, genesis); err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v\n", genesis.Alloc)
	// Write to file
	balances := []*BootstrapBalance{}
	for k, v := range genesis.Alloc {
		if v.Balance.String() == "0" {
			continue
		}
		balances = append(balances, &BootstrapBalance{
			// rosetta CLI expects "0x..." format; case sensitive
			Account: &types.AccountIdentifier{
				Address: k.Hex(),
			},
			Value: v.Balance.String(),
			Currency: &types.Currency{
				Symbol:   "cGLD",
				Decimals: 18,
			},
		})
	}
	file, err := json.MarshalIndent(balances, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(bootstrapBalancesFile, file, os.FileMode(0600)); err != nil {
		log.Fatal(err)
	}
	log.Printf("Bootstrap file contains %d balances\n", len(balances))
}
