package redisqgo

// Redigo
import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

// NewPool : Creates New Redis Connectio Pool
func NewPool(maxIdle int, maxActive int, hostname string, passw string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:   maxIdle,
		MaxActive: maxActive,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", hostname)
			if err != nil {
				panic(err.Error())
			}

			response, err := c.Do("AUTH", passw)
			if err != nil {
				panic(err.Error())
			}

			fmt.Println(response)
			return c, err
		},
	}
}

// Enqueue : Add Item to List (Queue)
func Enqueue(c *redis.Conn, list string, item string) {
	(*c).Do("LPUSH", list, item)
}

// Dequeue : Remove Item to List (Queue)
func Dequeue(c *redis.Conn, list string) string {
	val, err := redis.String((*c).Do("RPOP", list))
	if err != nil {
		fmt.Println(err)
	}
	return val
}

// GetQueueItems : Get queue items
func GetQueueItems(c *redis.Conn, list string) []string {
	value, err := redis.Strings((*c).Do("LRANGE", list, 0, -1))
	if err != nil {
		fmt.Println(err)
	}
	return value
}
