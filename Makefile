PROJECT_ROOT := $(shell pwd)

.PHONY: test cover

test:
	go test -coverprofile c.out ./...

cover: test
	go tool cover -html=c.out

lint: $(GOPATH)/bin/golangci-lint
	@echo "--> Running linter with default config"
	golangci-lint run -c $(PROJECT_ROOT)/.golangcli.yml

$(GOPATH)/bin/golangci-lint:
	@echo "--> Installing linter"
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.16.0