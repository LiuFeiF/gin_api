package config

import "gin_api/internal"

func InitConfig() {
	//internal.InitDB()//数据库
	internal.InitViper() //获取配置文件相关函数
}
