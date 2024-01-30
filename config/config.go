package config

import "gin_api/internal"

func InitConfig() {
	internal.InitDB()    //数据库
	internal.InitRedis() //redis
}
func init() {
	internal.InitViper()
}
