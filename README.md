# Celo Rosetta

A monitoring server for celo-blockchain

## Overview


## Starting rosetta

Rosetta has 2 operations modes:
  * local: Rosetta will launch it's own instance of celo-blockchain 
  * remote: Rosetta will rely on an external instance of celo-blockchain


### Examples of runs:

Run on RC0 using forno node
```bash
rosetta serve remote --datadir ./rosetta --nodeUri https://rc0-forno.celo-testnet.org/ --epoch 17280
```

Run with a local node
```bash
rosetta serve local \
  --genesis ./envs/rc0/genesis.json \
  --geth ../celo-blockchain/build/bin/geth \
  --staticNode "enode://33ac194052ccd10ce54101c8340dbbe7831de02a3e7dcbca7fd35832ff8c53a72fd75e57ce8c8e73a0ace650dc2c2ec1e36f0440e904bc20a3cf5927f2323e85@34.83.199.225:30303" \
  --datadir "./envs/rc0"
```

## Running with the docker image

Docker image is configured by default to:
  * Use `geth` binary inside the image
  * Use `/data` as datadir (should to be mounted)
  * Expects genesis.json to be at `/data/genesis.json`
  * HttpServer listens on port 8080

To run the docker image do:
```bash 
docker run -v "${PWD}/envs/rc0:/data" -p 8080:8080--name rosetta gcr.io/celo-testnet/rosetta:0.1 serve local \
  --staticNode "enode://33ac194052ccd10ce54101c8340dbbe7831de02a3e7dcbca7fd35832ff8c53a72fd75e57ce8c8e73a0ace650dc2c2ec1e36f0440e904bc20a3cf5927f2323e85@34.83.199.225:30303"
```


## Running on development

To run on development, you need to call `go run main.go`

* Run `make rc0-env` to initialize RC0 datadir with the genesis.json on `./envs/rc0`

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
* `make gen-rpc`: Regenerates server scaffold based on `swagger.json` spec
* `make test` or `go test ./...` to run unit tests
* `go build ./...` to build all modules (only compiles, doesn't generate or compile rust library)

### Changing Generated Files

* The `templates` folder contains all mustache templates for the scaffold.
* Do changes to the templates and run `make gen-rpc`
* If you modify a generated file, and want to avoid to overwrite it. Add it to `.openapi-generator-ignore`

### Managing Generated Contracts

Rosetta requires a few Celo Core Contracts

* The list of required contracts is defined on `scripts/gen-contracts.go` file
* Generation requires acces to `celo-blockchain` & `celo-monorepo`.
* Generation assumes both projects are **already properly built**
* To run generator do `make gen-contracts`

## How to build Docker Image

To build the docker image:
```bash
export COMMIT_SHA=$(git rev-parse HEAD)
docker build --build-arg COMMIT_SHA=$COMMIT_SHA -t gcr.io/celo-testnet/rosetta:$COMMIT_SHA .
docker push gcr.io/celo-testnet/rosetta:$COMMIT_SHA
```

