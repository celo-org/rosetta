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

```txt
Usage:
  rosetta run [flags]

Flags:
      --datadir string            datadir to use
      --geth.binary string        Path to the celo-blockchain binary
      --geth.bootnodes string     Bootnodes to use (separated by ,)
      --geth.genesis string       path to the genesis.json
      --geth.syncmode string      Geth blockchain sync mode (fast, full, light)
      --geth.gcmode string        Geth garbage collection mode (full, archive)
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

```sh
ROSETTA_DATADIR="/my/dir"
ROSETTA_GETH_GENESIS="/path/to/genesis.json"
```

## Running the Rosetta RPC Server

Running the Rosetta RPC Server from scratch will take some time to sync, since it runs a full archive node in the background. While it may be possible to run the Construction API in the future with a non-archive node, this is still required by the Rosetta spec for the Data API implementation in order to perform balance reconciliation.

### Version 1: Running from `rosetta` source code

You will need the following three repositories cloned locally: `rosetta` (this repo), `celo-monorepo`, `celo-blockchain`. You also need: `go >= 1.14`, `rust >= 1.41.0` (blockchain dependency)

#### Running on Alfajores (Testnet)

Prerequisites:

* Checkout `celo-monorepo` branch `alfajores` and run `yarn && yarn build`
* Checkout `celo-blockchain` tag `v1.1.1` (NOTE: check that this matches the version specified in the `rosetta` `go.mod` file) and `make all`
* Export paths to `celo-monorepo` and `celo-blockchain` as `CELO_MONOREPO_PATH` and `CELO_BLOCKCHAIN_PATH` respectively (can be paths relative to `rosetta` repo)
* Checkout `rosetta` tag `v0.7.5` (or latest released tag) and `make gen-contracts && make all`
* Run `make alfajores-env` to create an empty datadir with the genesis block (only needs to be run the first time, upon initializing the service)

Then run:

```bash
go run main.go run \
  --geth.genesis ./envs/alfajores/genesis.json \
  --geth.binary ../celo-blockchain/build/bin/geth \
  --geth.syncmode full \
  --geth.gcmode archive \
  --geth.staticnodes "enode://05977f6b7d3e16a99d27b714f8a029a006e41ec7732167d373dd920d31f72b3a1776650798d8763560854369d36867e9564dad13b4b60a90c347feeb491d83a9@34.83.42.50:30303" \
  --datadir "./envs/alfajores"
```

You can stop the service and restart by re-running just the last command above (`go run main.go` ... )

#### Running on RC1

This is the same as above with a few differences (generally: specifying `rc1` vs. `alfajores`)

Prerequisites:

* Checkout `celo-monorepo` branch `rc1` instead of `alfajores`, run `yarn && yarn build` as above
* `celo-blockchain`: same as above
* Export paths: same as above
* Checkout `rosetta`: same as above
* Run `make rc1-env` to create an empty datadir with the genesis block

Then run:

```bash
go run main.go run \
  --geth.genesis ./envs/rc1/genesis.json \
  --geth.binary ../celo-blockchain/build/bin/geth \
  --geth.syncmode full \
  --geth.gcmode archive \
  --geth.staticnodes "enode://5e0f4e3aaa096e2a2db76622b335cab4d3224d08d16cb11e8855a3a5f30c19d35d81a74b21271562e459495ab203c2f3a5a5747a83eb53ba046aeeb09aa240ff@34.83.110.24:30303" \
  --datadir "./envs/rc1"
```

### Version 2: Running Rosetta Docker Image

**Note:** This is currently slightly out of date, but we are working on getting this up-to-date and improving our release process.

Rosetta is released as a docker image: `us.gcr.io/celo-testnet/rosetta`. All version can be found on the [registry page](https://us.gcr.io/celo-testnet/rosetta)

Within the docker image, we pack `rosetta` binary and also `geth` binary from celo-blockchain. Rosetta will run both.

To run Rosetta using the docker container, the following options must be configured:

* `genesis.json` for the target network (can be found by `curl 'https://storage.googleapis.com/genesis_blocks/baklava' > genesis.json`)
* `staticNodes` or `bootnodes`.
  * With `staticNodes` Rosetta will directly peer to the list of staticNode provided. This node can be any you have access to. For a public list check `https://storage.cloud.google.com/static_nodes/baklava`

Additionaly, it needs a data directory for the geth datadir & rosetta.db

To run Celo Rosetta:

```bash
export STATICNODE="enode://e99a883d0b7d0bacb84cde98c4729933b49adbc94e718b77fdb31779c7ed9da6c49236330a9ae096f42bcbf6e803394229120201704b7a4a3ae8004993fa0876@34.83.92.243:30303"
export RELEASE="0.7.2"
# folder for rosetta to use as data directory (saves rosetta.db & celo-blockchain datadir)
export DATADIR="${PWD}/datadir"
mkdir $DATADIR
curl 'https://storage.googleapis.com/genesis_blocks/alfajores' > "${DATADIR}/genesis.json"
docker pull us.gcr.io/celo-testnet/rosetta:$RELEASE
docker run --name rosetta --rm \
  -v "${DATADIR}:/data" \
  -p 8080:8080 \
  us.gcr.io/celo-testnet/rosetta:$RELEASE \
  run --geth.staticnodes $STATICNODE
```

## Airgap Client Guide

The Celo Rosetta Airgap module is designed to facilitate signing transactions, parameterized by contemporaenous network metadata, in an offline context.

Examples of this metadata include:

* network wide state like "gas price minimum"
* argument specific state like vote amount "effect on validator priority queue"

```js
AirGapServer {
  ObtainMetadata(TxArgs): TxMetadata
  SubmitTx(Tx): Status
}

AirGapClient {
  ConstructTxFromMetadata(TxMetadata): Tx
  SignTx(Tx, PrivateKey): Tx
}
```

### Custody: Staking and Voting

For a documentation resource, please see the [custody docs](https://docs.celo.org/developer-guide/overview/integrations/custody).

For a code resource, please see the [examples](./examples/airgap/main.go).

## Developer Guide

### Setup

You need:

* go >= 1.14
* rust >= 1.41.0
* openapi-generator To re-generate rpc scaffold ([install link](https://openapi-generator.tech))
* golangci To run linter (check https://github.com/golangci/golangci-lint#install )

`Makefile` requires the following env variables:

* `CELO_BLOCKCHAIN_PATH`
* `CELO_MONOREPO_PATH`

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

## How to run rosetta-cli-checks

Install the [`rosetta-cli`](https://github.com/coinbase/rosetta-cli) according to the instructions. Current testing has been done with `v0.5.16`.

* Run the Rosetta service in the background for the respective network (currently only alfajores)
* Run the CLI checks as follows:

```sh
# alfajores; specify construction or data
rosetta-cli check:construction --configuration-file PATH/TO/rosetta/rosetta-cli-conf/testnet/cli-config.json
```

### How to generate `bootstrap_balances.json`

This is only necessary for running the data checks if it has not already been created for the particular network. Here's how to generate this for alfajores (for another network, specify the appropriate genesis block URL and output path):

```sh
go run examples/generate_balances/main.go \
  https://storage.googleapis.com/genesis_blocks/alfajores \
  rosetta-cli-conf/testnet/bootstrap_balances.json
```
