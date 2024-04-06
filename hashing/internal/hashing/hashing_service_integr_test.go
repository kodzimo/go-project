package hashing

import (
	"context"
	"testing"

	"final-project-kodzimo-hashing/internal/storage"
	pb "final-project-kodzimo-shared/proto"

	"github.com/stretchr/testify/assert"
)

/*
Этот тест проверяет, что метод CreateHash не только создает хеш, но и сохраняет его в базе данных Redis.
*/

func TestCreateHashIntegration(t *testing.T) {
	// Подключаемся к реальной базе данных Redis
	client, err := storage.ConnectToRedis()
	if err != nil {
		t.Fatalf("failed to connect to Redis: %v", err)
	}

	service := NewHashingService(client)
	req := &pb.HashRequest{Payload: "test"}

	// Создаем хеш
	resp, err := service.CreateHash(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Hash)

	// Проверяем, что хеш был сохранен в Redis
	hash, err := client.Get(context.Background(), resp.Hash).Result()
	assert.NoError(t, err)
	assert.Equal(t, req.Payload, hash)
}
