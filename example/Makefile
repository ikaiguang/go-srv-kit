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

# make run app=xxx
# go run ./app/${app}/service/cmd/main/... -conf=./app/${app}/service/configs
.PHONY: run
run:
	go run ./cmd/main/... -conf=./configs

# make migrate app=xxx
# go run ./app/${app}/service/cmd/migration/... -conf=./app/${app}/service/configs
.PHONY: migrate
migrate:
	go run ./cmd/migration/... -conf=./configs

# make proto path=api/xxx
# go run ./cmd/proto/... -path=${path}
.PHONY: proto
proto:
	go run ./cmd/proto/... -path=${path}

.PHONY: ping
ping:
	curl http://127.0.0.1:8081/api/v1/ping/hello && \
    echo "\n" && \
    curl http://127.0.0.1:8081/api/v1/ping/error && \
    echo "\n"
