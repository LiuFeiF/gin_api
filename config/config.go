package config

import "gin_api/internal"

func InitConfig() {
	internal.InitDB()    //数据库
	internal.InitRedis() //redis
	internal.InitES()    //初始化es
}
func init() {
	internal.InitViper()
}
