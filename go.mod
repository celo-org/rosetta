module github.com/celo-org/rosetta

go 1.14

require (
	github.com/celo-org/kliento v0.2.1-0.20210311112647-e151ebba2d7e
	github.com/coinbase/rosetta-sdk-go v0.5.9
	github.com/ethereum/go-ethereum v1.9.23
	github.com/felixge/httpsnoop v1.0.1
	github.com/google/addlicense v0.0.0-20200622132530-df58acafd6d5 // indirect
	github.com/gorilla/handlers v1.4.2
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
	github.com/miguelmota/go-ethereum-hdwallet v0.0.0-20200123000308-a60dcd172b4c
	github.com/onsi/ginkgo v1.10.2 // indirect
	github.com/onsi/gomega v1.9.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.6.3
	golang.org/x/net v0.0.0-20201021035429-f5854403a974
	golang.org/x/sync v0.0.0-20201020160332-67f06af15bc9
)

// DO NOT CHANGE DIRECTORY: Create a symlink so this works
// replace github.com/ethereum/go-ethereum => ../celo-blockchain

// Use this to use external build
replace github.com/ethereum/go-ethereum => github.com/celo-org/celo-blockchain v1.2.2
