package main

import (
	dbmigrate "github.com/ikaiguang/go-srv-kit/example/cmd/migration/migrate"
)

// go run ./cmd/migration/... -conf=./configs
func main() {
	dbmigrate.Run(dbmigrate.WithCloseEngineHandler())
}
