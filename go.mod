module github.com/celo-org/rosetta

go 1.13

require (
	github.com/Nvveen/Gotty v0.0.0-20120604004816-cd527374f1e5 // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-metrics v0.0.1 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/ethereum/go-ethereum v1.9.8
	github.com/gorilla/mux v1.7.3
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/sirupsen/logrus v1.5.0 // indirect
)

replace github.com/celo-org/bls-zexe/go => ./external/bls-zexe/go

replace github.com/ethereum/go-ethereum => ../geth

// Use this to use external build
// replace github.com/ethereum/go-ethereum => github.com/celo-org/celo-blockchain master
