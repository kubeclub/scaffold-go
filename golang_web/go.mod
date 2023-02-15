module kubeclub/scaffold_go_web

go 1.16

require (
	exexm.com/golang_common v1.0.0
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/boj/redistore v0.0.0-20180917114910-cd5dcc76aeff // indirect
	github.com/garyburd/redigo v1.6.0
	github.com/gin-gonic/contrib v0.0.0-20190526021735-7fb7810ed2a0
	github.com/gin-gonic/gin v1.7.7
	github.com/go-playground/locales v0.13.0
	github.com/go-playground/universal-translator v0.17.0
	github.com/golang-jwt/jwt/v4 v4.1.0
	github.com/gorilla/sessions v1.1.3 // indirect
	github.com/graphql-go/graphql v0.8.0
	github.com/graphql-go/handler v0.2.3
	github.com/pkg/errors v0.9.1
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.6.5
	gopkg.in/go-playground/validator.v9 v9.29.0
	gorm.io/gorm v1.22.4
)

replace exexm.com/golang_common => ./../golang_common //当前出于修改方便把 common 工程放一起，后续需要多服务共有可以独立出去
