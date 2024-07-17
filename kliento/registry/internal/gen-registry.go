package main

import (
	"bytes"
	"go/format"
	"os"
	"path"
	"text/template"
)

// Map of contract id (ie registry string) => generated contract struct
// This is used by templateStr to generate gen_registry
var contractIDToGeneratedContractType = map[string]string{
	"Accounts":             "Accounts",
	"Attestations":         "Attestations",
	"BlockchainParameters": "BlockchainParameters",
	"DoubleSigningSlasher": "DoubleSigningSlasher",
	"DowntimeSlasher":      "DowntimeSlasher",
	"Election":             "Election",
	"EpochRewards":         "EpochRewards",
	"Escrow":               "Escrow",
	"Exchange":             "Exchange",
	"ExchangeEUR":          "Exchange",
	"ExchangeBRL":          "Exchange",
	"FeeHandler":           "FeeHandler",
	"GasPriceMinimum":      "GasPriceMinimum",
	"GoldToken":            "GoldToken",
	"Governance":           "Governance",
	"GovernanceSlasher":    "GovernanceSlasher",
	"LockedGold":           "LockedGold",
	"Random":               "Random",
	"Reserve":              "Reserve",
	"SortedOracles":        "SortedOracles",
	"StableToken":          "StableToken",
	"StableTokenEUR":       "StableToken",
	"StableTokenBRL":       "StableToken",
	"Validators":           "Validators",
}

var templateStr = `
// Code generated - DO NOT EDIT.

package registry

import (
	"context"
	"fmt"
	"math/big"

	"github.com/celo-org/rosetta/kliento/contracts"
)

{{ range $contractId, $generatedContract := . }}
// {{ $contractId }}ContractID is the registry identifier for '{{ $contractId }}' contract
var {{ $contractId }}ContractID ContractID = "{{ $contractId }}"
{{ end }}

// RegisteredContractIDs are all (known) registered contract identifiers
var RegisteredContractIDs = []ContractID{
	{{ range $contractId, $generatedContract := . }}
	{{ $contractId }}ContractID,
	{{ end }}
}

type boundContracts struct {
	{{ range $contractId, $generatedContract := . }}
	{{ $contractId }}Contract *contracts.{{ $generatedContract }}
	{{ end }}
}

type generatedRegistry interface {
	GetContractByID(ctx context.Context, identifier string, blockNumber *big.Int) (interface{}, error)
	{{ range $contractId, $generatedContract := . }}
	Get{{ $contractId }}Contract(ctx context.Context, blockNumber *big.Int) (*contracts.{{ $generatedContract }}, error)
	{{ end }}
}

func (r *registryImpl) GetContractByID(ctx context.Context, identifier string, blockNumber *big.Int) (interface{}, error) {
	switch identifier {
		{{ range $contractId, $generatedContract := . }}
		case {{ $contractId }}ContractID.String():
			return r.Get{{ $contractId }}Contract(ctx, blockNumber)
		{{ end }}
	}
	return nil, fmt.Errorf("identifier '%s' not found", identifier)
}

{{ range $contractId, $generatedContract := . }}
func (r *registryImpl) Get{{ $contractId }}Contract(ctx context.Context, blockNumber *big.Int) (*contracts.{{ $generatedContract }}, error) {
	identifier := {{ $contractId }}ContractID.String()
	if (r.{{ $contractId }}Contract == nil || r.cache.isDirty(identifier)) {
		address, err := r.GetAddressFor(ctx, blockNumber, {{ $contractId }}ContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.New{{ $generatedContract }}(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.{{ $contractId }}Contract = contract
	}
	return r.{{ $contractId }}Contract, nil
}
{{ end }}
`

func main() {
	template := template.Must(template.New("registryTemplate").Parse(templateStr))

	var buf bytes.Buffer

	template.Execute(&buf, contractIDToGeneratedContractType)
	p, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}

	f, err := os.Create(path.Join("registry/gen_registry.go"))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Write(p)
}
