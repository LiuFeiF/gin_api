package route

import (
	api "gin_api/controller/v1"
	"gin_api/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	lb := middleware.NewLoadBalancer()
	r := gin.Default()
	r.Any("/", lb.HandleRequest) //做个简单的负载均衡政策的分流代理，或许有机会可以设置nginx代理做分流。
	r.Use(middleware.Cors())     //解决一下跨域问题，桶网上用的最多的跨域中间件判断
	v1 := r.Group("api/v1")
	{
		v1.GET("test", api.Test)
	}
	return r
}
