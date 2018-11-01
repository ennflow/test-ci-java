package models

import (
	"github.com/go-redis/redis"

	"fmt"
	"time"
)

var RedisClient *redis.Client

func InitRedisClient() {
	client := redis.NewClient(
		&redis.Options{
			Addr:"redis-test-qinzhao.enncloudtest.tenxcloud.cc:61275",
			Password:"qinzhao",
			DB:0,
		},
	)
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println("get redisClient failed", err)
		return
	}
	fmt.Println("pong===",pong)

	if client != nil {
		fmt.Println("redis-client",client)
		RedisClient = client
	}

}

func SetValue()(string,error) {

	return RedisClient.Set("xinzhiyun","xinzhiyun-redis-test",0*time.Second).Result()

}

func GetValue()(string,error)  {

	return RedisClient.Get("xinzhiyun").Result()

}