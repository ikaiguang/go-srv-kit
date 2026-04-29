module github.com/ikaiguang/go-srv-kit/data/postgres

go 1.25.9

require (
	github.com/ikaiguang/go-srv-kit/data/gorm v0.0.0
	github.com/jackc/pgx/v5 v5.9.2
	github.com/stretchr/testify v1.11.1
	google.golang.org/protobuf v1.36.11
	gorm.io/driver/postgres v1.6.0
	gorm.io/gorm v1.31.1
)

require (
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/ikaiguang/go-srv-kit/kit v0.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	golang.org/x/sync v0.20.0 // indirect
	golang.org/x/text v0.36.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/hints v1.1.2 // indirect
)

replace github.com/ikaiguang/go-srv-kit/kit => ../../kit

replace github.com/ikaiguang/go-srv-kit/data/gorm => ../gorm
