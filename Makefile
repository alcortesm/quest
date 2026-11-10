BINARY := quest

.PHONY: build test vet clean

build:
	go build -o bin/$(BINARY) ./cmd/quest

test:
	go test ./...

vet:
	go vet ./...
	go tool staticcheck ./...

clean:
	- rm -rf bin
