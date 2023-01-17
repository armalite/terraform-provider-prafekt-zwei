MAKEFLAGS += --warn-undefined-variables
SHELL = /bin/bash -o pipefail
.DEFAULT_GOAL := help
.PHONY: help fmt apply state lint fix tidy test testacc vet

GOBIN := $(shell go env GOPATH)/bin

## display help message
help:
	@awk '/^##.*$$/,/^[~\/\.0-9a-zA-Z_-]+:/' $(MAKEFILE_LIST) | awk '!(NR%2){print $$0p}{p=$$0}' | awk 'BEGIN {FS = ":.*?##"}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' | sort

## format
fmt:
	terraform fmt -recursive ./examples/
	go fmt ./...

## finds Go programs that use old APIs and rewrites them to use newer ones.
fix:
	go fix ./...

## run the lint aggreator golangci-lint over the codebase
lint:
	(which golangci-lint || go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.47.2)
	$(GOBIN)/golangci-lint run ./...

## update go.mod to match the source code in the module
tidy:
	go mod tidy

## examines Go source code and reports suspicious constructs
vet:
	go vet ./...

## run tests
test:
	go test -cover ./...

# install into ~/go/bin, needed to generate the docs
install: $(GOBIN)/terraform-provider-prefect

$(GOBIN)/terraform-provider-prefect: internal/*/* api/* api/operations.go
	go install

## make docs
docs: $(GOBIN)/terraform-provider-prefect examples/*
	go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
	touch docs

## apply examples
apply:
	terraform -chdir=examples apply
