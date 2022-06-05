package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

var mhsDB *redis.Client

func initDB() *redis.Client {
	activeDB := getEnv("REDIS_URL", "localhost:6379")

	if activeDB == "heroku" {
		opt, errDB := redis.ParseURL(activeDB)
		if errDB != nil {
			panic(errDB)
		}

		mhsDB = redis.NewClient(opt)
		return mhsDB

	} else {
		mhsDB = redis.NewClient(&redis.Options{
			Addr:     activeDB,
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       0,
		})

		return mhsDB
	}
}

func main() {
	mhsDB := initDB()
	_, err := mhsDB.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.New()
	router.SetTrustedProxies([]string{
		"localhost",
		"infralabs.cs.ui.ac.id",
		"herokuapp.com",
	})

	router1 := router.Group("")
	{
		router1.POST("", updateHandler)
	}

	port := getEnv("PORT", "8081")
	router.Run(":" + port)
}
