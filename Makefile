BINARY_NAME=aws-cron

.PHONY: help build run clean test dep vendor vet tidy

help:   ## show this help
	@echo 'usage: make [target] ...'
	@echo ''
	@echo 'targets:'
	@egrep '^(.+)\:\ .*##\ (.+)' ${MAKEFILE_LIST} | sed 's/:.*##/#/' | column -t -c 2 -s '#'

default: help

build-all: ## build all
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows main.go

build: ## build
	GOARCH=arm64 GOOS=darwin go build -o ${BINARY_NAME} main.go

run: build ## build and run binary
	./${BINARY_NAME}

clean: ## cleanup
	go clean
	rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-windows

test: ## run tests
	go test ./...

test_coverage: ## tests coverage
	go test ./... -coverprofile=coverage.out

dep: ## download dependencies
	go mod download

update-dependencies:  ## update golang dependencies
	dep ensure

vendor: ## save all dependencies in the repo
	go mod vendor

vet: ## vet
	go vet

tidy: ## tidy modfiles and format .go files, removes unused dependencies
	go mod tidy -v
	go fmt ./...
