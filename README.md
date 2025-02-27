# Go Expression Parser

## Введение

#### Проект представляет собой http-сервер на языке golang, который позволяет пользователю считать арифметические выражения

#### Поддерживаемые операции: +, -, /, *, ^ (возведение в степень)

----

## Способы запуска

- #### Склонировать проект `git clone https://github.com/Kripipastt/go-expression-parser`

1. ### Docker

- #### Необходимо установить [docker](https://www.docker.com/products/docker-desktop/)
- #### Перейти в директорию проекта
- #### Запустить `docker compose up --build`

[//]: # (2. ### Go)

[//]: # (- #### Необходимо установить [golang]&#40;https://go.dev/dl/&#41;)

[//]: # (- #### Перейти в директорию проекта)

[//]: # (- #### Скачать нужные пакеты `go mod download`)

[//]: # (- #### Запустить сервер `go run cmd/main.go`)

## Swagger

#### После запуска сервера на `http://localhost:8080/swagger/index.html` будет доступен swagger, в котором расписаны все имеющиеся эндпоинты

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
  "id": "53aa35c8-e659-44b2-882f-f6056e443c99"
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

### Get `/api/v1/expressions`

#### Request example:

```bash
curl -X GET --location 'localhost:8080/api/v1/expressions'
```

#### Response:

```json
{
  "expressions": [
    {
      "id": "53aa35c8-e659-44b2-882f-f6056e443c99",
      "expression": "2 + 2 * 2",
      "status": "finish",
      "result": 6
    }
  ]
}
```

#### 4 Вида статусов:

- create - выражение отправлено на сервер и ожидает своей очереди
- pending - выражение высчитывается
- finish - выражение успешно посчитано и получен ответ
- reject - выражение было отменено из-за определенных причин (например, деление на 0)

### Get `/api/v1/expressions/<expression_id>`

#### Request example:

```bash
curl -X GET --location 'localhost:8080/api/v1/expressions/53aa35c8-e659-44b2-882f-f6056e443c99'
```

#### Response:

```json
{
  "expression": {
    "id": "53aa35c8-e659-44b2-882f-f6056e443c99",
    "expression": "2 + 2 * 2",
    "status": "finish",
    "result": 6
  }
}
```

### Get `/internal/task`

#### Request example:

```bash
curl -X GET --location 'localhost:8080/internal/task'
```

#### Response:

```json
{
  "task": {
    "id": "53aa35c8-e659-44b2-882f-f6056e443c99",
    "arg1": 2,
    "arg2": 2,
    "operation": "+",
    "operation_time": 3
  }
}
```

### Post `/internal/task`

#### Request example:

```bash
curl --location 'localhost/internal/task' \
  --header 'Content-Type: application/json' \
  --data '{
    "id": 1,
    "result": 4
  }'
```

## Тестирование

#### Введите `go test ./...` для запуска тестов