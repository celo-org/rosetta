# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: all test lint clean

GO ?= latest
CELO_BLOCKCHAIN_PATH?=../celo-blockchain
CELO_MONOREPO_PATH?=../celo-monorepo

GOLANGCI_exists := $(shell command -v golangci-lint 2> /dev/null)

COMMIT_SHA=$(shell git rev-parse HEAD)

build:
	go build ./...

gen-contracts:
	go run contracts/internal/gen-contracts.go -gcelo $(CELO_BLOCKCHAIN_PATH) -monorepo $(CELO_MONOREPO_PATH)

gen-registry:
	go run registry/internal/gen-registry.go

test: 
	go test ./...

lint: ## Run linters.
ifeq ("$(GOLANGCI_exists)","")
	$(error "No golangci in PATH, consult https://github.com/golangci/golangci-lint#install")
else
	golangci-lint run -c .golangci.yml
endif

clean: 
	go clean -cache

