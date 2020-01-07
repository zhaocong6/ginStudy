package util

import (
	"ginApi/pkg/setting"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

//获取下一页取值范围
func GetPage(c *gin.Context) (result int){

	page, _ := com.StrTo(c.Query("page")).Int()

	if page > 0 {
		result = (page - 1) * setting.PageSize
	}

	return
}
