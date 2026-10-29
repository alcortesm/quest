BINARY := quest

.PHONY: build test vet clean

build:
	go build -o bin/$(BINARY) ./cmd

test:
	go test ./...

vet:
	go vet ./...
	go tool staticcheck ./...

clean:
	- rm -rf bin
