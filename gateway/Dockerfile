# Используем официальный образ Go как базовый
FROM golang:1.22

# Устанавливаем рабочую директорию в контейнере
WORKDIR /app

# Копируем go.mod и go.sum в рабочую директорию
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем исходный код в рабочую директорию
COPY . .

# Собираем приложение
RUN go build -o gateway ./cmd/gateway/main.go

# Запускаем приложение
CMD ["./gateway"]
