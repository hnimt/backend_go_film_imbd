package cache

import (
	"encoding/json"
	"log"
	microbackendfilm "micro_backend_film"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

func ConnectRedis(config microbackendfilm.TomlConfig) *redis.Client {
	db, err := strconv.Atoi(config.Redis.Db)
	if err != nil {
		log.Fatal(err)
	}

	red := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Host,
		Password: config.Redis.Pass, // no password set
		DB:       db,                // use default DB
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
