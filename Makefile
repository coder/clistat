GO_SRC_FILES := $(shell find . -type f -name '*.go')
SHELL_SRC_FILES := $(shell find . -type f -name '*.sh')

GOLANGCI_LINT_VERSION := v1.64.7

.PHONY: test
test:
	go test -count=1 ./...

.PHONY: test-race
test-race:
	go test -race -count=3 ./...

.PHONY: lint
lint: lint/go lint/shellcheck

.PHONY: lint/go
lint/go: $(GO_SRC_FILES)
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)
	golangci-lint run --timeout=10m

.PHONY: lint/shellcheck
lint/shellcheck: $(SHELL_SRC_FILES)
	shellcheck --external-sources $(SHELL_SRC_FILES)

.PHONY: fmt
fmt: fmt/go

.PHONY: fmt/go
fmt/go: $(shell find . -type f -name '*.go')
	go run mvdan.cc/gofumpt@v0.6.0 -l -w .
