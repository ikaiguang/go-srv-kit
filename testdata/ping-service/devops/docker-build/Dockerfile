# service arg
ARG BUILD_FROM_IMAGE=golang:1.23.10
ARG RUN_SERVICE_IMAGE=debian:stable-20250520
ARG APP_DIR=app
ARG SERVICE_NAME=ping-service
ARG VERSION=v1.0.0

# system arg
ARG GOPATH=/myworkspace/golang
ARG GOPROXY=https://goproxy.cn,direct
ARG GOPRIVATE=gitlab.uufff.com
ARG GOMODCACHE=/myworkspace/golang/pkg/mod
ARG GOCACHE=/myworkspace/golang/build_cache
ARG WORK_DIR=/myworkspace/golang/src/workspace
ARG TZ=Asia/Shanghai

FROM ${BUILD_FROM_IMAGE} AS builder

# service arg
ARG APP_DIR
ARG SERVICE_NAME
ARG VERSION

# golang arg
ARG GOPATH
ARG GOPROXY
ARG GOPRIVATE
ARG GOMODCACHE
ARG GOCACHE
ARG WORK_DIR
ARG TZ

# golang env
ENV GO111MODULE=on
ENV GOPATH=${GOPATH}
ENV GOPROXY=${GOPROXY}
ENV GOPRIVATE=${GOPRIVATE}
ENV GOMODCACHE=${GOMODCACHE}
ENV GOCACHE=${GOCACHE}
ENV TZ=${TZ}

WORKDIR ${WORK_DIR}

COPY . .

RUN go mod download && go mod tidy

# go build -ldflags="-s -w -X 'main.Version=1.0.0' -extldflags '-static'"
RUN go build -ldflags="-s -w -X 'main.Version=${VERSION}' -extldflags '-static'" \
    -o ./bin/ping-service \
    ./${APP_DIR}/${SERVICE_NAME}/cmd/ping-service/main.go

FROM ${RUN_SERVICE_IMAGE}

# arg
ARG WORK_DIR
ARG SERVICE_NAME
ARG TZ

# env
ENV TZ=$TZ
ENV DEBIAN_FRONTEND=noninteractive

WORKDIR ${WORK_DIR}

# https://developer.aliyun.com/mirror/debian
# cat /etc/os-release && echo -e "\n\n\n" && cat /etc/apt/sources.list.d/debian.sources
#RUN sed -i 's|deb.debian.org|mirrors.aliyun.com|g' /etc/apt/sources.list.d/debian.sources

RUN apt-get update && \
    apt-get install -y --no-install-recommends -y tzdata curl less && \
    ln -sf /usr/share/zoneinfo/${TZ} /etc/localtime && \
    echo ${TZ} > /etc/timezone && \
    apt-get autoremove -y && apt-get autoclean -y && apt-get clean -y && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder ${WORK_DIR}/bin/ping-service ./ping-service
#COPY --from=builder ${WORK_DIR}/${APP_DIR}/${SERVICE_NAME}/configs ./configs

# http 8080 8081
EXPOSE 8080
# grpc 50051 50052
EXPOSE 50051

ENTRYPOINT ["./ping-service", "-conf=./configs"]
