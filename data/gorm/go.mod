module github.com/ikaiguang/go-srv-kit/data/gorm

go 1.25.9

require (
	github.com/google/uuid v1.6.0
	github.com/ikaiguang/go-srv-kit/kit v0.0.0
	github.com/stretchr/testify v1.11.1
	gorm.io/gorm v1.31.1
	gorm.io/hints v1.1.2
)

require (
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	golang.org/x/text v0.36.0 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/ikaiguang/go-srv-kit/kit => ../../kit
