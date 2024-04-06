package gateway

import (
	"context"
	"io"
	"net/http"

	pb "final-project-kodzimo-shared/proto"
)

/*
В этом коде GatewayService содержит клиент gRPC, который будет использоваться для взаимодействия
с Hashing Service. Затем у нас есть три обработчика HTTP: CheckHashHandler, GetHashHandler и CreateHashHandler,
которые соответствуют трем методам в вашем сервисе Hashing.

Каждый из этих обработчиков будет принимать HTTP-запрос, извлекать необходимые данные из запроса, вызывать
соответствующий метод на клиенте gRPC, а затем отправлять ответ обратно клиенту.
*/

type GatewayService struct {
	HashingClient pb.HashingClient
}

/*
Конечно, вот пример HTTP POST запроса, который вы можете использовать для тестирования обработчика `CheckHashHandler`:

```http
POST /checkhash HTTP/1.1
Host: localhost:8080
Content-Type: text/plain
Content-Length: 13

Hello, world!
```

В этом примере `/checkhash` - это путь, по которому ваш обработчик `CheckHashHandler` прослушивает запросы.
`localhost:8080` - это адрес и порт вашего сервера (замените их на реальные значения, если они отличаются).
`Hello, world!` - это полезная нагрузка запроса, которую вы хотите проверить.
*/

func (g *GatewayService) CheckHashHandler(w http.ResponseWriter, r *http.Request) {
	// Проверяем, что метод запроса - POST.
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Извлекаем полезную нагрузку из тела запроса.
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// Создаем и заполняем HashRequest.
	req := &pb.HashRequest{
		Payload: string(body),
	}

	// Вызываем метод CheckHash на клиенте gRPC.
	res, err := g.HashingClient.CheckHash(context.Background(), req)
	if err != nil {
		http.Error(w, "Error calling CheckHash: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Возвращаем полученный хеш обратно клиенту.
	w.Write([]byte(res.Hash))
}

/*
```http
POST /gethash HTTP/1.1
Host: localhost:8080
Content-Type: text/plain
Content-Length: 13

Hello, world!
```

Этот обработчик будет принимать HTTP-запрос, извлекать полезную нагрузку из запроса, вызывать метод GetHash
на клиенте gRPC, а затем отправлять ответ обратно клиенту.
*/

func (g *GatewayService) GetHashHandler(w http.ResponseWriter, r *http.Request) {
	// Проверяем, что метод запроса - POST.
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Извлекаем полезную нагрузку из тела запроса.
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// Создаем и заполняем HashRequest.
	req := &pb.HashRequest{
		Payload: string(body),
	}

	// Вызываем метод GetHash на клиенте gRPC.
	res, err := g.HashingClient.GetHash(context.Background(), req)
	if err != nil {
		http.Error(w, "Error calling GetHash: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Возвращаем полученный хеш обратно клиенту.
	w.Write([]byte(res.Hash))
}

/*
```http
POST /createhash HTTP/1.1
Host: localhost:8080
Content-Type: text/plain
Content-Length: 13

Hello, world!
```
Этот обработчик будет принимать HTTP-запрос, извлекать полезную нагрузку из запроса, вызывать метод CreateHash
на клиенте gRPC, а затем отправлять ответ обратно клиенту.
*/

func (g *GatewayService) CreateHashHandler(w http.ResponseWriter, r *http.Request) {
	// Проверяем, что метод запроса - POST.
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Извлекаем полезную нагрузку из тела запроса.
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// Создаем и заполняем HashRequest.
	req := &pb.HashRequest{
		Payload: string(body),
	}

	// Вызываем метод CreateHash на клиенте gRPC.
	res, err := g.HashingClient.CreateHash(context.Background(), req)
	if err != nil {
		http.Error(w, "Error calling CreateHash: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Возвращаем полученный хеш обратно клиенту.
	w.Write([]byte(res.Hash))
}
