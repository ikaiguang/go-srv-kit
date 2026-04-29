module github.com/ikaiguang/go-srv-kit

go 1.25.9

// 注意：以下依赖包含 testdata/ 中使用的模块。
// testdata/ 被 Go 的 ./... 模式排除，go mod tidy 会移除这些依赖。
// 请勿对根模块执行 go mod tidy，或执行后手动恢复 testdata 所需的依赖。
require (
	github.com/envoyproxy/protoc-gen-validate v1.3.3
	github.com/go-kratos/kratos/v2 v2.9.2
	github.com/google/wire v0.7.0
	github.com/gorilla/websocket v1.5.3
	github.com/ikaiguang/go-srv-kit/kit v0.0.0
	github.com/ikaiguang/go-srv-kit/kratos v0.0.0
	google.golang.org/grpc v1.80.0 // indirect
	google.golang.org/protobuf v1.36.11
)

replace github.com/ikaiguang/go-srv-kit/kit => ./kit

replace github.com/ikaiguang/go-srv-kit/kratos => ./kratos

replace github.com/ikaiguang/go-srv-kit/data/consul => ./data/consul

replace github.com/ikaiguang/go-srv-kit/data/etcd => ./data/etcd

replace github.com/ikaiguang/go-srv-kit/data/gorm => ./data/gorm

replace github.com/ikaiguang/go-srv-kit/data/jaeger => ./data/jaeger

replace github.com/ikaiguang/go-srv-kit/data/migration => ./data/migration

replace github.com/ikaiguang/go-srv-kit/data/mongo => ./data/mongo

replace github.com/ikaiguang/go-srv-kit/data/mysql => ./data/mysql

replace github.com/ikaiguang/go-srv-kit/data/postgres => ./data/postgres

replace github.com/ikaiguang/go-srv-kit/data/rabbitmq => ./data/rabbitmq

replace github.com/ikaiguang/go-srv-kit/data/redis => ./data/redis

replace github.com/ikaiguang/go-srv-kit/service => ./service

require (
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible // indirect
	github.com/lestrrat-go/strftime v1.1.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.1 // indirect
	golang.org/x/net v0.53.0 // indirect
	golang.org/x/sys v0.43.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260414002931-afd174a4e478 // indirect
)
