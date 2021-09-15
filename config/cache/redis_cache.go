package cache

import (
	"encoding/json"
	"log"
	"micro_backend_film/pkg/entity"
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

func SetCache(red *redis.Client, key string, value []entity.Film) {
	data, err := json.Marshal(value)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	err = red.Set(key, data, 30*time.Second).Err()
	if err != nil {
		log.Println(err)
	}
}

func GetCache(red *redis.Client, key string) ([]entity.Film, error) {
	var data []entity.Film
	val, err := red.Get(key).Result()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(val), &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
