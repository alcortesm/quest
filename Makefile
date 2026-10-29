BINARY := quest

.PHONY: build test vet clean

build: test vet
	go build -o bin/$(BINARY) ./cmd

test:
	go test ./...

vet:
	go vet ./...
	staticcheck ./...

clean:
	- rm -rf bin
