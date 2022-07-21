# 版本信息
GO_VERSION:=$(shell go env GOPATH)
GOPATH:=$(shell go env GOPATH)
GIT_BRANCH=$(shell git branch | grep '*' | awk '{print $$2}')
GIT_VERSION=$(shell git describe --tags --always)

.PHONY: info
info:
	@echo "==> GO_VERSION: $(GO_VERSION)"
	@echo "==> GOPATH: $(GOPATH)"
	@echo "==> GIT_BRANCH: $(GIT_BRANCH)"
	@echo "==> GIT_VERSION: $(GIT_VERSION)"

.PHONY: run
run:
	go run ./cmd/main/... -conf=./configs

.PHONY: migrate
migrate:
	go run ./cmd/migration/... -conf=./configs

.PHONY: proto_xxx
proto_xxx:
	go run ./cmd/proto/... -path=./api/xxx

.PHONY: ping
ping:
	curl http://127.0.0.1:8081/api/v1/ping/hello && \
    echo "\n" && \
    curl http://127.0.0.1:8081/api/v1/ping/error && \
    echo "\n"
