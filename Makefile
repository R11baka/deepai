.PHONY: lint vet test help
.DEFAULT_GOAL := help

help: Makefile
	@echo Choose a command to run
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

## build: main.go in cmd folder
build:
	@go build cmd/main.go

## lint: Golang files
lint:
	@golint ./

## vet: run go vet
vet:
	@go vet ./

## test: run unittests
test: 
	@go test ./test
