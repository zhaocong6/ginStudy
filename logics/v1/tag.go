package v1

import (
	"ginApi/models"
	"ginApi/pkg/e"
	"ginApi/pkg/setting"
	"ginApi/pkg/util"
	"github.com/gin-gonic/gin"
)

func GetTags(c *gin.Context) {

	maps := make(map[string]interface{})

	if arg := c.Query("name"); arg != "" {
		maps["name LIKE ?"] = "%" + arg + "%"
	}

	data := models.GetTags(util.GetPage(c), setting.PageSize, maps)

	c.JSON(e.SUCCESS, gin.H{
		"code": e.SUCCESS,
		"data": data,
		"msg":  e.GetMsg(e.SUCCESS),
	})
}

func AddTag(c *gin.Context) {
	c.JSON(e.SUCCESS, gin.H{
		"message": "增加tag",
	})
}

func EditTag(c *gin.Context) {
	c.JSON(e.SUCCESS, gin.H{
		"message": "编辑tag",
	})
}

func DelTag(c *gin.Context) {
	c.JSON(e.SUCCESS, gin.H{
		"message": "删除tag",
	})
}
