# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: evm all test clean


GO ?= latest

GITHUB_ORG?=celo-org
GITHUB_REPO?=rosetta

GOLANGCI_VERSION=1.50.0
GOLANGCI_exists := $(shell command -v golangci-lint 2> /dev/null)
GOLANGCI_v_installed := $(shell echo $(shell golangci-lint --version) | cut -d " " -f 4)

COMMIT_SHA=$(shell git rev-parse HEAD)

LICENCE_SCRIPT=addlicense -c "Celo Org" -l "apache" -v

all: 
	go build ./...

deps:
	go get ./...
	go install github.com/google/addlicense@v1.1.1

test: 
	go test ./...

test-cover:
	go test ./... -covermode=count -coverprofile=coverage.out

fmt: 
	go fmt ./...

install-lint-ci:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v$(GOLANGCI_VERSION)

lint: ## Run linters.
ifeq ("$(GOLANGCI_exists)","")
	$(error "No golangci in PATH, consult https://github.com/golangci/golangci-lint#install")
else ifneq ($(GOLANGCI_v_installed), $(GOLANGCI_VERSION))
	$(error "Installed golangci version $(GOLANGCI_v_installed) \
	 does not match expected version $(GOLANGCI_VERSION)")
else
	golangci-lint run -c .golangci.yml
endif

clean:
	go clean -cache

add-license:
	${LICENCE_SCRIPT} analyzer airgap cmd db examples internal kliento service main.go

check-license:
	${LICENCE_SCRIPT} -check analyzer airgap cmd db examples internal kliento service main.go

.PHONY: gen-contracts lint fmt 
