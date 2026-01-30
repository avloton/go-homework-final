.PHONY: help build run populate clean docker-build docker-up docker-down docker-stop

# Переменные
BINARY_NAME=server
MAIN_PATH=cmd/main.go
MODULE_NAME=mywebsite
DOCKER_IMAGE=mybakery:latest

# Цвета для вывода (опционально)
GREEN=\033[0;32m
YELLOW=\033[1;33m
NC=\033[0m # No Color

help: ## Показать справку по командам
	@echo "$(GREEN)Доступные команды:$(NC)"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  $(YELLOW)%-15s$(NC) %s\n", $$1, $$2}'

build: ## Скомпилировать приложение
	@echo "$(GREEN)Компиляция приложения...$(NC)"
	@go build -o $(BINARY_NAME) $(MAIN_PATH)
	@echo "$(GREEN)Готово! Исполняемый файл: $(BINARY_NAME)$(NC)"

run: ## Запустить приложение
	@echo "$(GREEN)Запуск приложения...$(NC)"
	@go run $(MAIN_PATH)

populate: ## Инициализировать базу данных
	@echo "$(GREEN)Инициализация базы данных...$(NC)"
	@go run $(MAIN_PATH) populate

docker-build: ## Собрать Docker образ
	@echo "$(GREEN)Сборка Docker образа...$(NC)"
	@docker build -t $(DOCKER_IMAGE) .
	@echo "$(GREEN)Docker образ собран: $(DOCKER_IMAGE)$(NC)"

docker-up: ## Запустить приложение через Docker Compose
	@echo "$(GREEN)Запуск через Docker Compose...$(NC)"
	@docker compose up -d
	@echo "$(GREEN)Приложение запущено на http://localhost:8080$(NC)"

docker-down: ## Остановить и удалить Docker контейнеры
	@echo "$(GREEN)Остановка и удаление контейнеров ...$(NC)"
	@docker compose down

docker-stop: ## Остановить Docker контейнеры
	@echo "$(GREEN)Остановка контейнеров ...$(NC)"
	@docker compose stop

docker-logs: ## Показать логи Docker контейнеров
	@docker compose logs -f

docker-restart: docker-stop docker-up ## Перезапустить Docker Compose

dev: ## Запустить в режиме разработки (с populate)
	@echo "$(GREEN)Запуск в режиме разработки...$(NC)"
	@go run $(MAIN_PATH) populate
	@go run $(MAIN_PATH)

clean: ## Очистить скомпилированные файлы
	@echo "$(GREEN)Очистка...$(NC)"
	@rm -f $(BINARY_NAME) $(BINARY_NAME).exe
	@echo "$(GREEN)Очистка завершена$(NC)"
