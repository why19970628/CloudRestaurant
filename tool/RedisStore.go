package tool

import (
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/mojocn/base64Captcha"
)

// 定义一个redis结构体
type RedisStore struct {
	client *redis.Client
}

// 定义一个变量等于这个结构体
var RediStore RedisStore

// 这个结构体初始化方法 在main 里面调用
func InitRedisStore() *RedisStore {
	config := GetConfig().RedisConfig //引入redis配置文件

	// 链接上redis
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr + ":" + config.Port,
		Password: config.Password,
		DB:       config.Db,
	})

	RediStore = RedisStore{client: client}
	base64Captcha.SetCustomStore(&RediStore)

	return &RediStore
}

// set
func (rs *RedisStore) Set(id string, value string) {
	err := rs.client.Set(id, value, time.Minute*10).Err()
	if err != nil {
		log.Println(err)
	}
}

// get
func (rs *RedisStore) Get(id string, clear bool) string {
	val, err := rs.client.Get(id).Result()
	if err != nil {
		log.Println(err)
		return ""
	}
	if clear {
		err := rs.client.Del(id).Err()
		if err != nil {
			log.Println(err)
			return ""
		}
	}
	return val
}
