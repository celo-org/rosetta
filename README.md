# Celo Rosetta

A monitoring server for celo-blockchain

## Overview


## Running the Server

* To run the server, we require to set `ROSETTA_DATADIR` which targets to the datadir that will be used by rosetta and celo-blockchain
* Datadir **must contain** the `genesis.json` for the node.
* Datadir **can contain** a configuration file: `rosetta-cfg.yaml` (other extensions work. Check vyper library)
* Run `make rc0-env` to initialize RC0 datadir with the genesis.json on `./envs/rc0`

To run the server while developing do:

```shell
ROSETTA_DATADIR="envs/rc0" go run main.go
```

## Docker Container

To run the server in a docker container
```
docker build --network=host -t api .
```

Once image is built use
```
docker run --rm -it api 
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

* `make all`: Builds project (generates contract wrappers, compiles go project, compiles bls-zexe)
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

* The list of required contracts is defined on `cmd/gen-contracts/main.go` file
* Generation requires acces to `celo-blockchain` & `celo-monorepo`.
* Generation assumes both projects are **already properly built**
* To run generator do `make gen-contracts`



