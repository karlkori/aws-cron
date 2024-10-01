BINARY_NAME=aws-cron

BUILD_DATE?="$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')"
GIT_COMMIT?="$(shell git rev-parse HEAD)"
VERSION?="$(shell git describe --tags --abbrev=0 | tr -d '\n')"

LDFLAGS ?= -s -w \
	-X github.com/karlkori/aws-cron/internal/version.buildDate=$(BUILD_DATE) \
	-X github.com/karlkori/aws-cron/internal/version.gitCommit=$(GIT_COMMIT) \
	-X github.com/karlkori/aws-cron/internal/version.gitVersion=$(VERSION)

.PHONY: help build run clean test dep vendor vet tidy

help: ## show this help
	@echo 'usage: make [target] ...'
	@echo ''
	@echo 'targets:'
	@egrep '^(.+)\:\ .*##\ (.+)' ${MAKEFILE_LIST} | sed 's/:.*##/#/' | column -t -c 2 -s '#'

default: help

build-all: ## build all
	goreleaser release --snapshot --clean

build: ## build
	CGO_ENABLED=0 GOARCH=arm64 GOOS=darwin go build -o ${BINARY_NAME} -ldflags "$(LDFLAGS)" *.go

run: build ## build and run binary
	./${BINARY_NAME}

clean: ## cleanup
	go clean
	rm -rf ./dist

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
