module github.com/celo-org/rosetta

go 1.13

require (
	github.com/celo-org/kliento v0.0.0-20200428005346-e515c2c18075
	github.com/coinbase/rosetta-sdk-go v0.1.8
	github.com/ethereum/go-ethereum v1.9.8
	github.com/felixge/httpsnoop v1.0.1
	github.com/google/addlicense v0.0.0-20200422172452-68a83edd47bc // indirect
	github.com/gorilla/handlers v1.4.2
	github.com/jstemmer/go-junit-report v0.9.1 // indirect
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
	github.com/miguelmota/go-ethereum-hdwallet v0.0.0-20200123000308-a60dcd172b4c
	github.com/onsi/gomega v1.9.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.6.3
	golang.org/x/net v0.0.0-20190921015927-1a5e07d1ff72
	golang.org/x/sync v0.0.0-20200317015054-43a5402ce75a
)

replace github.com/celo-org/kliento => ../kliento

// DO NOT CHANGE DIRECTORY: Create a symlink so this works
// replace github.com/ethereum/go-ethereum => ../celo-blockchain

// Use this to use external build
replace github.com/ethereum/go-ethereum => github.com/celo-org/celo-blockchain v0.0.0-20200519153823-adbdc7f8c27e
