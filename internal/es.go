package internal

import (
	"fmt"
	"gin_api/pkg/log"
	"github.com/olivere/elastic/v7"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var ESClient *elastic.Client //全局生命，后续业户函数进行引用

type Item struct {
	ID    int    `db:"id"`
	Title string `db:"title"`
}

func InitES() {
	host := viper.GetString("es.host")
	port := viper.GetString("es.port")
	// auth_user := viper.GetString("es.auth_user")
	// auth_password := viper.GetString("es.auth_password")
	addr := fmt.Sprintf("http://%s:%s", host, port)
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(addr))
	if err != nil {
		zap.S().Error("NewClient err:" + err.Error())
		log.Errorf("NewClient err:" + err.Error())
		log.Errorf("NewClient err:" + err.Error())
	}
	ESClient = client
	log.Info("es 初始化成功")

}
