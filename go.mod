module blog

go 1.16

require github.com/gin-gonic/gin v1.7.7

require github.com/sirupsen/logrus v1.8.1

require (
	github.com/casbin/casbin/v2 v2.37.4
	gorm.io/driver/postgres v1.3.1 // indirect
	gorm.io/driver/sqlserver v1.3.1 // indirect

)

require (
	github.com/antonfisher/nested-logrus-formatter v1.3.1 // indirect
	github.com/casbin/gorm-adapter/v3 v3.5.0
	github.com/go-playground/validator/v10 v10.10.1 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/howeyc/fsnotify v0.9.0 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/pilu/config v0.0.0-20131214182432-3eb99e6c0b9a // indirect
	github.com/pilu/fresh v0.0.0-20190826141211-0fa698148017 // indirect
	github.com/pkg/errors v0.9.1
	github.com/spf13/viper v1.10.1
	gorm.io/driver/mysql v1.3.3
	gorm.io/gorm v1.23.4
)
