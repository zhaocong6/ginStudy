package routers

import (
	v1 "ginApi/logics/v1"
	"ginApi/pkg/setting"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	gin.SetMode(setting.RunMode)

	addV1Routes(r)

	return r
}

//增加v1路由
func addV1Routes(r *gin.Engine) *gin.RouterGroup {
	apiV1 := r.Group("/api/v1")
	apiV1.POST("auth", v1.GetAuth)

	{
		apiV1.GET("/tags", v1.GetTags)
		apiV1.POST("/tags", v1.AddTag)
		apiV1.PUT("/tags/:id", v1.EditTag)
		apiV1.DELETE("/tags/:id", v1.DelTag)
	}

	return apiV1
}
