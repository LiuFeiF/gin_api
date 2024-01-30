package internal

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"time"
)

var RedisClient *redis.Client

func InitRedis() {

	host := viper.GetString("redis.host")
	port := viper.GetString("redis.port")
	addr := fmt.Sprintf("%s:%s", host, port)

	password := viper.GetString("redis.password")
	//RedisClient = redis.NewFailoverClient(&redis.FailoverOptions{
	//	//MasterName: "redis-master", //设置了MasterName字段，指定了Redis主节点的名称。在Redis Sentinel模式下，主节点是由哨兵进行监控和管理的。
	//	SentinelAddrs: []string{
	//		addr,
	//	},
	//	Password:     password,
	//	DialTimeout:  time.Second * 5, //DialTimeout字段是连接Redis服务器的超时时间。在这个示例中，设置为5秒。
	//	ReadTimeout:  time.Second * 3, //ReadTimeout字段是从Redis服务器读取数据的超时时间。在这个示例中，设置为3秒。
	//	WriteTimeout: time.Second * 3, //WriteTimeout字段是向Redis服务器写入数据的超时时间。在这个示例中，设置为3秒。
	//	PoolTimeout:  4 * time.Second, //PoolTimeout字段是连接池的超时时间。连接池用于管理与Redis服务器的连接。在这个示例中，设置为4秒。
	//	OnConnect: func(conn *redis.Conn) error {
	//		return nil
	//	},
	//})
	//上面是配置哨兵模式下的配置操作，但是我本地的redis没有配置，先用下面基础的连接，后续有需求的话配置一下。
	RedisClient := redis.NewClient(&redis.Options{
		Addr:         addr,            // Redis服务器地址和端口
		Password:     password,        // 如果Redis服务器需要密码验证，可以在这里设置密码
		DB:           0,               // Redis数据库索引
		DialTimeout:  time.Second * 5, //DialTimeout字段是连接Redis服务器的超时时间。在这个示例中，设置为5秒。
		ReadTimeout:  time.Second * 3, //ReadTimeout字段是从Redis服务器读取数据的超时时间。在这个示例中，设置为3秒。
		WriteTimeout: time.Second * 3, //WriteTimeout字段是向Redis服务器写入数据的超时时间。在这个示例中，设置为3秒。
		PoolTimeout:  4 * time.Second, //PoolTimeout字段是连接池的超时时间。连接池用于管理与Redis服务器的连接。在这个示例中，设置为4秒。
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		zap.S().Error("RedisCluster.Ping err" + err.Error())
		fmt.Println("RedisCluster.Ping err" + err.Error())

	}
	fmt.Println("redis 初始化成功")
}
