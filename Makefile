# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: evm all test clean


GO ?= latest
BLS_RS_PATH ?= external/bls-zexe
CELO_BLOCKCHAIN_PATH?=../celo-blockchain
CELO_MONOREPO_PATH?=../celo-monorepo

CARGO_exists := $(shell command -v cargo 2> /dev/null)
LSB_exists := $(shell command -v lsb_release 2> /dev/null)
GOLANGCI_exists := $(shell command -v golangci-lint 2> /dev/null)

ifdef CARGO_exists
.PHONY: $(BLS_RS_PATH)/target/release/libepoch_snark.a
endif

OS :=
ifeq ("$(LSB_exists)","")
	OS = darwin
else
	OS = linux
endif

all: bls-zexe
	go run cmd/gen-contracts/main.go -gcelo $(CELO_BLOCKCHAIN_PATH) -monorepo $(CELO_MONOREPO_PATH)
	go build ./...

bls-zexe: $(BLS_RS_PATH)/target/release/libepoch_snark.a

$(BLS_RS_PATH)/target/release/libepoch_snark.a:
ifeq ("$(CARGO_exists)","")
	$(error "No cargo in PATH, consult https://github.com/celo-org/celo-monorepo/blob/master/SETUP.md")
else
	cd $(BLS_RS_PATH) && cargo build --release
endif


test: all
	build/env.sh go run build/ci.go test $(TEST_FLAGS)

lint: ## Run linters.
ifeq ("$(GOLANGCI_exists)","")
	$(error "No golangci in PATH, consult https://github.com/golangci/golangci-lint#install")
else
	golangci-lint run -c .golangci.yml
endif

clean-geth:
	go clean -cache
	
clean-bls-zexe:
	rm -rf $(BLS_RS_PATH)/target

clean: clean-geth clean-bls-zexe
