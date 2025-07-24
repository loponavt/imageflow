# 🖼️ ImageFlow

ImageFlow — это микросервис для приёма задач на обработку изображений (например, `resize`, `blur`) и асинхронного их выполнения с помощью worker pool. Поддерживается проверка статуса задач по ID.

## 📦 Возможности

- Отправка задач на обработку изображений
- Поддержка разных типов обработки (resize, blur)
- Хранение задач в PostgreSQL
- Обработка задач в фоне через worker pool
- Swagger-документация

## 🚀 Быстрый старт

### 1. Запуск через Docker Compose

```bash
make docker-compose-up
```

Остановить:

```bash
make docker-compose-down
```

### 2. Сборка и запуск вручную

```bash
make build
./imageflow
```

Или через Docker:

```bash
make docker-build
make docker-run
```

---

## 📖 API

### Upload задачи

```http
GET /api/v1/upload?filename=test.jpg&type=resize
```

Параметры:

- `filename` — имя изображения (будет использоваться внутри обработки)
- `type` — тип обработки (`resize`, `blur`, и т.д.)

### Проверка статуса

```http
GET /api/v1/status?id=task-id
```

---

## 📂 Структура

- `cmd/app/main.go` — точка входа
- `internal/usecase` — бизнес-логика (обработка задач, worker pool)
- `internal/repository/postgres` — реализация репозитория на GORM/PostgreSQL
- `internal/delivery` — HTTP-обработчики на Gin
- `internal/worker` — реализация worker pool и задач (`Job`)
- `internal/model` — структуры данных

---

## 🔧 Переменные окружения

Используется `config.Load()`, по умолчанию ожидается порт `:8080`, можно адаптировать `.env` при необходимости.

---

## 🧪 Swagger

Доступен по адресу:

```
http://localhost:8080/swagger/index.html
```

---

## 📋 Пример задачи

```json
{
  "id": "4f8b6c21-7f85-41d5-b9d2-2f5fd8baf03a",
  "filename": "cat.jpg",
  "type": "resize",
  "status": "done"
}
```

---

## 🛠️ TODO

- Загрузка и сохранение реальных изображений (s3 minio)
- Расширение типов обработки (например, grayscale, watermark)
- UI для отправки и отслеживания задач
