.PHONY: test build lint vet

test: build lint
	go test -v ./...

build:
	go build -v ./...

lint: vet
	golint ./...
	errcheck -verbose ./...
