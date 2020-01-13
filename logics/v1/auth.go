package v1

import (
	"ginApi/models"
	"ginApi/pkg/e"
	"ginApi/pkg/util"
	"github.com/gin-gonic/gin"
)

func GetAuth(c *gin.Context) {
	_, ok := models.GetAuth(c.PostForm("username"), c.PostForm("password"))

	code := e.SUCCESS
	data := make(map[string]interface{})

	if ok {
		token, err := util.GenerateToken(c.PostForm("username"), c.PostForm("password"))
		if err != nil {
			code = e.ERROR_AUTH_TOKEN
		}

		data["token"] = token
	} else {
		code = e.INVALID_PARAMS
	}

	c.JSON(e.SUCCESS, gin.H{
		"code": code,
		"data": data,
		"msg":  e.GetMsg(code),
	})
}
