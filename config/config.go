package config

import "gin_api/internal"

func InitConfig() {
	internal.InitDB()       //数据库
	internal.InitRedis()    //redis
	internal.InitES()       //初始化es
	internal.InitRabbitMQ() //初始话RabbitMQ
}
func init() {
	internal.InitViper()
}
