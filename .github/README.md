# PetGoApi - Todo REST API

[![CI/CD Pipeline](https://github.com/stavagg/petGoApi/actions/workflows/ci.yml/badge.svg)](https://github.com/stavagg/petGoApi/actions/workflows/ci.yml)
[![PR Check](https://github.com/stavagg/petGoApi/actions/workflows/pr-check.yml/badge.svg)](https://github.com/stavagg/petGoApi/actions/workflows/pr-check.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/stavagg/petGoApi)](https://goreportcard.com/report/github.com/stavagg/petGoApi)
[![Coverage](https://codecov.io/gh/stavagg/petGoApi/branch/main/graph/badge.svg)](https://codecov.io/gh/stavagg/petGoApi)
[![Docker Pulls](https://img.shields.io/docker/pulls/stavagg/petgoapi)](https://hub.docker.com/r/stavagg/petgoapi)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/go-1.23-blue.svg)](https://golang.org)

🚀 **REST API** для управления Todo задачами, построенный на Go с применением принципов Clean Architecture и современных практик разработки.

## ✨ Основные возможности

- ✅ **Clean Architecture** - четкое разделение слоев (Handler → Service → Repository → Database)
- ✅ **RESTful API** - полный CRUD функционал с правильными HTTP методами и статус кодами
- ✅ **PostgreSQL** - надежная реляционная база данных с GORM ORM
- ✅ **Docker** - полная контейнеризация приложения и базы данных
- ✅ **Unit тесты** - покрытие ключевой бизнес-логики с отчетами
- ✅ **CI/CD Pipeline** - автоматическое тестирование и деплой через GitHub Actions
- ✅ **Валидация данных** - проверка входящих запросов на уровне модели
- ✅ **Error Handling** - правильная обработка ошибок и информативные сообщения
- ✅ **CORS поддержка** - готовность для фронтенд интеграции
- ✅ **Статистика** - аналитика по задачам и прогрессу выполнения

## 🏗️ Архитектура

Проект построен с использованием **Clean Architecture** принципов:

┌─────────────┐ ┌─────────────┐ ┌─────────────┐ ┌─────────────┐
│ Handler │───▶│ Service │───▶│ Repository │───▶│ Database │
│ (HTTP) │ │ (Business) │ │ (Data) │ │ (Storage) │
└─────────────┘ └─────────────┘ └─────────────┘ └─────────────┘
▲ ▲ ▲


### Слои приложения:

- **Handler** (`internal/handler/`) - HTTP обработчики, маршрутизация, валидация запросов
- **Service** (`internal/service/`) - бизнес-логика, правила валидации, координация операций  
- **Repository** (`internal/repository/`) - работа с базой данных, SQL запросы через GORM
- **Model** (`internal/model/`) - структуры данных, запросы и ответы API

## 📝 API Endpoints

### Todo Management

| Метод | Путь | Описание | Тело запроса |
|-------|------|----------|--------------|
| `POST` | `/api/v1/todos` | Создать новую задачу | `{"title": "string", "description": "string"}` |
| `GET` | `/api/v1/todos` | Получить все задачи | - |
| `GET` | `/api/v1/todos?completed=true` | Фильтр по статусу | - |
| `GET` | `/api/v1/todos/:id` | Получить задачу по ID | - |
| `PUT` | `/api/v1/todos/:id` | Обновить задачу | `{"title": "string", "description": "string", "completed": boolean}` |
| `DELETE` | `/api/v1/todos/:id` | Удалить задачу | - |

### Additional Features

| Метод | Путь | Описание |
|-------|------|----------|
| `POST` | `/api/v1/todos/:id/toggle` | Переключить статус выполнения |
| `GET` | `/api/v1/todos/stats` | Статистика по задачам |
| `GET` | `/health` | Проверка работоспособности API |
| `GET` | `/` | Информация о доступных endpoints |

### Примеры запросов

Создание задачи
curl -X POST http://localhost:8080/api/v1/todos
-H "Content-Type: application/json"
-d '{"title":"Изучить Go","description":"Создать REST API проект"}'

Получение всех задач
curl http://localhost:8080/api/v1/todos

Обновление задачи
curl -X PUT http://localhost:8080/api/v1/todos/1
-H "Content-Type: application/json"
-d '{"title":"Изучить Go - ЗАВЕРШЕНО","completed":true}'

Статистика
curl http://localhost:8080/api/v1/todos/stats


## 🚀 Быстрый старт

### Docker Compose (рекомендуется)

Клонировать репозиторий
git clone https://github.com/stavagg/petGoApi.git
cd petGoApi

Запустить приложение и базу данных
docker-compose up --build

API будет доступен на http://localhost:8080


### Локальная разработка

1. Запустить только PostgreSQL в Docker
docker-compose up -d db

2. Установить зависимости Go
go mod download

3. Настроить переменные окружения
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASS=password
export DB_NAME=mydb
export PORT=:8080

4. Запустить приложение
go run cmd/api/main.go

### Переменные окружения

| Переменная | Описание | Значение по умолчанию |
|------------|----------|-----------------------|
| `PORT` | Порт для HTTP сервера | `:8080` |
| `DB_HOST` | Хост PostgreSQL | `localhost` |
| `DB_PORT` | Порт PostgreSQL | `5432` |
| `DB_USER` | Пользователь БД | `postgres` |
| `DB_PASS` | Пароль БД | `password` |
| `DB_NAME` | Название базы данных | `mydb` |

## 🧪 Тестирование

### Запуск тестов

Запустить все тесты
go test ./...

Тесты с подробным выводом
go test ./... -v

Тесты с покрытием
go test ./... -cover

Генерация HTML отчета покрытия
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out


### Текущее покрытие

- **Service слой**: ~54% - основная бизнес-логика покрыта тестами
- **Handler слой**: ~19% - базовые HTTP тесты реализованы  
- **Общее покрытие**: ~25% - стабильная база для дальнейшего развития

### Структура тестов

internal/
├── service/
│ ├── todo.go
│ └── todo_test.go # Unit тесты бизнес-логики
├── handler/
│ ├── todo.go
│ └── todo_test.go # HTTP интеграционные тесты
└── repository/mocks/ # Моки для тестирования
└── todo_repository_mock.go


## 🏗️ Технологический стек

### Backend
- **Go 1.23** - язык программирования
- **Gin** - высокопроизводительный HTTP веб-фреймворк
- **GORM** - ORM библиотека для работы с базой данных
- **PostgreSQL 15** - реляционная база данных

### Инфраструктура  
- **Docker & Docker Compose** - контейнеризация приложения
- **GitHub Actions** - CI/CD pipeline для автоматического тестирования

### Тестирование
- **Testify** - библиотека для написания тестов и создания моков
- **httptest** - HTTP тестирование для endpoints

### Инструменты разработки
- **Air** - live reload для разработки (опционально)
- **golangci-lint** - статический анализ кода
- **goimports** - форматирование и организация импортов

## 📊 CI/CD Pipeline

Каждый commit и Pull Request автоматически проверяется через GitHub Actions:

### Этапы проверки:
1. **Тестирование** - запуск всех unit и integration тестов
2. **Покрытие** - проверка минимального порога покрытия кода
3. **Линтинг** - статический анализ качества кода
4. **Сборка Docker** - проверка успешной контейнеризации
5. **Деплой** - автоматическая публикация образов (на main ветке)

### Статус проекта
- ✅ **Тесты**: автоматически проходят при каждом изменении
- ✅ **Качество кода**: соответствует стандартам Go
- ✅ **Docker**: готов к деплою в любое окружение

## 📁 Структура проекта

petGoApi/
├── cmd/api/
│ └── main.go # Точка входа приложения
├── internal/
│ ├── config/
│ │ └── config.go # Конфигурация приложения
│ ├── handler/
│ │ ├── todo.go # HTTP обработчики
│ │ └── todo_test.go # Тесты обработчиков
│ ├── service/
│ │ ├── todo.go # Бизнес-логика
│ │ ├── todo_test.go # Unit тесты
│ │ └── mocks/ # Моки для тестирования
│ ├── repository/
│ │ ├── todo.go # Работа с БД
│ │ └── mocks/ # Моки репозитория
│ └── model/
│ └── todo.go # Модели данных
├── .github/workflows/
│ └── ci.yml # CI/CD конфигурация
├── docker-compose.yml # Docker Compose для разработки
├── Dockerfile # Образ приложения
├── go.mod # Зависимости Go
└── README.md # Документация проекта

