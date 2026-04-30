module github.com/ikaiguang/go-srv-kit/kratos/auth/redis

go 1.25.9

require (
	github.com/go-kratos/kratos/v2 v2.9.2
	github.com/ikaiguang/go-srv-kit/kit v0.0.0
	github.com/ikaiguang/go-srv-kit/kratos v0.0.0
	github.com/redis/go-redis/v9 v9.18.0
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/go-playground/form/v4 v4.3.0 // indirect
	github.com/golang-jwt/jwt/v5 v5.3.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/rs/xid v1.6.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	golang.org/x/sys v0.43.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260414002931-afd174a4e478 // indirect
	google.golang.org/grpc v1.80.0 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/ikaiguang/go-srv-kit/kit => ../../../kit

replace github.com/ikaiguang/go-srv-kit/kratos => ../..
