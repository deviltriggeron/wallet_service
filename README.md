# Wallet Service

Простой сервис для управления кошельками с REST API на Go, использованием PostgreSQL и Docker.  
Позволяет создавать депозиты, выполнять снятия средств, пополнение и получать баланс кошелька.  

## 🛠 Стек технологий

- **Backend:** Golang  
- **База данных:** PostgreSQL  
- **Фреймворк для роутинга:** Gorilla Mux  
- **Контейнеризация:** Docker + Docker Compose  
- **Миграции:** SQL скрипты (`db/model.sql`)  
- **Тестирование:** стандартный пакет `testing`  


## 🐳 Запуск

- Создайте файл `config.env`:

```env
POSTGRES_USER=wallet_user
POSTGRES_PASSWORD=wallet_pass
POSTGRES_DB=wallet_db
POSTGRES_HOST=db
POSTGRES_PORT=5432
SERVER_PORT=8080
```
- Запустить сервис с Docker:

```
    docker-compose up
```
- Postgres автоматически создаст таблицу wallets из db/model.sql.
- Go-сервер будет доступен на http://localhost:8080.

## 👨🏻‍💻 Проверка работы

- Для запуска тестов:

```
bash test_run.sh
```
- Перед запуском тестов убедитесь что Go-сервер запущен

- Примеры запросов:

```
curl -X POST http://localhost:8080/api/v1/createWallets \
  -H "Content-Type: application/json" \
  -d '{"amount": 1000}'

curl -X GET http://localhost:8080/api/v1/wallets/

curl -X POST http://localhost:8080/api/v1/wallet \
  -H "Content-Type: application/json" \
  -d '{
        "walletId": "f6336cb8-c129-4ac2-a141-5e86eb34ac27",
        "operationType": "DEPOSIT",
        "amount": 1000
      }'
```