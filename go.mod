module github.com/celo-org/rosetta

go 1.13

require (
	github.com/celo-org/kliento v0.1.2-0.20200608140637-c5afc8cf0f44
	github.com/coinbase/rosetta-sdk-go v0.5.9
	github.com/ethereum/go-ethereum v1.9.23
	github.com/felixge/httpsnoop v1.0.1
	github.com/google/addlicense v0.0.0-20200622132530-df58acafd6d5 // indirect
	github.com/gorilla/handlers v1.4.2
	github.com/jstemmer/go-junit-report v0.9.1 // indirect
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/k0kubun/pp v3.0.1+incompatible // indirect
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
	github.com/miguelmota/go-ethereum-hdwallet v0.0.0-20200123000308-a60dcd172b4c
	github.com/onsi/gomega v1.9.0
	github.com/segmentio/golines v0.0.0-20200306054842-869934f8da7b // indirect
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.6.3
	golang.org/x/net v0.0.0-20200904194848-62affa334b73
	golang.org/x/sync v0.0.0-20200625203802-6e8e738ad208
)

// DO NOT CHANGE DIRECTORY: Create a symlink so this works
// replace github.com/ethereum/go-ethereum => ../celo-blockchain

// Use this to use external build
replace github.com/ethereum/go-ethereum => github.com/celo-org/celo-blockchain v0.0.0-20200612100840-bf2ba25426f9
