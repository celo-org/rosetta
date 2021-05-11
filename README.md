# Celo Rosetta

[![CircleCI](https://circleci.com/gh/celo-org/rosetta/tree/master.svg?style=shield)](https://circleci.com/gh/celo-org/rosetta/tree/master)
[![License](https://img.shields.io/github/license/celo-org/rosetta.svg)](https://github.com/celo-org/rosetta/blob/master/LICENSE.txt)

A monitoring server for celo-blockchain

## What is Celo Rosetta?

Celo Rosetta is an RPC server that exposes an API to:

- Query Celo's Blockchain
- Obtain Balance Changing Operations
- Construct Airgapped Transactions

With a special focus on getting balance change operations, Celo Rosetta provides an easy way to obtain changes that are not easily queryable using
the celo-blockchain rpc; such as:

- Gas Fee distribution
- Gold transfers (internal & external). Taking in account Tobin Tax
- Epoch Rewards Distribution
- LockedGold & Election Operations

## RPC endpoints

Rosetta exposes the following endpoints:

- `POST /network/list`: Get List of Available Networks
- `POST /network/status`: Get Network Status
- `POST /network/options`: Get Network Options
- `POST /block`: Get a Block
- `POST /block/transaction`: Get a Block Transaction
- `POST /mempool`: Get All Mempool Transactions
- `POST /mempool/transaction`: Get a Mempool Transaction
- `POST /account/balance`: Get an Account Balance
- `POST /construction/metadata`: Get Transaction Construction Metadata
- `POST /construction/submit`: Submit a Signed Transaction

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

You will need the following three repositories cloned locally:

- `rosetta` (this repo)
- [`celo-blockchain`](https://github.com/celo-org/celo-blockchain)

You also need the following dependencies to be met:

- `go >= 1.14`
- `golangci` ([installation instructions](https://golangci-lint.run/usage/install/#local-installation)) (linter dependency for the Makefile)

#### Running on Alfajores (Testnet)

Prerequisites:

- Checkout `celo-blockchain` tag `v1.3.2` (`git fetch --all && git checkout v1.3.2`) (NOTE: check that this matches the version specified in `rosetta`'s `go.mod` file) and `make geth`
- Checkout `rosetta` tag `v0.8.4` (`git fetch --all && git checkout v0.8.4`) (or latest released tag) and `make all`
- Run `make alfajores-env` to create an empty datadir with the genesis block (only needs to be run the first time, upon initializing the service). The output should look something like this:

  ```sh
  mkdir -p ./envs/alfajores
  curl 'https://storage.googleapis.com/genesis_blocks/alfajores' > ./envs/alfajores/genesis.json
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                  Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:-100 12600  100 12600    0     0  36842      0 --:--:-- --:--:-- --:--:-- 36842
  ```

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

- `celo-blockchain`: same as above
- Export paths: same as above
- Checkout `rosetta`: same as above
- Run `make rc1-env` to create an empty datadir with the genesis block. The output should look something like this:

  ```sh
  mkdir -p ./envs/rc1
  curl 'https://storage.googleapis.com/genesis_blocks/rc1' > ./envs/rc1/genesis.json
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                  Dload  Upload   Total   Spent    Left  Speed
  100 25643  100 25643    0     0  56732      0 --:--:-- --:--:-- --:--:-- 56858
  ```

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

You should start to see continuous output looking similar to this:

```sh
INFO [01-28|14:09:03.434] Press CTRL-C to stop the process
INFO [01-28|14:09:03.440] Running geth init                        service=geth
INFO [01-28|14:09:04.104] Geth Already initialized... skipping init service=geth
geth --networkid 44787 --nousb --rpc --rpcaddr 127.0.0.1 --rpcport 8545 --rpcvhosts localhost --syncmode full --gcmode archive --rpcapi eth,net,web3,debug,admin,personal --ipcpath <YourPathToRosetta>/rosetta/envs/alfajores/celo/geth.ipc --light.serve 0 --light.maxpeers 0 --maxpeers 1100 --consoleformat term
INFO [01-28|14:09:05.110] Detected Chain Parameters                chainId=44787 epochSize=17280
INFO [01-28|14:09:05.120] Starting httpServer                      listen_address=:8080
INFO [01-28|14:09:05.120] Resuming operation from last persisted  block srv=celo-monitor block=0
INFO [01-28|14:09:05.121] SubscriptionFetchMode:Start              srv=celo-monitor pipe=header_listener start=1

...

INFO [01-28|14:09:25.731] Stored 1000 blocks                       srv=celo-monitor pipe=persister       block=1000 registryUpdates=0
```

### Version 2: Running Rosetta Docker Image

Prerequisites:

- [Install](https://docs.docker.com/engine/install/) and run `docker` (tested with version `19.03.12`)

Rosetta is released as a docker image: `us.gcr.io/celo-testnet/rosetta`. All versions can be found on the [registry page](https://us.gcr.io/celo-testnet/rosetta). Within the docker image, we pack the `rosetta` binary and also the `geth` binary from `celo-blockchain`. Rosetta will run both.

The command below runs the Celo Rosetta RPC server for `alfajores`:

```bash
export STATICNODE="enode://e99a883d0b7d0bacb84cde98c4729933b49adbc94e718b77fdb31779c7ed9da6c49236330a9ae096f42bcbf6e803394229120201704b7a4a3ae8004993fa0876@34.83.92.243:30303"
export RELEASE="latest"  # or specify a release version
# folder for rosetta to use as data directory (saves rosetta.db & celo-blockchain datadir)
export DATADIR="${PWD}/datadir"
mkdir $DATADIR
curl 'https://storage.googleapis.com/genesis_blocks/alfajores' > "${DATADIR}/genesis.json"
docker pull us.gcr.io/celo-testnet/rosetta:$RELEASE
docker run --name rosetta --rm \
  -v "${DATADIR}:/data" \
  -p 8080:8080 \
  us.gcr.io/celo-testnet/rosetta:$RELEASE \
  run --geth.staticnodes $STATICNODE \
  --geth.syncmode full \
  --geth.gcmode archive

```

To run this for a different network, replace the genesis block generation and staticnode lines with values specific to the network, as detailed directly below:

- `genesis.json` for the target network (can be found by running the following, selecting one of `alfajores`, `baklava`, `rc1` as `<NETWORK>` in `curl 'https://storage.googleapis.com/genesis_blocks/<NETWORK>' > genesis.json`).
- `staticNodes` or `bootnodes`.
  - With `staticNodes` Rosetta will directly peer to the list of staticNode provided. This node can be any you have access to. For a public list check `https://storage.cloud.google.com/static_nodes/<NETWORK>`

## Airgap Client Guide

The Celo Rosetta Airgap module is designed to facilitate signing transactions, parameterized by contemporaenous network metadata, in an offline context.

Examples of this metadata include:

- network wide state like "gas price minimum"
- argument specific state like vote amount "effect on validator priority queue"

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

In addition to the dependencies listed above under the instructions for running from `rosetta` source code, you also need:

- `openapi-generator` To re-generate rpc scaffold ([install link](https://openapi-generator.tech))

### Build Commands

Important commands:

- `make all`: Builds project (compiles all modules), same as `go build ./...`
- `make test` or `go test ./...` to run unit tests

### Interaction with Celo Core Contracts

Rosetta uses [kliento](https://github.com/celo-org/kliento) to interact with the necessary Celo Core Contracts.

## How to run rosetta-cli-checks

- Install the [`rosetta-cli`](https://github.com/coinbase/rosetta-cli) according to the instructions. (Note that on Mac, installing the `rosetta-cli` to `/usr/local/bin` or adding its location to you `$PATH` will allow you to call `rosetta-cli` directly on the command line rather than needing to provide the path to the executable). Current testing has been done with `v0.5.16` of the `rosetta-cli`.
- Run the Rosetta service in the background for the respective network (currently only alfajores for both Data and Construction checks)
- Run the CLI checks for alfajores as follows:

```sh
# alfajores; specify construction or data
rosetta-cli check:construction --configuration-file PATH/TO/rosetta/rosetta-cli-conf/testnet/cli-config.json
```

_Note that running the checks to completion will take a long time if this is the first time you are running Rosetta locally. Under the hood, the service is syncing a full archive node, which takes time (likely a couple of days on a normal laptop). The construction service needs to reach the tip before submitting transactions. The data checks will take a while to complete as well (likely a couple of days on a normal laptop with the current settings) as they reconcile balances for the entire chain._

### How to generate `bootstrap_balances.json`

This is only necessary for running the data checks if it has not already been created for the particular network. Here's how to generate this for alfajores (for another network, specify the appropriate genesis block URL and output path):

```sh
go run examples/generate_balances/main.go \
  https://storage.googleapis.com/genesis_blocks/alfajores \
  rosetta-cli-conf/testnet/bootstrap_balances.json
```
