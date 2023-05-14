.PHONY: build

build: fmt
	go build -o build/rest.exe cmd/main/main.go

fmt:
	go fmt ./...

.DEFAULT_GOAL := build
