# api

api

## generate

```shell

# mac & linux
protoc -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/ikaiguang/go-srv-kit/third_party/ \
  --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  --go-errors_out=paths=source_relative:. \
  ./*.proto

# windows
# protoc -I. -I%GOPATH%/src --go_out=paths=source_relative:. ./*.proto
#protoc -I. \
#  -I%GOPATH%/src \
#  -I%GOPATH%/src/github.com/ikaiguang/go-srv-kit/third_party/ \
#  --go_out=paths=source_relative:. \
#  --go-errors_out=paths=source_relative:. \
#  ./*.proto

```