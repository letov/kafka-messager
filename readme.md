# Goka messenger

## Подготовка окружения

```bash
make up
```
```bash
make init
```

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
