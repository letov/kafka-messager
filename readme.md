# Goka messenger

## Архитектура проекта
![C2!](./docs/c2.png "C2")

## Структура проекта

    .
    ├── cmd                  # Основное приложение
    ├── docker               # KSQL скрипты и параметры коннектора
    ├── internal             # Внутренний код
    │   ├── application      # Application layer  
    │   │   ├── app         
    │   │   └── dto
    │   ├── domain           # Domain layer
    │   ├── infra            # Infrastructure layer
    │   │   ├── config        
    │   │   ├── db         
    │   │   ├── di
    │   │   ├── logger
    │   │   ├── msg   
    │   │   ├── repo   
    │   │   └── repo
    │   └── test         
    └── ...

## Запуск проекта

Создать в корне .env

```
DATABASE_DSN=
MSG_TOPIC=
MSG_FILTERED_BLOCK_USERS_TOPIC=
MSG_FILTERED_TOPIC=
SCHEMA_REGISTRY_URL=
KAFKA_BROKERS=
KAFKA_SESSION_TIMEOUT_MS=
KAFKA_AUTO_OFFSET_RESET=
KAFKA_CONSUMER_PULL_TIMEOUT_MS=
KAFKA_ACKS=
```

Подготовка окружения и запуск теста

```bash
make up
```
```bash
make init
```
```bash
go test -v ./...
```

## Основные структуры
* msg.Codec - кодер для сериализации/десериализации данных goka процессоров
* msg.MsgEmitter - эмиттер для отправки новых сообщений в kafka
* msg.Processor - процессор goka, содержит колбек для маскировки запрещенных слов
* msg.Receiver - слушатель топика kafka, содержащий отфильтрованные сообщения (без забаненных пользователей и запрещенных слов)
* msg.Schema - вспомогательные функции для schema registry