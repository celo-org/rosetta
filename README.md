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
- CELO transfers (internal & external)
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

Arguments are described in the help output of the CLI.

Every argument can be defined using environment variables using `ROSETTA_` prefix; and replacing `.` for `_`; for example:

```sh
ROSETTA_DATADIR="/my/dir"
ROSETTA_GETH_NETWORK="alfajores"
```

## Running the Rosetta RPC Server

Running the Rosetta RPC Server from scratch will take a very long time to sync and require at least 1.5TB of storage space, since it runs a full archive node in the background. While it may be possible to run the Construction API in the future with a non-archive node, this is still required by the Rosetta spec for the Data API implementation in order to perform balance reconciliation.

### Version 1: Running from `rosetta` source code

You will need the following three repositories cloned locally:

- `rosetta` (this repo)
- [`celo-blockchain`](https://github.com/celo-org/celo-blockchain)

You also need the following dependencies to be met:

- `go >= 1.19`
- `golangci` ([installation instructions](https://golangci-lint.run/usage/install/#local-installation)) (linter dependency for the Makefile)

#### Running Rosetta

Prerequisites:

- Checkout the rosetta version that you want and run `make all`.
- Find the `celo-blockchain` version in the rosetta `go.mod` file. Look for a line containing `github.com/celo-org/celo-blockchain` the version comes after separated by a space.
- Checkout `celo-blockchain` at the version specified in rosetta's `go.mod` and run `make geth`
- Replace `<NETWORK>` with one of `alfajores` (developer testnet), `baklava` (validator testnet) or `mainnet`.
- Replace `<PATH-TO-DATADIR>` below, which is the location for the celo-blockchain data directory (the directory does not need to exist before passing it in).
The data directory is network specific so when switching networks you will also need to change the data directory.

Then run:

```bash
go run main.go run \
  --geth.network <NETWORK> \
  --geth.binary ../celo-blockchain/build/bin/geth \
  --geth.syncmode full \
  --geth.gcmode archive \
  --datadir <PATH-TO-DATADIR>
```

You should start to see continuous output looking something like this:

```sh
INFO [01-28|14:09:03.434] Press CTRL-C to stop the process
--nousb --rpc --rpcaddr 127.0.0.1 --rpcport 8545 --rpcvhosts localhost --syncmode full --gcmode archive --rpcapi eth,net,web3,debug,admin,personal --ipcpath <YourPathToRosetta>/rosetta/envs/alfajores/celo/geth.ipc --light.serve 0 --light.maxpeers 0 --maxpeers 1100 --consoleformat term
INFO [01-28|14:09:05.110] Detected Chain Parameters                chainId=44787 epochSize=17280
INFO [01-28|14:09:05.120] Starting httpServer                      listen_address=:8080
INFO [01-28|14:09:05.120] Resuming operation from last persisted  block srv=celo-monitor block=0
INFO [01-28|14:09:05.121] SubscriptionFetchMode:Start              srv=celo-monitor pipe=header_listener start=1

...

INFO [01-28|14:09:25.731] Stored 1000 blocks                       srv=celo-monitor pipe=persister       block=1000 registryUpdates=0
```

You can stop the service and restart by re-running just the last command above (`go run main.go` ... )

### Version 2: Running Rosetta Docker Image

Prerequisites:

- [Install](https://docs.docker.com/engine/install/) and run `docker` (tested with version `19.03.12`)

Rosetta is released as a docker image: `us.gcr.io/celo-testnet/rosetta`. All versions can be found on the [registry page](https://us.gcr.io/celo-testnet/rosetta). Within the docker image, we pack the `rosetta` binary and also the `geth` binary from `celo-blockchain`. Rosetta will run both.

The command below runs the Celo Rosetta RPC server for `alfajores`:

```bash
export RELEASE="latest"  # or specify a release version
# folder for rosetta to use as data directory (saves rosetta.db & celo-blockchain datadir)
export DATADIR="${PWD}/datadir"
mkdir $DATADIR
docker pull us.gcr.io/celo-testnet/rosetta:$RELEASE
docker run --name rosetta --rm \
  -v "${DATADIR}:/data" \
  -p 8080:8080 \
  us.gcr.io/celo-testnet/rosetta:$RELEASE \
  run \
  --geth.network alfajores \
  --geth.syncmode full \
  --geth.gcmode archive

```

To run this for a different network, change the `geth.network` flag from `alfajores` to `mainnet` or `baklava`.

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

_Note that running these checks is most likely infeasible for most people on
mainnet because you will need at least 1.5 terabytes (as of 2022/11) of space
and several days to be able to sync the chain. Under the hood, the service is
syncing a full archive node, which takes a long time. The construction service
needs to reach the tip before submitting transactions. The data checks will
take a while to complete as well (likely a couple of days on a normal laptop
with the current settings) as they reconcile balances for the entire chain._

- Install the [`rosetta-cli`](https://github.com/coinbase/rosetta-cli) according to the instructions. (Note that on Mac, installing the `rosetta-cli` to `/usr/local/bin` or adding its location to you `$PATH` will allow you to call `rosetta-cli` directly on the command line rather than needing to provide the path to the executable). Current testing has been done with `v0.10.3` of the `rosetta-cli`.
- Run the Rosetta service in the background for the respective network (currently only alfajores for both data and construction checks, the data checks work for all other networks).
- Run the CLI checks for `NETWORK_NAME` as follows:

```sh
# alfajores; specify construction or data
rosetta-cli check:data --configuration-file $PATH_TO_ROSETTA/rosetta-cli-conf/$NETWORK_NAME/cli-config.json
```

[See below](#running-rosetta-with-a-mycelo-testnet) for more details on running the reconciliation checks against a mycelo testnet.

### How to generate `bootstrap_balances.json`

This is only necessary for running the data checks if it has not already been created for the particular network. Here's how to generate this for alfajores (for another network, specify the appropriate genesis block filepath and output path):

```sh
go run examples/generate_balances/main.go \
  $PATH_TO_GENESIS_FILE rosetta-cli-conf/mycelo/bootstrap_balances.json
```

### Running Rosetta with a mycelo testnet

- Set `--geth.genesis` to point to the genesis file for the testnet.
- Set `--geth.networkid` to the network ID (if this is a deployed testnet, this may be different from the `ChainID` in the genesis file). If this value is the same as the `ChainID`, it is not necessary to set this parameter.
- Set the `--monitor.initcontracts` flag (at least on the first run), which fetches necessary state from the genesis block and updates the Rosetta DB accordingly.

To run reconciliation tests on this network, you can use the `cli-config` in `./rosetta-cli-conf/mycelo` as a basis. (Note: the default `cli-config` uses the chain ID that is set when the `loadtest` template is used to set up the mycelo network).

- [Generate `bootstrap_balances.json`](#how-to-generate-bootstrap_balancesjson) within the same folder, using the mycelo network's genesis block.
- In the `cli-config.json`:
  - Set `network` parameter to match the `ChainID` in the genesis file (not the network ID).
  - Point `bootstrap_balances` to the generated `bootstrap_balances.json`.

#### Example

Run the following command to run a local testnet with three validators:
```sh
build/bin/mycelo genesis --buildpath compiled-system-contracts --dev.accounts 2 --newenv tmp/rosetta --mnemonic "miss fire behind decide egg buyer honey seven advance uniform profit renew"
build/bin/mycelo validator-init tmp/rosetta
build/bin/mycelo validator-run tmp/rosetta
```

Get the first validator's enode with the following command:
```sh
build/bin/geth attach tmp/rosetta/validator-00/geth.ipc --exec "admin.nodeInfo"
```

In the `rosetta` repository run the following command to start rosetta:
```sh
go run main.go run \
  --geth.genesis ../celo-blockchain/tmp/rosetta/genesis.json \
  --geth.binary ../celo-blockchain/build/bin/geth \
  --geth.syncmode full \
  --geth.gcmode archive \
  --datadir ./test-mycelo \
  --monitor.initcontracts \
  --geth.port 30308 \
  --geth.rpcport 8553 \
  --geth.bootnodes "<ENODE>"
```

## Releasing rosetta

### Versioning

Rosetta uses [semantic versioning](https://semver.org/) and contains 3 distinct
interfaces that should be taken into account when setting the version. Having
so many interfaces exported from this repo makes it likely that most changes
will constitute a breaking change in at least one interface.

* The rosetta HTTP API exposed by the services in this repo.
* The CLI exposed by the rosetta binary.
* All exported code in the repo, since this repo is imported as a module by Coinbase projects.

Breaking changes constitute any change that could cause a downstream consumer's
code to break. E.g:

* Any change that for a given API request results in different response bytes.
  (This seems quite strict, but in the absence of any framework for classifying
  breaking changes (_that would need to be communicated to and agreed upon by
  downstream consumers_) we have to assume that any change in the bytes of a
  response for a given request could break downstream consumers.)
* Any changes in CLI flags or default values, except if a new flag has been
  added which has no effect unless explicitly set.
* Any changes to exported code in this repo or changes that would change the
  value returned by code exported in this repo.
* Bear in mind that breaking changes in the API of celo-blockchain can imply
  breaking changes for this repo, for example changing the value returned by gas
  estimation, or adding a field to a response. Also note that celo-blockchain
  does not follow semantic versioning, so when updating it you will need to
  investigate all changes included in an update and understand how they affect
  rosetta to see if any are breaking.

Changes that are not breaking and are not bug fixes are considered minor
versions, there will likely be very few of these.

Changes that are not breaking and are bug fixes are considered patch versions,
it seems quite likely bug fixes could also be breaking.

### Making a release

* Ensure all required changes are merged to master.
* Create a PR to update the `MiddlewareVersion` in `./service/rpc/versions.go`
  set the version to the version of the release and merge to master.
* Tag the release in GitHub, which will trigger cloudbuild to automatically build the docker image from tagged commit and tag the image as: `us.gcr.io/celo-testnet/rosetta:<TAG>`, `us.gcr.io/celo-testnet/rosetta:<COMMIT_SHA>`, `us.gcr.io/celo-testnet/rosetta:latest`. (Note: this can be done by creating a new tag while creating a release in GitHub.)
  * If there are any build issues, update the master branch on your local repo and run `./scripts/set_tag_and_build_docker_images.sh`, which will attempt to manually tag the commit and build the docker image, which may be helpful for debugging the build.
* In GitHub create a release for the aforementioned git tag, describe all
  changes and breaking changes and add the tags for the docker images built by
  the script.
