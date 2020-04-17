module github.com/celo-org/rosetta

go 1.13

require (
	github.com/coinbase/rosetta-sdk-go v0.1.4
	github.com/ethereum/go-ethereum v1.9.8
	github.com/felixge/httpsnoop v1.0.1
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/k0kubun/pp v3.0.1+incompatible
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
	github.com/miguelmota/go-ethereum-hdwallet v0.0.0-20200123000308-a60dcd172b4c
	github.com/onsi/ginkgo v1.10.2 // indirect
	github.com/onsi/gomega v1.9.0
	github.com/smartystreets/assertions v1.0.1 // indirect
	github.com/spf13/cobra v0.0.7
	github.com/spf13/viper v1.6.2
)

replace github.com/celo-org/bls-zexe/go => ./external/bls-zexe/go

// DO NOT CHANGE DIRECTORY: Create a symlink so this works
// replace github.com/ethereum/go-ethereum => ../celo-blockchain

// Use this to use external build
replace github.com/ethereum/go-ethereum => github.com/celo-org/celo-blockchain v0.0.0-20200408223849-d021a14a4676
