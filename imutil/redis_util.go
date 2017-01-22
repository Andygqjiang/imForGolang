package imutil

import (
	"github.com/garyburd/redigo/redis"
	"time"
	m "../model"
)

var redisPool *redis.Pool

const (
	SET = "SET"
	GET = "GET"
	EXPIRE = "EXPIRE"
	EXISTS = "EXISTS"
	HSET = "HSET"
	HGET = "HGET"
	HGETALL = "HGETALL"
	HDEL = "HDEL"
)

func init() {
	redisPool = &redis.Pool{
		MaxIdle: 3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error){
			c, err := redis.Dial("tcp", "localhost:6379")
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", "123456"); err != nil {
				c.Close()
				return nil, err
			}
			if _, err := c.Do("SELECT", "0"); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func getConn() redis.Conn {
	return redisPool.Get()
}

func Set(key string, value interface{}, expire int)  {
	conn := getConn()
	conn.Do(SET, key, value)
	conn.Do(EXPIRE, key, expire)
	conn.Close()
}

func Get(key string) string {
	conn := getConn()
	result, err := redis.String(conn.Do(GET, key))
	if err != nil {
		Ilog("redis Get error: ", err)
	}
	conn.Close()
	return result
}

func Hset(key, field string, value interface{}, expire int)  {
	conn := getConn()
	conn.Do(HSET, key, field, value)
	conn.Do(EXPIRE, key, expire)
	conn.Close()
}

func Hget(key, field string) *m.Message {
	conn := getConn()
	result, err := redis.Bytes(conn.Do(HGET, key, field))
	if err != nil {
		Ilog("redis Hget error: ", err)
	}
	conn.Close()
	return Parse(result)
}

func HgetAll(key string) map[string]string {
	conn := getConn()
	result, err := redis.StringMap(conn.Do(HGETALL, key))
	if err != nil {
		Ilog("redis HgetAll error: ", err)
	}
	conn.Close()
	return result
}

func Hdel(key, field string)  {
	conn := getConn()
	conn.Do(HDEL, key, field)
	conn.Close()
}





