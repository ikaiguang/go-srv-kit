# api

api

## generate

```shell

# mac & linux
protoc -I. -I$GOPATH/src --go_out=. --go_opt=paths=source_relative ./*.proto

# windows
# protoc -I. -I%GOPATH%/src --go_out=. --go_opt=paths=source_relative ./*.proto

```