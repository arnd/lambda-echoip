
GITHASH         ?= $(shell git rev-parse HEAD)
FUNCTIONNAME    ?= lambda-echoip

help: ## Print this help message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

clean: ## Cleans workdir from generated files
	rm -r main function.zip

main: main.go
	GO_ENABLED=0 GOOS=linux go build -ldflags "-s" main.go

test: ## Run tests
	go test -race -cover -coverprofile=coverage.txt -v ./...

function.zip: main
	zip function.zip main

deploy: function.zip ## Updates the function on aws
	aws lambda update-function-code \
	--function-name $(FUNCTIONNAME) \
 	--zip-file fileb://function.zip

build: function.zip ## Builds the code bundle for deployment

all: build

.PHONY: help test build clean deploy
