package cache

import (
	"encoding/json"
	"log"
	"micro_backend_film/config"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

func ConnectRedis(config config.TomlConfig) *redis.Client {
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

func GetCache(red *redis.Client, key string, dest interface{}) error {
	val, err := red.Get(key).Result()
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(val), &dest)
	if err != nil {
		return err
	}

	return nil
}
