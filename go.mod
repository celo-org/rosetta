module github.com/celo-org/rosetta

go 1.13

require (
	github.com/ethereum/go-ethereum v1.9.8
	github.com/gorilla/mux v1.7.3
)

replace github.com/celo-org/bls-zexe/go => ./external/bls-zexe/go

replace github.com/ethereum/go-ethereum => ../geth

// Use this to use external build
// replace github.com/ethereum/go-ethereum => github.com/celo-org/celo-blockchain master
