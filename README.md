# Celo Rosetta

[![CircleCI](https://circleci.com/gh/celo-org/rosetta/tree/master.svg?style=shield)](https://circleci.com/gh/celo-org/rosetta/tree/master)
[![License](https://img.shields.io/github/license/celo-org/rosetta.svg)](https://github.com/celo-org/rosetta/blob/master/LICENSE.txt)

A monitoring server for celo-blockchain

## What is Celo Rosetta?

Celo Rosetta is an RPC server that exposes an API to:
 * Query Celo's Blockchain
 * Obtain Balance Changing Operations
 * Construct Airgapped Transactions

With a special focus on getting balance change operations, Celo Rosetta provides an easy way to obtain changes that are not easily queryable using
the celo-blockchain rpc; such as:

* Gas Fee distribution
* Gold transfers (internal & external). Taking in account Tobin Tax
* Epoch Rewards Distribution
* LockedGold & Election Operations

## RPC endpoints

Rosetta exposes the following endpoints:

 * `POST /network/list`: Get List of Available Networks
 * `POST /network/status`: Get Network Status
 * `POST /network/options`: Get Network Options
 * `POST /block`: Get a Block
 * `POST /block/transaction`: Get a Block Transaction
 * `POST /mempool`: Get All Mempool Transactions
 * `POST /mempool/transaction`: Get a Mempool Transaction
 * `POST /account/balance`: Get an Account Balance
 * `POST /construction/metadata`: Get Transaction Construction Metadata
 * `POST /construction/submit`: Submit a Signed Transaction

For an understanding of inputs & outputs check [servicer.go](./service/rpc/servicer.go)

## Command line arguments

The main command is `rosetta run`, whose arguments are:

```
Usage:
  rosetta run [flags]

Flags:
      --datadir string            datadir to use
      --geth.binary string        Path to the celo-blockchain binary
      --geth.bootnodes string     Bootnodes to use (separated by ,)
      --geth.genesis string       path to the genesis.json
      --geth.ipcpath string       Path to the geth ipc file
      --geth.logfile string       Path to logs file
      --geth.publicip string      Public Ip to configure geth (sometimes required for discovery)
      --geth.staticnodes string   StaticNode to use (separated by ,)
      --geth.verbosity string     Geth log verbosity (number between [1-5])
  -h, --help                      help for run
      --rpc.address string        Listening address for http server
      --rpc.port uint             Listening port for http server (default 8080)
      --rpc.reqTimeout duration   Timeout when serving a request (default 25s)
```

Every argument can be defined using environment variables using `ROSETTA_` prefix; and replacing `.` for `_`; for example:
```
ROSETTA_DATADIR="/my/dir"
ROSETTA_GETH_GENESIS="/path/to/genesis.json"
``` 

## Running Rosetta Docker Image

Rosetta is released as a docker image: `us.gcr.io/celo-testnet/rosetta`. All version can be found on the [registry page](https://us.gcr.io/celo-testnet/rosetta)

Within the docker image, we pack `rosetta` binary and also `geth` binary from celo-blockchain. Rosetta will run both.

To run Rosetta using the docker container, the following options must be configured:
  * `genesis.json` for the target network (can be found by `curl 'https://storage.googleapis.com/genesis_blocks/baklava' > genesis.json`)
  * `staticNodes` or `bootnodes`.
    * With `staticNodes` Rosetta will directly peer to the list of staticNode provided. This node can be any you have access to. For a public list check `https://storage.cloud.google.com/static_nodes/baklava`

Additionaly, it needs a data directory for the geth datadir & rosetta.db

To run Celo Rosetta:

```bash
# Use the last release
export RELEASE="0.5.4" #might be outdated
# folder for rosetta to use as data directory (saves rosetta.db & celo-blockchain datadir)
export DATADIR="/var/rosetta"
docker pull us.gcr.io/celo-testnet/rosetta:$RELEASE
docker run --name rosetta --rm \
  -v "${DATADIR}:/data" \
  -p 8080:8080 \
  -e ROSETTA_GETH_STATICNODES="enode://33ac194052ccd10ce54101c8340dbbe7831de02a3e7dcbca7fd35832ff8c53a72fd75e57ce8c8e73a0ace650dc2c2ec1e36f0440e904bc20a3cf5927f2323e85@34.83.199.225:30303"
  us.gcr.io/celo-testnet/rosetta:$RELEASE \
  run --staticNode $STATICNODE
```

## Developer Guide

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
mkdir validator-data 
go run examples/generate_balances/main.go \
  https://storage.googleapis.com/genesis_blocks/alfajores \
  validator-data/bootstrap_balances.json
rosetta-validator check:complete
```

### Running on development

#### Running on RC1:

Prerequisites:
  * Download `celo-monorepo` branch `rc1` and `yarn && yarn build`
  * Download `celo-blockchain` branch `rc1-tracing-fix` and `make all`
  * Download `rosetta` branch `master` update go.mod and `make gen-contracts && make all`
  * Run `make rc1-env` to create an empty datadir with the genesis block

```bash
go run main.go run \
  --geth.genesis ./envs/rc1/genesis.json \
  --geth.binary ../celo-blockchain/build/bin/geth \
  --geth.staticnodes "enode://5e0f4e3aaa096e2a2db76622b335cab4d3224d08d16cb11e8855a3a5f30c19d35d81a74b21271562e459495ab203c2f3a5a5747a83eb53ba046aeeb09aa240ff@34.83.110.24:30303"
  --datadir "./envs/rc1"
```

#### Running on Alfajores:

Prerequisites:
  * Download `celo-monorepo` branch `alfajores` and `yarn && yarn build`
  * Download `celo-blockchain` branch `alfajores-tracing-fix` and `make all`
  * Download `rosetta` branch `master` update go.mod and `make gen-contracts && make all`
  * Run `make alfajores-env` to create an empty datadir with the genesis block

```bash
go run main.go run \
  --geth.genesis ./envs/alfajores/genesis.json \
  --geth.binary ../celo-blockchain/build/bin/geth \
  --geth.staticnodes "enode://05977f6b7d3e16a99d27b714f8a029a006e41ec7732167d373dd920d31f72b3a1776650798d8763560854369d36867e9564dad13b4b60a90c347feeb491d83a9@34.83.42.50:30303"
  --datadir "./envs/alfajores"
```