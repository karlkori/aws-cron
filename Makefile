BINARY_NAME=aws-cron

.PHONY: help
help: ## Show this help
	@echo 'Usage:'
	@grep -E '^\S+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-40s\033[0m %s\n", $$1, $$2}'

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

vet: ## vet
	go vet

.PHONY: tidy
tidy: ## tidy tidy modfiles and format .go files
	go mod tidy -v
	go fmt ./...

# lint: ## lint
# 	golangci-lint run --enable-all