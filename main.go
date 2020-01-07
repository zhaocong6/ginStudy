package main

import (
	"fmt"
	"ginApi/pkg/setting"
	"ginApi/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := routers.InitRouter()

	run(router)
}

//启动服务
func run(r *gin.Engine) {
	r.Run(fmt.Sprintf(":%d", setting.HTTPPort))
}
