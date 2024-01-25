package config

import "gin_api/internal"

func InitConfig() {
	internal.InitDB() //数据库
}
func init() {
	internal.InitViper()
}
