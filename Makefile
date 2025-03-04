.PHONY: test
test:
	go test -count=1 ./...

.PHONY: test-race
test-race:
	go test -race -count=3 ./...

.PHONY: fmt
fmt: fmt/go

.PHONY: fmt/go
fmt/go: $(shell find . -type f -name '*.go')
	go run mvdan.cc/gofumpt@v0.6.0 -l -w .
