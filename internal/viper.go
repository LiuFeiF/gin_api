package internal

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
)

func InitViper() {
	workDir, _ := os.Getwd() //工作目录
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "test"
	}
	viper.SetConfigName(fmt.Sprintf("config_%s", env)) //配置文件文件名
	viper.SetConfigType("yml")                         //配置文件的类型
	viper.AddConfigPath(workDir + "/config")           //工作目录
	err := viper.ReadInConfig()
	if err != nil {
		zap.S().Error("ReadInConfig err:" + err.Error())
		//log.WaymonLogger.Error("ReadInConfig err:" + err.Error())
		fmt.Println("ReadInConfig err:" + err.Error())
	}
	fmt.Println("Viper 初始化成功")
}
