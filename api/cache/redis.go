package cache

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
)

// Redis struct
type Redis struct {
	Address  string
	Password string
	Expire   time.Duration
	DB       int
}

// IRedis interface
type IRedis interface {
	Ping()
	Connect() *redis.Client
	Set(data interface{})
	Get(key string, data interface{}) string
}

func marshallStruct(data interface{}) (string, []byte, error) {
	stringify, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return "", nil, err
	}

	var result map[string]interface{}
	if err = json.Unmarshal([]byte(stringify), &result); err != nil {
		return "", nil, err
	}
	return result["uuid"].(string), stringify, nil
}

// Connect connect to redis
func (r Redis) Connect() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     r.Address,
		Password: r.Password,
		DB:       r.DB,
	})
}

// Ping checking ping to redis
func (r *Redis) Ping() {
	client := redis.NewClient(&redis.Options{
		Addr:     r.Address,
		Password: r.Password,
		DB:       r.DB,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(fmt.Errorf("Redis is not up, Reason: %v", err))
	}
	fmt.Println(pong)
}

// Set store cache
func (r *Redis) Set(data interface{}) {
	client := r.Connect()
	id, json, err := marshallStruct(data)
	if err != nil {
		fmt.Println(err)
	}
	client.Set(id, json, r.Expire)
}

// Get retrieve cache and save data if possible
func (r *Redis) Get(key string, data interface{}) string {
	client := r.Connect()
	value, err := client.Get(key).Result()
	if err == redis.Nil {
		if data != nil {
			fmt.Println(fmt.Errorf("key: %s does not exist", key))
			fmt.Println(fmt.Errorf("Store new data into key: %s", key))
			id, stringify, err := marshallStruct(data)
			if err != nil {
				fmt.Println(err)
			}
			client.Set(id, stringify, r.Expire)
			return ""
		}
		fmt.Println("Skipping")
		return ""
	} else if err != nil {
		fmt.Println(err)
		return ""
	}
	return value
}
