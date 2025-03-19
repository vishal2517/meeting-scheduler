package db

import (
	"fmt"
	"log"
	"os"
	"context"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/redis/go-redis/v9"
)

var DB *gorm.DB
var RedisClient *redis.Client


// ConnectRedis initializes Redis connection
func ConnectRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"), // Redis host (from env variable)
		Password: "",                      // No password by default
		DB:       0,                        // Default DB
	})

	ctx := context.Background()
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}

	fmt.Println("✅ Connected to Redis!")
}

func ConnectDatabase() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = database
}
