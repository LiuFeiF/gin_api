package main

import (
	"gin_api/config"
	"gin_api/pkg/log"
	"gin_api/route"
	"github.com/spf13/viper"
)

func main() {
	config.InitConfig()                    //初始化读取配置文件的viper中间件对后续需要用的sql和redis都有很好的帮助
	log.InitWithConfig()                   //过年宅在家弄了日志模块和es的初始化,这样控制台就可以更加直观的查看日志和排错了，但是有一说一真的很难啊，很多地方真的很有意义阅读一些优秀的开源配置作品。
	port := viper.GetString("server.port") //本地调试或者docker部署是可以对外链接的port接口
	r := route.InitRoutes()                //路由文件夹
	err := r.Run(port)
	if err != nil {
		return
	}

}
