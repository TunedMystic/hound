APP=hound

help:  ## This help
	@echo "Usage:"
	@echo "  make <target>"
	@echo ""
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[1m%-15s\033[0m %s\n", $$1, $$2}'

build: clean  ## Build the binary
	@go build -ldflags="-s -w"

clean:  ## Clean workspace
	@rm -f ${APP}
	@rm -rf tmp
	@rm -rf coverage.txt

dev:  ## Run the program in dev mode.
	@BASE_DIR=$(shell go env GOMOD) DATABASE_NAME=hound.sqlite go run main.go

install:  ## Install project dependencies
	@go mod download

test:  ## Run tests
	@go clean -testcache; BASE_DIR=$(shell go env GOMOD) DATABASE_NAME=hound-test.sqlite go test ./app/... -v -covermode=atomic -coverprofile coverage.txt; go tool cover -func coverage.txt

.PHONY: help build clean install test dev
