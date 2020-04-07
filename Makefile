# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: evm all test clean


GO ?= latest
BLS_RS_PATH ?= external/bls-zexe
CELO_BLOCKCHAIN_PATH?=../celo-blockchain
CELO_MONOREPO_PATH?=../celo-monorepo
GITHUB_ORG?=celo-org
GITHUB_REPO?=rosetta

CARGO_exists := $(shell command -v cargo 2> /dev/null)
LSB_exists := $(shell command -v lsb_release 2> /dev/null)
GOLANGCI_exists := $(shell command -v golangci-lint 2> /dev/null)
OPENAPIGEN_exists := $(shell command -v openapi-generator 2> /dev/null)

COMMIT_SHA=$(shell git rev-parse HEAD)

.PHONY:
	gen-rpc 
	ifdef CARGO_exists
		$(BLS_RS_PATH)/target/release/libepoch_snark.a
	endif

OS :=
ifeq ("$(LSB_exists)","")
	OS = darwin
else
	OS = linux
endif

all: bls-zexe
	go build ./...

rosetta:
	go build ./...

bls-zexe: $(BLS_RS_PATH)/target/release/libepoch_snark.a

$(BLS_RS_PATH)/target/release/libepoch_snark.a:
ifeq ("$(CARGO_exists)","")
	$(error "No cargo in PATH, consult https://github.com/celo-org/celo-monorepo/blob/master/SETUP.md")
else
	cd $(BLS_RS_PATH) && cargo build --release
endif

gen-contracts:
	go run scripts/gen-contracts.go -gcelo $(CELO_BLOCKCHAIN_PATH) -monorepo $(CELO_MONOREPO_PATH)

gen-rpc:
ifeq ("$(OPENAPIGEN_exists)","")
	$(error "No openapi-generator in PATH, consult https://github.com/OpenAPITools/openapi-generator#1---installation")
else
	GO_POST_PROCESS_FILE="$(GOPATH)/bin/goimports -w" \
	openapi-generator generate -g go-server \
	--input-spec swagger.json \
	--git-user-id $(GITHUB_ORG) --git-repo-id $(GITHUB_REPO) \
	--package-name "api" --additional-properties "sourceFolder=api" \
	--template-dir ./templates \
	--enable-post-process-file
endif

test: 
	go test ./...

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

rc0-env:
	mkdir -p ./envs/rc0
	curl 'https://storage.googleapis.com/genesis_blocks/rc0' > ./envs/rc0/genesis.json

docker-publish: docker-build
	docker push gcr.io/celo-testnet/rosetta:$(COMMIT_SHA)

docker-build:
	echo "Creating docker image with commit: $(COMMIT_SHA)"
	docker build --build-arg COMMIT_SHA=$(COMMIT_SHA) -t us.gcr.io/celo-testnet/rosetta:$(COMMIT_SHA) .
