package cache

import (
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis"
)

func ConnectRedis() *redis.Client {
	red := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	log.Println("Connect redis successfully")
	return red
}

func SetCache(red *redis.Client, key string, value interface{}) {
	data, err := json.Marshal(value)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	err = red.Set(key, data, 24*time.Hour).Err()
	if err != nil {
		log.Println(err)
	}
}

func GetCache(red *redis.Client, key string, dest interface{}) {
	val, err := red.Get(key).Result()
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal([]byte(val), &dest)
	if err != nil {
		log.Println(err)
	}
}
