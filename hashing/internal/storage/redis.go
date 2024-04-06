package storage

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

func ConnectToRedis() (*redis.Client, error) {
	// Загружаем переменные из .env файла
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPasswd := os.Getenv("REDIS_PASSWD")
	dbNum, err := strconv.Atoi(os.Getenv("DB_NUM"))
	if err != nil {
		log.Fatalf("Error converting DB_NUM to integer: %v", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPasswd,
		DB:       dbNum,
	})

	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
