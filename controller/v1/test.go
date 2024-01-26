package v1

import (
	"gin_api/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Test(c *gin.Context) {
	Service := service.TestService{}
	if err := c.ShouldBind(&Service); err != nil {
		code := 400
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    "入参失败",
			"data":   nil,
		})
		return
	}
	res := Service.TestInfo(c.Request.Context()) //掉用service层进行相关逻辑处理
	c.JSON(http.StatusOK, res)
}
