# Celo Rosetta

A monitoring server for celo-blockchain

## Overview


## Starting rosetta

To run rosetta do:
```
    rosetta run [options]
```

Where rosetta is the binary. 

  * If on development you can replace rosetta by `go run main.go`

### Example for RC0:

Prerequisites:
  * Download `celo-monorepo` branch `rc0` and `yarn && yarn build`
  * Download `celo-blockchain` branch `mc/rosetta-rc0` and `make all`
  * Download `rosetta` branch `rc0` update go.mod and `make gen-contracts && make all`
  * Run `make rc0-env` to create an empty datadir with the genesis block

```bash
rosetta run \
  --genesis ./envs/rc0/genesis.json \
  --geth ../celo-blockchain/build/bin/geth \
  --staticNode "enode://33ac194052ccd10ce54101c8340dbbe7831de02a3e7dcbca7fd35832ff8c53a72fd75e57ce8c8e73a0ace650dc2c2ec1e36f0440e904bc20a3cf5927f2323e85@34.83.199.225:30303" \
  --datadir "./envs/rc0"
```

### Example for Alfajores:

Prerequisites:
  * Download `celo-monorepo` branch `alfajores` and `yarn && yarn build`
  * Download `celo-blockchain` branch `alfajores-tracing-fix` and `make all`
  * Download `rosetta` branch `master` update go.mod and `make gen-contracts && make all`
  * Run `make alfajores-env` to create an empty datadir with the genesis block

```bash
rosetta run \
  --genesis ./envs/alfajores/genesis.json \
  --geth ../celo-blockchain/build/bin/geth \
  --staticNode "enode://05977f6b7d3e16a99d27b714f8a029a006e41ec7732167d373dd920d31f72b3a1776650798d8763560854369d36867e9564dad13b4b60a90c347feeb491d83a9@34.83.42.50:30303"
  --datadir "./envs/alfajores"
```

## Running with the docker image

Docker image is configured by default to:
  * Use `geth` binary inside the image
  * Use `/data` as datadir (should to be mounted)
  * Expects genesis.json to be at `/data/genesis.json`
  * HttpServer listens on port 8080

To run the docker image do:
```bash 
docker run -v "${PWD}/envs/rc0:/data" -p 8080:8080--name rosetta gcr.io/celo-testnet/rosetta:0.1 run \
  --staticNode "enode://33ac194052ccd10ce54101c8340dbbe7831de02a3e7dcbca7fd35832ff8c53a72fd75e57ce8c8e73a0ace650dc2c2ec1e36f0440e904bc20a3cf5927f2323e85@34.83.199.225:30303"
```

## Dev Guide

### Setup

You need:
  * go >= 1.13
  * rust >= 1.41.0
  * openapi-generator To re-generate rpc scaffold ([install link](https://openapi-generator.tech))
  * golangci To run linter (check https://github.com/golangci/golangci-lint#install )

`Makefile` requires the following env variables:
  * `CELO_BLOCKCHAIN_PATH`: By default defines as `../celo-blockchain`
  * `CELO_MONOREPO_PATH`: By default defines as `../celo-monorepo`

`go.mod` is set up to build `celo-blockchain` from `../celo-blockchain`. Which is the default path,
if you need to change it **DON'T COMMIT IT**

### Build Commands

Important commands:

* `make all`: Builds project (compiles go project, compiles bls-zexe)
* `make gen-contracts`: Regenerates contract wrappers
* `make test` or `go test ./...` to run unit tests
* `go build ./...` to build all modules (only compiles, doesn't generate or compile rust library)

### Managing Generated Contracts

Rosetta requires a few Celo Core Contracts

* The list of required contracts is defined on `scripts/gen-contracts.go` file
* Generation requires acces to `celo-blockchain` & `celo-monorepo`.
* Generation assumes both projects are **already properly built**
* To run generator do `make gen-contracts`

## How to build Docker Image

Commands:
  * `make docker-build`
  * `make docker-publish`

## How to run rosetta-validator

```
go get -u github.com/coinbase/rosetta-validator@v0.1.2
go run examples/generate_balances/main.go https://storage.googleapis.com/genesis_blocks/alfajores
mkdir validator-data && cp ./bootstrap_balances.csv ./validator-data
rosetta-validator check:quick
```

