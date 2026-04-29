# build

示例服务镜像构建说明。

## 统一端口

* http: 8080
* grpc: 50051

## 构建示例

```bash
docker build \
  --build-arg BUILD_FROM_IMAGE=golang:1.25.9 \
  --build-arg RUN_SERVICE_IMAGE=debian:stable-20250520 \
  --build-arg APP_DIR=testdata \
  --build-arg SERVICE_NAME=ping-service \
  --build-arg VERSION=latest \
  -t ping-service:latest \
  -f ./testdata/ping-service/devops/docker-build/Dockerfile .
```

如果需要不同的 Go 基础镜像版本，可通过 `BUILD_FROM_IMAGE` 覆写。
