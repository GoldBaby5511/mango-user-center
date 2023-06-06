module mango-user-center

go 1.15

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/ethereum/go-ethereum v1.10.4
	github.com/gin-gonic/gin v1.8.1
	github.com/go-playground/locales v0.14.0
	github.com/go-playground/universal-translator v0.18.0
	github.com/go-playground/validator/v10 v10.11.0
	github.com/google/uuid v1.3.0
	github.com/jinzhu/copier v0.3.5
	github.com/robfig/cron/v3 v3.0.0
	github.com/sirupsen/logrus v1.9.0
	github.com/spf13/viper v1.8.1
	github.com/stretchr/testify v1.8.0 // indirect
	golang.org/x/time v0.0.0-20220722155302-e5dcc9cfc0b9
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gorm.io/driver/mysql v1.3.5
	gorm.io/gorm v1.23.8
)

replace github.com/spf13/afero => github.com/spf13/afero v1.5.1
