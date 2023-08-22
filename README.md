# Пример gRPC клиента и сервера

## [Установка зависимостей](https://grpc.io/docs/languages/go/quickstart/#prerequisites)

### Запуск
1) ```go run cmd/server/main.go```
2) ```go run cmd/client/main.go```

### Генерация .pb.go файлов
``` protoc --go_out=. --go_opt=paths=source_relative \ --go-grpc_out=. --go-grpc_opt=paths=source_relative \ helloworld/helloworld.proto ```