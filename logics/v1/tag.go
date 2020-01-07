package v1

import (
	"ginApi/pkg/e"
	"github.com/gin-gonic/gin"
)

func GetTags(c *gin.Context) {
	c.JSON(e.SUCCESS, gin.H{
		"message":"tags列表",
	})
}

func AddTag(c *gin.Context) {
	c.JSON(e.SUCCESS, gin.H{
		"message":"增加tag",
	})
}

func EditTag(c *gin.Context) {
	c.JSON(e.SUCCESS, gin.H{
		"message":"编辑tag",
	})
}

func DelTag(c *gin.Context) {
	c.JSON(e.SUCCESS, gin.H{
		"message":"删除tag",
	})
}
