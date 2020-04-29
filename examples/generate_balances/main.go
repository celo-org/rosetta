package main

import (
	"encoding/json"
	"fmt"
	"github.com/coinbase/rosetta-sdk-go/types"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type genesis struct {
	Alloc map[string]genesisAllocation `json:"alloc"`
}
type genesisAllocation struct {
	Balance string `json:"balance"`
}
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
	genesisAllocations := &genesis{}
	if err := json.Unmarshal(body, genesisAllocations); err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v\n", genesisAllocations)
	// Write to file
	balances := []*BootstrapBalance{}
	for k, v := range genesisAllocations.Alloc {
		if v.Balance == "0" {
			continue
		}
		balances = append(balances, &BootstrapBalance{
			Account: &types.AccountIdentifier{
				Address: k,
			},
			Value: v.Balance,
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
