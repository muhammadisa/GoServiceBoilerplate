package cache

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/muhammadisa/go-service-boilerplate/api/utils/message"
	uuid "github.com/satori/go.uuid"
)

// Redis struct
type Redis struct {
	Address  string
	Password string
	Expire   time.Duration
	Debug    bool
	DB       int
}

// IRedis interface
type IRedis interface {
	Ping()
	Connect() *redis.Client
	Set(data interface{})
	Get(key string, data interface{}) string
}

func redisLogger(operation string, even string, key string, debug bool) {
	if debug {
		fmt.Println(fmt.Sprintf("%s CACHE WITH KEY: %s IS %s", operation, key, even))
	}
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
	structName := message.GetType(data)
	identifier := fmt.Sprintf("%s:%s", structName, result["id"].(string))
	return identifier, stringify, nil
}

// Key decide proper cache key
func Key(data interface{}, id uuid.UUID) string {
	return fmt.Sprintf("%s:%s", message.GetType(data), id.String())
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
	redisLogger("SET", "SUCCESS", id, r.Debug)
}

// Get retrieve cache and save data if possible
func (r *Redis) Get(key string) string {
	client := r.Connect()
	value, err := client.Get(key).Result()
	if err == redis.Nil {
		redisLogger("GET", "NIL", key, r.Debug)
		return "nil"
	} else if err != nil {
		redisLogger("GET", "ERROR", key, r.Debug)
		fmt.Println(err)
		return ""
	}
	redisLogger("GET", "SUCCESS", key, r.Debug)
	return value
}
