# API

## Docker

Перед запуском `docker-compose` нужно создать `.env` файл с переменными окружения в корне проекта.
В корне лежит файл `.env.example` со всеми нужными переменными.

#### Запуск через docker-compose:

```bash
docker-compose up --build -d
```

При использовании `docker-compose` будут созданы три контейнера: `db`, `migrator` и `api`.

- `db` - поднимает PostgreSQL
- `migrator` - автоматически запускает миграции при старте
- `api` - запускает http сервер для API

## Тестирование

Для примера реализован только тест для сервиса `deparment`, который лежит здесь:
`internal/service/department/service_test.go`.

## Структура проекта

```
├───cmd
│   ├───api
│   │       main.go
│   └───migrator
│       │   main.go
│       └───migration
│               20260520090442_base.sql
└───internal
    ├───errors
    │       errors.go
    ├───handler                           # HTTP-хендлеры и роутинг
    │   │   router.go
    │   ├───department
    │   │   │   handler.go
    │   │   └───dto
    │   │           converter.go
    │   │           request.go
    │   │           response.go
    │   └───employee
    │       │   handler.go
    │       └───dto
    │               converter.go
    │               request.go
    │               response.go
    ├───model                             # Модели и валидация данных
    │       department.go
    │       department_delete_mode.go
    │       employee.go
    │       title.go
    ├───repository                        # Работа с БД
    │   ├───department
    │   │   │   repository.go
    │   │   └───record
    │   │           converter.go
    │   │           record.go
    │   └───employee
    │       │   repository.go
    │       └───record
    │               converter.go
    │               record.go
    └───service                           # Бизнес-логика
        ├───department
        │       service.go
        └───employee
                service.go
```