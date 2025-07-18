# Документация для сервиса управления подписками

## Описание сервиса

Сервис представляет собой REST API для управления онлайн-подписками пользователей. Основные функции:
- CRUDL-операции (Create, Read, Update, Delete, List) над подписками
- Расчет суммарной стоимости подписок за указанный период с возможностью фильтрации
- Хранение данных в PostgreSQL
- Логирование операций
- Swagger-документация API
- Запуск через Docker Compose

Каждая подписка содержит:
- Название сервиса
- Стоимость месячной подписки (целое число рублей)
- UUID пользователя
- Дата начала подписки (месяц и год)
- Опциональную дату окончания подписки

## Makefile документация

### Переменные
- `BINARY_NAME` - имя исполняемого файла (по умолчанию `myapp`)
- `DOCKER_IMAGE` - имя Docker-образа (по умолчанию `myapp-image`)
- `GO_BUILD_ENV` - переменные окружения для сборки Go (отключен CGO)
- `GO_FILES` - список всех Go-файлов в проекте (исключая vendor)

### Цели (targets)

#### Основные цели
- `all` - цель по умолчанию, выполняет `build`
- `build` - собирает основной бинарный файл из `cmd/main.go`

#### Очистка
- `clean` - удаляет скомпилированные бинарные файлы и выполняет `go clean`

#### Docker-операции
- `docker-build` - собирает Docker-образ через `docker compose build`
- `docker-up` - запускает сервис через `docker compose up -d`
- `docker-down` - останавливает сервис через `docker compose down`

#### Кросс-компиляция
- `build-linux-arm` - собирает версию для Linux ARM64
- `build-windows` - собирает версию для Windows (amd64)

### Использование
1. Для локальной разработки:
   ```bash
   make build
   ./myapp
   ```

2. Для сборки и запуска через Docker:
   ```bash
   make docker-build
   make docker-up
   ```

3. Для остановки сервиса:
   ```bash
   make docker-down
   ```

4. Для очистки:
   ```bash
   make clean
   ```

## Конфигурация
Конфигурационные параметры вынесены в `.env` файл, который должен быть создан на основе примера перед запуском сервиса.

## API Документация
# Project: svc-subman

## End-point: create subscribe
### Method: POST
>```
>{{local}}/subscribe
>```
### Body (**raw**)

```json
{
  "service_name": "Netflix Premium",
  "price": 15,
  "user_id": "a1b2c3dx-e5f6-7890-g1h2-i3j4k5l6m7n2",
  "start_date": "01-2024",
  "end_date": "12-2024"
}
```

### Response: 200
```json
{
    "status": "success",
    "data": "30d548cb-caa5-4866-bdbe-22c7e93f745a"
}
```

