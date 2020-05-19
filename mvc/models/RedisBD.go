package models

import (
	"fmt"
	"hash/fnv"

	"github.com/garyburd/redigo/redis"
)

var redisPool *redis.Pool

func Hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

// Redis_db_init Init Redis
func RedisDbInit() {
	//redisAddr := "elc-tutorial-001.acnysi.0001.use2.cache.amazonaws.com:6379"
	redisAddr := "localhost:6379"
	redisPool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", redisAddr)
			return conn, err
		},
	}
}

func RedisDbSave(longUrl string) (interface{}, error, string) {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	newShortCode := Hash(longUrl)
	red, err := redisConn.Do("SET", newShortCode, longUrl)
	return red, err, fmt.Sprint(newShortCode)
}

func RedisDbGet(shortCode string) (string, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	redirectUrl, err := redis.String(redisConn.Do("GET", shortCode))
	return redirectUrl, err
}

func RedisDbDel(shortCode string) (interface{}, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	redirectUrl, err := redisConn.Do("DEL", shortCode)
	return redirectUrl, err
}

func RedisDbDelBulks(urls []string, hostName string) []string {
	redisConn := redisPool.Get()
	defer redisConn.Close()
	cad := make([]string, 0)
	for _, value := range urls {
		cad = append(cad, hostName+value)
		redisConn.Do("DEL", value)
	}
	return cad
}

func RedisDbSaveBulks(urls []string, hostName string) []string {
	redisConn := redisPool.Get()
	defer redisConn.Close()
	cad := make([]string, 0)
	for _, value := range urls {
		h := fmt.Sprint(Hash(value))
		cad = append(cad, hostName+h)
		redisConn.Do("SET", h, value)
	}
	return cad
}

/*
func RedisDbSaveBulks(urls []string, hostName string) []string {
	redisConn := redisPool.Get()
	defer redisConn.Close()
	//var cadUrls string
	cad := make([]string, 0)
	for _, value := range urls {
		h := fmt.Sprint(Hash(value))
		//cadUrls = cadUrls + h + " " + value + " "
		cad = append(cad, hostName+h)
		redisConn.Do("SET", h, value)
	}
	//log.Println("MSET" + cadUrls)
	//redisConn.Do("MSET " + cadUrls) // save to db
	return cad
}

func Redis_db_exists(shortLong string) bool {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	list_shorts, err := redis.Strings(redisConn.Do("key *"))
	if err != nil {
		return false
	} else {
		for _, value := range list_shorts {
			if shortLong == value {
				return true
			}
		}
		return false
	}

}
*/
