package hashing

import (
	"context"
	"crypto/sha256"
	"fmt"

	pb "final-project-kodzimo-shared/proto"

	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/*
В этом примере HashingService содержит клиента Redis, который используется для взаимодействия с базой данных.
Метод CreateHash вычисляет SHA-256 хеш от входных данных и сохраняет его в базе данных Redis.

Функция ConnectToRedis устанавливает соединение с сервером Redis и возвращает клиента Redis, который затем
передается в HashingService.
*/

type HashingService struct {
	redisClient *redis.Client
}

func NewHashingService(redisClient *redis.Client) *HashingService {
	return &HashingService{redisClient: redisClient}
}

/*
Метод CheckHash. Этот метод будет принимать входные данные, проверять, существует ли уже хеш для этих данных
в базе данных, и возвращать результат.
*/

func (s *HashingService) CheckHash(ctx context.Context, req *pb.HashRequest) (*pb.HashResponse, error) {
	// Получаем данные из запроса
	payload := req.GetPayload()

	// Ищем хеш в базе данных
	hash, err := s.redisClient.Get(ctx, payload).Result()
	if err != nil {
		// Если произошла ошибка при поиске хеша, возвращаем ошибку
		if err == redis.Nil {
			return nil, status.Errorf(codes.NotFound, "hash not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get hash: %v", err)
	}

	// Если хеш найден, возвращаем его
	return &pb.HashResponse{Hash: hash}, nil
}

/*
Метод `GetHash` принимает один параметр - экземпляр структуры `HashRequest`, которая определена в вашем файле
`hashing.proto`. Структура `HashRequest` содержит одно поле `payload` типа `string`, которое представляет
собой данные, для которых вы хотите получить хеш.

Вот как это выглядит в коде:

```proto
// The request message containing the payload's data
message HashRequest {
  string payload = 1;
}
```

Таким образом, когда вы вызываете метод `GetHash`, вы передаете ему `HashRequest`, содержащий `payload`,
для которого вы хотите получить хеш. Например:

```go
req := &pb.HashRequest{Payload: "your data"}
res, err := hashingService.GetHash(ctx, req)
```

В этом примере `"your data"` - это данные, для которых вы хотите получить хеш. `ctx` - это контекст, который
используется для управления временем выполнения и отменой запроса.
*/

func (s *HashingService) GetHash(ctx context.Context, req *pb.HashRequest) (*pb.HashResponse, error) {
	// Получаем данные из запроса
	payload := req.GetPayload()

	// Ищем хеш в базе данных
	hash, err := s.redisClient.Get(ctx, payload).Result()
	if err != nil {
		// Если произошла ошибка при поиске хеша, возвращаем ошибку
		if err == redis.Nil {
			return nil, status.Errorf(codes.NotFound, "hash not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get hash: %v", err)
	}

	// Если хеш найден, возвращаем его
	return &pb.HashResponse{Hash: hash}, nil
}

func (s *HashingService) CreateHash(ctx context.Context, req *pb.HashRequest) (*pb.HashResponse, error) {
	// Здесь вычисляется хеш SHA-256 от payload, который был передан в запросе
	hash := sha256.Sum256([]byte(req.Payload))
	// Здесь хеш, который является байтовым массивом, преобразуется в строку шестнадцатеричных символов
	hashString := fmt.Sprintf("%x", hash)

	// Здесь хеш hashString и соответствующий ему payload сохраняются в базе данных Redis
	err := s.redisClient.Set(ctx, hashString, req.Payload, 0).Err()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to save hash: %v", err)
	}

	// Если хеш успешно сохранен, функция возвращает ответ с хешем и nil в качестве ошибки
	return &pb.HashResponse{Hash: hashString}, nil
}
