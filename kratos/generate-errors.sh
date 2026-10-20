#!/usr/bin/env bash

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "${SCRIPT_DIR}/.." && pwd)"
GENERATOR_VERSION="v0.0.3"
MODULE_CACHE="$(go env GOMODCACHE)"
GENERATOR_SOURCE="${MODULE_CACHE}/github.com/ikaiguang/protoc-gen-go-errors@${GENERATOR_VERSION}"
TEMP_DIR="$(mktemp -d /tmp/go-srv-kit-errors-v3.XXXXXX)"

cleanup() {
	rm -rf "${TEMP_DIR}"
}
trap cleanup EXIT

GOWORK=off go mod download "github.com/ikaiguang/protoc-gen-go-errors@${GENERATOR_VERSION}"
cp -a "${GENERATOR_SOURCE}/." "${TEMP_DIR}/source"
chmod -R u+w "${TEMP_DIR}/source"
sed -i \
	's#github.com/go-kratos/kratos/v2/errors#github.com/go-kratos/kratos/v3/errors#g' \
	"${TEMP_DIR}/source/errors.go"

(
	cd "${TEMP_DIR}/source"
	GOWORK=off go build -o "${TEMP_DIR}/protoc-gen-go-errors" .
)

(
	cd "${REPO_ROOT}/kratos/auth"
	PATH="${TEMP_DIR}:${PATH}" protoc \
		--proto_path=. \
		--proto_path=../../third_party \
		--go_out=paths=source_relative:. \
		--go-errors_out=paths=source_relative:. \
		auth.kit.proto
)

(
	cd "${REPO_ROOT}"
	PATH="${TEMP_DIR}:${PATH}" protoc \
		--proto_path=. \
		--proto_path=third_party \
		--go_out=paths=source_relative:. \
		--go-errors_out=paths=source_relative:. \
		kratos/error/error.kit.proto
)
