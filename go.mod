module ginApi

go 1.13

require (
	github.com/gin-gonic/gin v1.5.0
	github.com/go-ini/ini v1.51.1 // indirect
)

replace (
	github.com/zhaocong6/ginStudy/conf => ./pkg/conf
	github.com/zhaocong6/ginStudy/middleware => ./middleware
	github.com/zhaocong6/ginStudy/models => ./models
	github.com/zhaocong6/ginStudy/pkg/setting => ./pkg/setting
	github.com/zhaocong6/ginStudy/routers => ./routers
)
