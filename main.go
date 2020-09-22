package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gomodule/redigo/redis"
)

func main() {
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "127.0.0.1:6379"
	}
	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisAddr)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conn := pool.Get()
		conn.Do("INCR", "read_count")
		fmt.Fprintf(w, "Success")

	})
	http.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		conn := pool.Get()
		reply, err := redis.Int(conn.Do("GET", "read_count"))
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}

		fmt.Fprintf(w, "%d", reply)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
