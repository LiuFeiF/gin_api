package main

import (
	"gin_api/config"
	"gin_api/route"
	"github.com/spf13/viper"
)

func main() {
	config.InitConfig()                    //初始化读取配置文件的viper中间件对后续需要用的sql和redis都有很好的帮助
	port := viper.GetString("server.port") //本地调试或者dockers部署是可以对外链接的port接口
	r := route.InitRoutes()                //路由文件夹
	err := r.Run(port)
	if err != nil {
		return
	}

}
