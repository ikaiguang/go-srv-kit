module github.com/ikaiguang/go-srv-kit/data/migration

go 1.25.9

require gorm.io/gorm v1.31.1

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/text v0.36.0 // indirect
)

replace github.com/ikaiguang/go-srv-kit/kit => ../../kit

replace github.com/ikaiguang/go-srv-kit/data/gorm => ../gorm
