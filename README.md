# Task API Service

HTTP API для управления долгими I/O bound задачами (выполнение 3-5 минут).  
Задачи хранятся в памяти сервера без использования внешних зависимостей.

---

## 🚀 Запуск сервера

### Требования
- Go 1.20+
- Для тестов: Postman или curl

### Установка и запуск
```bash
# Клонировать репозиторий
git clone https://github.com/swagxx/tasktoto.git
cd task-api

# Запустить сервер
go run cmd/server/main.go
```
### Тестирование API

- POST /tasks Создать задачу

- GET	/tasks/{id}	Проверить статус

- DELETE	/tasks/{id}	Удалить задачу

### Через cURL
```bash
# Создать задачу
curl -X POST http://localhost:8080/tasks -H "Content-Type: application/json"

# Проверить статус
curl http://localhost:8080/tasks/1739320423040000000

# Удалить задачу
curl -X DELETE http://localhost:8080/tasks/1739320423040000000
```

### Через Postman 
```bash
  http://localhost:8080
```
#### Метод POST
```bash
  #выбираете метод Post и нажимаете на Send
http://localhost:8080
Результат:
{
  "id": "1739320423040000000",
  "status": "pending",
  "created_at": "2023-10-25T12:00:00Z"
}
```


#### Метод GET 
```bash
http://localhost:8080/{id}
# тот который получили в результате post запроса
```

#### Метод DELETE
```bash
http://localhost:8080/{id}
```

# СПАСИБО ЗА ВНИМАНИЕ