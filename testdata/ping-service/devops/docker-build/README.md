# build

统一端口

* http: 8080
* grpc: 50051

```shell
docker build \
		--build-arg BUILD_FROM_IMAGE=golang:1.23.10 \
		--build-arg RUN_SERVICE_IMAGE=debian:stable-20250520 \
		--build-arg APP_DIR=testdata \
		--build-arg SERVICE_NAME=ping-service \
		--build-arg VERSION=latest \
		-t ping-service:latest \
		-f ./testdata/ping-service/devops/Dockerfile .
```