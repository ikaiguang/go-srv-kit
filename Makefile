# 定义环境变量
GOHOSTOS := $(shell go env GOHOSTOS)
GOPATH := $(shell go env GOPATH)
VERSION := $(shell git describe --tags --always)

# 定义项目变量
PROJECT_MAKEFILE := $(abspath $(lastword $(MAKEFILE_LIST)))
PROJECT_ABS_PATH := $(patsubst %/,%,$(dir $(PROJECT_MAKEFILE)))
PROJECT_PATH_NAME := $(notdir $(PROJECT_ABS_PATH))
PROJECT_REL_PATH := "./"

# 示例
ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	GIT_BASH= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	COMMON_PROTO_FILES=$(shell $(GIT_BASH) -c "find $(PROJECT_PATH)api/common -name *.proto")
else
endif

# 定义编译 protobuf
define protoc_protobuf
    if [ "$1" != "" ]; then \
		protoc \
			--proto_path=. \
			--proto_path=$(GOPATH)/src \
			--proto_path=./third_party \
			--go_out=paths=source_relative:. \
			--go-grpc_out=paths=source_relative:. \
			--go-http_out=paths=source_relative:. \
			--go-errors_out=paths=source_relative:. \
			--validate_out=paths=source_relative,lang=go:. \
			--openapiv2_out . \
			--openapiv2_opt logtostderr=true \
			--openapiv2_opt allow_delete_body=true \
			--openapiv2_opt json_names_for_fields=false \
			--openapiv2_opt enums_as_ints=true \
			--openapi_out=fq_schema_naming=true,enum_type=integer,default_response=true:. \
			$1 ; \
    	fi
endef

.DEFAULT_GOAL := help
# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.PHONY: init
# init and install necessary software
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.2
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
	go install github.com/ikaiguang/protoc-gen-go-errors@v0.0.2
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@v0.7.0
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.22.0
	go install github.com/envoyproxy/protoc-gen-validate@v1.1.0
	go install github.com/google/wire/cmd/wire@v0.6.0
	go install github.com/golang/mock/mockgen@v1.6.0
	go install golang.org/x/tools/cmd/goimports@v0.24.0

# ===== include =====
# ===== include =====
# ===== include =====

# api
include api/makefile_protoc.mk
include api/config/makefile_protoc.mk
#include testdata/ping-service/api/makefile_protoc.mk
include testdata/ping-service/api/ping-service/makefile_protoc.mk
include testdata/ping-service/api/ping-service/v1/makefile_protoc.mk
include testdata/ping-service/api/testdata-service/makefile_protoc.mk
include testdata/ping-service/api/testdata-service/v1/makefile_protoc.mk
include testdata/ping-service/internal/conf/makefile_protoc.mk

# run
include testdata/ping-service/cmd/makefile_run.mk

# build
include testdata/ping-service/devops/makefile_cicd.mk

# ===== include =====
# ===== include =====
# ===== include =====

.PHONY: echo
# echo test content
echo:
	@echo "==> GOHOSTOS: $(GOHOSTOS)"
	@echo "==> GOPATH: $(GOPATH)"
	@echo "==> VERSION: $(VERSION)"
	@echo "==> PROJECT_MAKEFILE: $(PROJECT_MAKEFILE)"
	@echo "==> PROJECT_ABS_PATH: $(PROJECT_ABS_PATH)"
	@echo "==> PROJECT_PATH_NAME: $(PROJECT_PATH_NAME)"
	@echo "==> PROJECT_REL_PATH: $(PROJECT_REL_PATH)"

.PHONY: generate
# generate : go:generate
generate:
	#go mod tidy
	go generate ./...

