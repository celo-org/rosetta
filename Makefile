# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: evm all test clean


GO ?= latest

GITHUB_ORG?=celo-org
GITHUB_REPO?=rosetta
CELO_BLOCKCHAIN_PATH=../celo-blockchain
CELO_MONOREPO_PATH=../celo-monorepo

GOLANGCI_exists := $(shell command -v golangci-lint 2> /dev/null)

COMMIT_SHA=$(shell git rev-parse HEAD)

LICENCE_SCRIPT=addlicense -c "Celo Org" -l "apache" -v

all: 
	go build ./...

deps:
	go get ./...
	go get github.com/google/addlicense
	# used in CI
	go get github.com/jstemmer/go-junit-report 
	# go get github.com/segmentio/golines
	# go get github.com/mattn/goveralls	

test: 
	go test ./...

test-cover:
	go test ./... -covermode=count

lint: ## Run linters.
ifeq ("$(GOLANGCI_exists)","")
	$(error "No golangci in PATH, consult https://github.com/golangci/golangci-lint#install")
else
	golangci-lint run -c .golangci.yml
endif

clean:
	go clean -cache

rc1-env:
	mkdir -p ./envs/rc1
	curl 'https://storage.googleapis.com/genesis_blocks/rc1' > ./envs/rc1/genesis.json

alfajores-env:
	mkdir -p ./envs/alfajores
	curl 'https://storage.googleapis.com/genesis_blocks/alfajores' > ./envs/alfajores/genesis.json

alfajoresstaging-env:
	mkdir -p ./envs/alfajoresstaging
	curl 'https://storage.googleapis.com/genesis_blocks/alfajoresstaging' > ./envs/alfajoresstaging/genesis.json

rc0-env:
	mkdir -p ./envs/rc0
	curl 'https://storage.googleapis.com/genesis_blocks/rc0' > ./envs/rc0/genesis.json

ci-test:
	mkdir -p /tmp/test-results
	trap "go-junit-report < /tmp/test-results/go-test.out > /tmp/test-results/go-test-report.xml" EXIT
	go test -v ./... | tee /tmp/test-results/go-test.out

add-license:
	${LICENCE_SCRIPT} analyzer airgap cmd db examples internal service main.go


check-license:
	${LICENCE_SCRIPT} -check analyzer airgap cmd db examples internal service main.go

gen-contracts:
	go run ./scripts/gen-contracts.go -gcelo $(CELO_BLOCKCHAIN_PATH) -monorepo $(CELO_MONOREPO_PATH)

.PHONY: gen-contracts