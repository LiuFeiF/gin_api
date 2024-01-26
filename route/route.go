package route

import (
	api "gin_api/controller/v1"
	"gin_api/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors()) //解决一下跨域问题，桶网上用的最多的跨域中间件判断
	v1 := r.Group("api/v1")
	{
		v1.GET("test", api.Test)
	}
	return r
}
