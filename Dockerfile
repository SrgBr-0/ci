# Используем официальный образ Go
FROM golang:1.23-alpine

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum для установки зависимостей (если они есть)
COPY go.mod go.sum ./

# Устанавливаем зависимости
RUN go mod tidy

# Копируем исходный код
COPY . .

# Сборка приложения
RUN go build -o main .

# Открываем порт 8080 для доступа к приложению
EXPOSE 8080

# Запускаем приложение
CMD ["./main"]
