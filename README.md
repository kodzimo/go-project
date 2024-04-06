# Go Portfolio Project

## Описание

Это проект, который включает в себя два микросервиса: Gateway Service и Hashing Service. Gateway Service принимает REST-запросы от пользователей и перенаправляет их в Hashing Service с использованием gRPC. Hashing Service обрабатывает эти запросы, выполняя операции с хешами.

## Структура проекта

go-portfolio-project
│
├── gateway
│   ├── cmd
│   │   └── gateway
│   │       └── main.go
│   ├── internal
│   │   └── gateway-service.go
│   └── Dockerfile
│
├── hashing
│   ├── cmd
│   │   └── hashing
│   │       └── main.go
│   ├── internal
│   │   ├── hashing-service.go
│   │   └── grpc-server.go
│   ├── storage
│   │   └── redis.go
│   └── Dockerfile
│
└── shared
    └── proto
        ├── hashing.proto
        ├── hashing.pb.go
        └── hashing_grpc.pb.go

## Установка и запуск

1. Установите Docker и Docker Compose на вашем компьютере.
2. Клонируйте этот репозиторий на ваш компьютер.
3. Перейдите в каталог проекта.
4. Запустите `docker-compose up -d`.

## Использование

Вы можете использовать `Makefile` для выполнения запросов к вашему `hashing-service` через `gateway`. Вот как это сделать:

1. Откройте терминал.
2. Перейдите в каталог вашего проекта.
3. Выполните одну из следующих команд:

   - Для выполнения запроса `checkhash`:
     ```bash
     make checkhash
     ```
   - Для выполнения запроса `gethash`:
     ```bash
     make gethash
     ```
   - Для выполнения запроса `createhash`:
     ```bash
     make createhash
     ```

Каждая из этих команд отправляет HTTP POST запрос на соответствующий эндпоинт вашего `gateway` сервиса (`localhost:8080/checkhash`, `localhost:8080/gethash` или `localhost:8080/createhash`), который затем перенаправляет запрос к `hashing-service`.

## Лицензия

Этот проект лицензирован под MIT License - см. файл LICENSE.md для подробностей.

## Доработка

- Проверить работоспособность
- Доделать тесты
- Сделать всё по требованиям go kit
