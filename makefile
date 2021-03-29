GIT_HASH = $(shell git describe --tags --dirty --always)

.PHONY: all run lint fmt test build clean gen-cert help

# Run help as first command
help:

all: clean test build run ## clean dev env, run all linting and tests, build binary and run tasky

run: ## Run tasky
	./bin/tasky
	
test: lint fmt unit ## Execute all linting and tests 

unit: ## Execute unit tests 
	go test ./... -coverprofile cover.out
	
lint:
	go vet ./...

fmt:
	go fmt ./...

build: test ## Build tasky binary
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on GOSUMDB=off go build -o bin/tasky main.go

docker: ## Build tasky dockerfile
	docker build . -f docker/Dockerfile
	
clean: ## Clean the development environment
	rm -rf cert.pem key.pem bin cover.out *.log

help: ## Show help messages
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'