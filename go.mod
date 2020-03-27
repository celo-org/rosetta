module github.com/celo-org/rosetta

go 1.13

require (
	github.com/Nvveen/Gotty v0.0.0-20120604004816-cd527374f1e5 // indirect
	github.com/davecgh/go-spew v1.1.1
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-metrics v0.0.1 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/ethereum/go-ethereum v1.9.8
	github.com/felixge/httpsnoop v1.0.1
	github.com/gorilla/mux v1.7.3
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/k0kubun/pp v3.0.1+incompatible
	github.com/miguelmota/go-ethereum-hdwallet v0.0.0-20200123000308-a60dcd172b4c
	github.com/onsi/gomega v1.7.1
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/sirupsen/logrus v1.5.0 // indirect
	github.com/spf13/viper v1.6.2
	github.com/stretchr/testify v1.4.0
	golang.org/x/tools v0.0.0-20200327185809-c12078ef0834 // indirect
)

replace github.com/celo-org/bls-zexe/go => ./external/bls-zexe/go

replace github.com/ethereum/go-ethereum => ../celo-blockchain

// Use this to use external build
// replace github.com/ethereum/go-ethereum => github.com/celo-org/celo-blockchain master
