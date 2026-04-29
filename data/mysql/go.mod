module github.com/ikaiguang/go-srv-kit/data/mysql

go 1.25.9

require (
	github.com/go-sql-driver/mysql v1.9.3
	github.com/ikaiguang/go-srv-kit/data/gorm v0.0.0
	github.com/ikaiguang/go-srv-kit/kit v0.0.0
	github.com/stretchr/testify v1.11.1
	google.golang.org/protobuf v1.36.11
	gorm.io/driver/mysql v1.6.0
	gorm.io/gorm v1.31.1
)

require (
	filippo.io/edwards25519 v1.2.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible // indirect
	github.com/lestrrat-go/strftime v1.1.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	golang.org/x/text v0.36.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/hints v1.1.2 // indirect
)

replace github.com/ikaiguang/go-srv-kit/kit => ../../kit

replace github.com/ikaiguang/go-srv-kit/data/gorm => ../gorm
