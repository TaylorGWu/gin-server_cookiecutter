package redis

import (
	"github.com/gomodule/redigo/redis"
	"{{cookiecutter.app_name}}/config"
)

var redisPool redis.Pool

type RedisClient struct {
	client	redis.Conn
}

func InitRedisPool() {
	cf := config.Config()
	redisConfig := cf.GetStringMap("redis")
	redisPool = redis.Pool{
		MaxIdle: redisConfig["max_idle"].(int),
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisConfig["addr"].(string))
			if err != nil {
				return nil, err
			}

			if redisConfig["auth"].(string) != "" {
				if _, err := c.Do("AUTH", redisConfig["auth"].(string)); err != nil {
					c.Close()
					return nil, err
				}
			}

			if _, err := c.Do("SELECT", redisConfig["db"].(int)); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
	}
}

func GetRedisClient() (redisClient RedisClient) {
	redisClient = RedisClient{
		client:	redisPool.Get(),
	}
	return
}

func (redisClient *RedisClient) Close() {
	redisClient.Close()
}

func (redisClient *RedisClient) Set(key string, value interface{}) (err error) {
	_, err = redisClient.client.Do("SET", key, value)
	if err != nil {
		// todo log
	}
	return
}