### Response: 400
```json
{
    "status": "error",
    "error": "Key: 'CreateSubscriptionRequest.Price' Error:Field validation for 'Price' failed on the 'required' tag",
    "details": 400
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: delete subscriber
### Method: DELETE
>```
>{{local}}/subscribe
>```
### Body (**raw**)

```json
{
    "ids": [
        "688164ab-354a-4908-ae66-5a3d53908361",
        "1ba5eb71-94b9-4f33-a2f8-77b85f52c9c9"
        
    ]
}
```

### Response: 200
```json
{
    "status": "success",
    "data": "All subscribers have been deleted"
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: agree subscribe
### Method: GET
>```
>{{local}}/subscribe/cost?user_id=a1b2c3dx-e5f6-7890-g1h2-i3j4k5l6m7n2&service_name=Netflix Premium12&start_date=01-2024&end_date=02-2025
>```
### Query Params

|Param|value|
|---|---|
|user_id|a1b2c3dx-e5f6-7890-g1h2-i3j4k5l6m7n2|
|service_name|Netflix Premium12|
|start_date|01-2024|
|end_date|02-2025|


### Response: 200
```json
{
    "status": "success",
    "data": 100000
}
```

### Response: 400
```json
{
    "status": "error",
    "error": "failed to get ent_date",
    "details": 400
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: get subscribes
### Method: GET
>```
>{{local}}/subscribes?user_id&service_name=Netflix Premium12&limit=10&offset=0
>```
### Query Params

|Param|value|
|---|---|
|user_id|null|
|service_name|Netflix Premium12|
|limit|10|
|offset|0|


### Response: 200
```json
{
    "status": "success",
    "data": [
        {
            "id": "af7a4ff5-6ce9-48d8-bf4e-ea28b005fadb",
            "service_name": "Netflix Premium",
            "price": 15,
            "user_id": "a1b2c3dx-e5f6-7890-g1h2-i3j4k5l6m7n2",
            "start_date": "2024-01-01T00:00:00Z",
            "end_date": "2024-12-01T00:00:00Z"
        },
        {
            "id": "fd57aa6f-88a7-4b04-9477-2bdeb3f3f20d",
            "service_name": "Netflix Premium",
            "price": 15,
            "user_id": "a1b2c3dx-e5f6-7890-g1h2-i3j4k5l6m7n2",
            "start_date": "2024-01-01T00:00:00Z",
            "end_date": "2024-12-01T00:00:00Z"
        },
        {
            "id": "a470dc53-2658-428f-8ac7-0c91f1e77ae2",
            "service_name": "Netflix Premium",
            "price": 15,
            "user_id": "a1b2c3dx-e5f6-7890-g1h2-i3j4k5l6m7n2",
            "start_date": "2024-01-01T00:00:00Z",
            "end_date": "2024-12-01T00:00:00Z"
        },
        {
            "id": "6eed7d72-8a00-49dc-b5a0-1237959cfef7",
            "service_name": "Netflix Premium",
            "price": 15,
            "user_id": "a1b2c3dx-e5f6-7890-g1h2-i3j4k5l6m7n2",
            "start_date": "2024-01-01T00:00:00Z",
            "end_date": "2024-12-01T00:00:00Z"
        },
        {
            "id": "0577ae90-6df1-4d48-833b-a0cc821d2bef",
            "service_name": "Netflix Premium",
            "price": 15,
            "user_id": "a1b2c3dx-e5f6-7890-g1h2-i3j4k5l6m7n2",
            "start_date": "2024-01-01T00:00:00Z",
            "end_date": "2024-12-01T00:00:00Z"
        },
        {
            "id": "cd904cbf-d629-4b2f-9b4c-b08570b22bbd",
            "service_name": "Netflix Premium",
            "price": 15,
            "user_id": "a1b2c3dx-e5f6-7890-g1h2-i3j4k5l6m7n2",
            "start_date": "2024-01-01T00:00:00Z",
            "end_date": "2024-12-01T00:00:00Z"
        },
        {
            "id": "411fe711-e2bd-4069-b070-f8078dc6f931",
            "service_name": "Netflix Premium",
            "price": 15,
            "user_id": "a1b2c3dx-e5f6-7890-g1h2-i3j4k5l6m7n2",
            "start_date": "2024-01-01T00:00:00Z",
            "end_date": "2024-12-01T00:00:00Z"
        },
        {
            "id": "d5760810-5ae6-4eb4-8947-92eed1da3eaa",
            "service_name": "Netflix Premium",
            "price": 15,
            "user_id": "a1b2c3dx-e5f6-7890-g1h2-i3j4k5l6m7n2",
            "start_date": "2024-01-01T00:00:00Z",
            "end_date": "2024-12-01T00:00:00Z"
        },
        {
            "id": "31110cd2-0ed0-482f-9002-efafd17c003a",
            "service_name": "Netflix Premium",
            "price": 15,
            "user_id": "a1b2c3dx-e5f6-7890-g1h2-i3j4k5l6m7n2",
            "start_date": "2024-01-01T00:00:00Z",
            "end_date": "0001-01-01T00:00:00Z"
        },
        {
            "id": "37df5f72-8cea-4b81-9b82-b8997fc277f5",
            "service_name": "Netflix Premium",
            "price": 15,
            "user_id": "a1b2c3dx-e5f6-7890-g1h2-i3j4k5l6m7n2",
            "start_date": "0001-01-01T00:00:00Z",
            "end_date": "0001-01-01T00:00:00Z"
        }
    ],
    "meta": 200
}
```

### Response: 400
```json
{
    "status": "error",
    "error": "failed to get offset",
    "details": 400
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: get subscribe
### Method: GET
>```
>{{local}}/subscribe/df73e03e-ea64-415e-a1e0-222eeb137a1c
>```
### Response: 200
```json
{
    "status": "success",
    "data": {
        "id": "df73e03e-ea64-415e-a1e0-222eeb137a1c",
        "service_name": "Netflix Premium",
        "price": 15,
        "user_id": "a1b2c3dx-e5f6-7890-g1h2-i3j4k5l6m7n2",
        "start_date": "2024-01-01T00:00:00Z",
        "end_date": "2024-12-01T00:00:00Z"
    },
    "meta": 200
}
```

### Response: 400
```json
{
    "status": "error",
    "error": "params id is not valid uuid",
    "details": 400
}
```



