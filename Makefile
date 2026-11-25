BINARY := quest

.PHONY: build test vet leaks-dir leaks-staged clean

build:
	go build -o bin/$(BINARY) ./cmd/quest

test:
	go test ./...

vet:
	go vet ./...
	go tool staticcheck ./...

# leaks-dir runs gitleaks over the contents of the working directory
leaks-dir:
	go tool gitleaks dir .

# leaks-staged runs gitleaks over the staged changes, good for a pre-commit hook
leaks-staged:
	go tool gitleaks git --pre-commit --staged

clean:
	- rm -rf bin
