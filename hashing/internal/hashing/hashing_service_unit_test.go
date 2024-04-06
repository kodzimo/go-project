package hashing

import (
	"context"
	"testing"

	"final-project-kodzimo-hashing/internal/storage"
	pb "final-project-kodzimo-shared/proto"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/*
Этот тест проверяет, что метод CheckHash не возвращает ошибку и возвращает непустой хеш.
*/
func TestCheckHash(t *testing.T) {
	client, err := storage.ConnectToRedis()
	if err != nil {
		t.Fatalf("failed to connect to Redis: %v", err)
	}

	service := NewHashingService(client)
	req := &pb.HashRequest{Payload: "test"}

	// Создаем хеш
	createResp, err := service.CreateHash(context.Background(), req)
	assert.NoError(t, err)

	// Проверяем хеш
	checkReq := &pb.HashRequest{Payload: createResp.GetHash()}
	checkResp, err := service.CheckHash(context.Background(), checkReq)

	assert.NoError(t, err)
	if checkResp != nil {
		assert.NotEmpty(t, checkResp.Hash)
	}

	/*
		В этом примере мы добавили новый запрос missingReq, который ищет хеш, которого нет в базе данных. Мы ожидаем,
		что CheckHash вернет ошибку с кодом codes.NotFound в этом случае. Это позволяет нам проверить, что ваш код
		правильно обрабатывает ситуацию, когда хеш не найден.

		Пожалуйста, учтите, что этот пример предполагает, что в вашей базе данных Redis нет значения с ключом “missing”.
		Если это не так, вам нужно будет выбрать другой ключ для этого теста.
	*/

	// Добавляем проверку на случай, когда хеша нет в базе данных
	missingReq := &pb.HashRequest{Payload: "missing"}
	_, err = service.CheckHash(context.Background(), missingReq)

	assert.Error(t, err)
	assert.Equal(t, codes.NotFound, status.Code(err))
}

/*
Этот тест проверяет, что метод GetHash не возвращает ошибку и возвращает тот же хеш, который был создан.
*/
func TestGetHash(t *testing.T) {
	client, err := storage.ConnectToRedis()
	if err != nil {
		t.Fatalf("failed to connect to Redis: %v", err)
	}

	service := NewHashingService(client)
	req := &pb.HashRequest{Payload: "test"}

	// Создаем хеш
	createResp, err := service.CreateHash(context.Background(), req)
	assert.NoError(t, err)

	// Получаем хеш
	getReq := &pb.HashRequest{Payload: createResp.GetHash()}
	getResp, err := service.GetHash(context.Background(), getReq)

	assert.NoError(t, err)
	assert.NotNil(t, getResp)
	assert.Equal(t, req.GetPayload(), getResp.GetHash())

	/*
		В этом примере мы добавили новый запрос missingReq, который ищет хеш, которого нет в базе данных. Мы ожидаем,
		что GetHash вернет ошибку с кодом codes.NotFound в этом случае. Это позволяет нам проверить, что ваш код правильно обрабатывает ситуацию, когда хеш не найден.

		Пожалуйста, учтите, что этот пример предполагает, что в вашей базе данных Redis нет значения с ключом “missing”.
		Если это не так, вам нужно будет выбрать другой ключ для этого теста.
	*/

	// Добавляем проверку на случай, когда хеша нет в базе данных
	missingReq := &pb.HashRequest{Payload: "missing"}
	_, err = service.GetHash(context.Background(), missingReq)

	assert.Error(t, err)
	assert.Equal(t, codes.NotFound, status.Code(err))
}

/*
Этот тест проверяет, что метод CreateHash не возвращает ошибку и возвращает непустой хеш.
*/
func TestCreateHash(t *testing.T) {
	client, err := storage.ConnectToRedis()
	if err != nil {
		t.Fatalf("failed to connect to Redis: %v", err)
	}

	service := NewHashingService(client)
	req := &pb.HashRequest{Payload: "test"}

	resp, err := service.CreateHash(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Hash)
}

/*
Для запуска тестов вы можете использовать go test -run 'Имя_теста'.
*/
