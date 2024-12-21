# Go Expression Parser

## Введение

#### Проект представляет собой небольшой http-сервер на языке golang, который позволяет пользователю считать арифметические выражения

#### Поддерживаемые операции: +, -, /, *, ^ (возведение в степень)

----

## Способы запуска

1. ### Docker

- #### Необходимо установить [docker](https://www.docker.com/products/docker-desktop/)
- #### Перейти в директорию проекта
- #### Ввести команду `docker compose up --build`

2. ### Go

- #### Необходимо установить [golang](https://go.dev/dl/)
- #### Перейти в директорию проекта
- #### Ввести команду `go run cmd/main.go`

## Эндпоинты

### Get `/ping`
```bash
curl -X GET --location localhost:8080/ping
```

#### Возвращает `pong`, необходим для проверки старта сервера

### Post `/api/v1/calculate`

#### Request example:

```bash
curl -X POST --location 'localhost:8080/api/v1/calculate' \
  --header 'Content-Type: application/json' \
  --data '{"expression": "2 + 2 * 2"}'
```

#### Response:  
Status 200
```json
{
  "result": 6
}
```

Если было введено некорректное выражение:  
Status 422
```json
{
  "error": "Expression is not valid"
}
```

Если произошла непредвиденная ошибка:  
Status 500
```json
{
  "error": "Internal server error"
}
```

## Тестирование
#### Введите `go test ./...` для запуска тестов