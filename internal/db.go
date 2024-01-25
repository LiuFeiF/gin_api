package internal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"strings"
)

var (
	DB        *gorm.DB
	WriteConn string
	_         string
)

func InitDB() {
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	database := viper.GetString("mysql.database")
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	charset := viper.GetString("mysql.charset")
	WriteConn = strings.Join([]string{username, ":", password, "@tcp(", host, ":", port, ")/", database, "?charset=", charset, "&parseTime=true"}, "")

	var gormLogger logger.Interface
	if gin.Mode() == "debug" {
		gormLogger = logger.Default.LogMode(logger.Info)
	} else {
		gormLogger = logger.Default //根据当前的运行模式（通过gin.Mode()获取），设置日志记录级别。如果是调试模式，则将日志记录级别设置为Info，否则使用默认级别
	}
	//通过gorm.Open函数打开MySQL数据库连接，并传入gorm.Config进行一些额外的配置，比如日志记录和命名策略。主要这里是为了更清晰的展示我们运行的sql语句
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       WriteConn,
		DefaultStringSize:         256,   //数据库字符串最大长度
		DisableDatetimePrecision:  true,  //禁用datetime精度
		DontSupportRenameIndex:    true,  //重命名索引
		DontSupportRenameColumn:   true,  //重命名列 兼容低版本
		SkipInitializeWithVersion: false, //根据版本自动配置
	}), &gorm.Config{
		Logger: gormLogger, //输出日志
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //table不加s
		},
	})
	if err != nil {
		fmt.Println("db Open err:", err)
	}
	DB = db //这个地方主要是把gorm功能加入到数据库中//成功打开数据库连接后，将连接赋值给全局变量DB，以便后续在应用程序中使用。
	////主从配置//这里其实可以进行数据库的均衡配置一个数据库进行写一个进行读，或者一主多从，但是我就一个数据库就先这样把
	//_ = DB.Use(dbresolver.
	//	Register(dbresolver.Config{
	//		Sources: []gorm.Dialector{
	//			mysql.Open(WriteConn), //写操作
	//		},
	//		Replicas: []gorm.Dialector{
	//			mysql.Open(WriteConn), //读操作
	//		},
	//		Policy: dbresolver.RandomPolicy{}, //负载均衡政策
	//	}).SetMaxIdleConns(10).SetMaxOpenConns(30).SetConnMaxLifetime(time.Hour * 1))
	fmt.Println("mysql 初始化成功")
}
