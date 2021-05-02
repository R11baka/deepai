.PHONY: lint vet test help

build: ## build main.go in cmd folder
	@go build cmd/main.go

lint: ## Lint Golang files
	@golint ./

vet: ## Run go vet
	@go vet ./

test: ## Run unittests
	@go test ./test

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
