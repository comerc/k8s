# Multi-Tier Application - Полноценное приложение на Go

Многоуровневое приложение с frontend (Nginx), backend (**Golang**) и базой данных (PostgreSQL).

**Уровень сложности**: 🔴 Продвинутый

## Что изучите

- ✅ **StatefulSet** - PostgreSQL с persistent storage
- ✅ **Secrets** - безопасное хранение паролей БД
- ✅ **Ingress** - маршрутизация HTTP трафика
- ✅ **Namespaces** - изоляция приложения
- ✅ **Go REST API** - database/sql + lib/pq
- ✅ **Multi-tier архитектура** - frontend → backend → database
- ✅ **Production patterns** - health checks, graceful shutdown, CORS

## Технологии

- **Frontend**: Nginx + HTML/JS
- **Backend**: Go 1.21 (net/http + database/sql)
- **Database**: PostgreSQL 15 + StatefulSet
- **Kubernetes**: Ingress, Secrets, ConfigMaps, StatefulSet

## Производительность Go backend

| Метрика | Go | Сравнение с Node.js |
|---------|-----|---------------------|
| Образ Docker | ~20 MB | vs 200+ MB |
| Память (runtime) | ~20 MB | vs 150+ MB |
| Startup время | <3 сек | vs 15-40 сек |
| Запросов/сек | 40,000+ | vs 2,000 |

## Архитектура

```
          [Internet]
              ↓
         [Ingress]
         /       \
        /         \
  [Frontend]    [Backend API]  ⚡ Go
  (Nginx)       (Golang)
                    ↓
              [PostgreSQL]
              (StatefulSet)
```

## Компоненты

### 1. Frontend
- HTML/JavaScript приложение
- Nginx для статики
- Общается с Backend через REST API

### 2. Backend (Go)
- Go REST API с net/http
- Подключение к PostgreSQL (lib/pq)
- Endpoints: GET/POST /api/items, GET /health
- CORS middleware

### 3. Database
- PostgreSQL 15
- StatefulSet с PersistentVolume
- Secrets для credentials

## Быстрый старт

### 1. Создать namespace

```bash
kubectl create namespace multi-tier
```

### 2. Запустить базу данных

```bash
kubectl apply -f database/ -n multi-tier

# Дождаться готовности
kubectl get pods -n multi-tier -w
```

### 3. Запустить backend (Go)

```bash
kubectl apply -f backend/ -n multi-tier

# Проверить логи
kubectl logs -n multi-tier -l app=backend
```

### 4. Запустить frontend

```bash
kubectl apply -f frontend/ -n multi-tier
```

### 5. Настроить Ingress (опционально)

```bash
# Включить Ingress в minikube
minikube addons enable ingress

# Применить Ingress
kubectl apply -f ingress.yaml -n multi-tier

# Добавить в /etc/hosts
echo "$(minikube ip) app.local" | sudo tee -a /etc/hosts
```

### 6. Открыть приложение

```bash
# Через NodePort
minikube service frontend -n multi-tier

# Или через Ingress
open http://app.local
```

## Структура проекта

```
multi-tier/
├── backend/               # Go backend
│   └── deployment.yaml   # Deployment + ConfigMap + Service
├── database/             # PostgreSQL
│   ├── statefulset.yaml  # StatefulSet + Service
│   └── secret.yaml       # Credentials
├── frontend/             # Nginx + HTML/JS
│   └── deployment.yaml   # Deployment + ConfigMap + Service
├── ingress.yaml          # Ingress для HTTP routing
└── README.md             # Эта документация
```

## Go Backend API

### Endpoints

```go
GET  /health            // Health check
GET  /api/items         // Получить все items
POST /api/items         // Создать новый item
```

### Код (упрощенно)

```go
// PostgreSQL подключение
db, _ := sql.Open("postgres", connStr)

// REST API
http.HandleFunc("/api/items", handleItems)

// JSON response
json.NewEncoder(w).Encode(items)
```

## Эксперименты

### Проверка связности

```bash
# Backend может подключиться к БД?
kubectl logs -n multi-tier -l app=backend | grep "Connected to PostgreSQL"

# Список items в БД
kubectl exec -n multi-tier deploy/postgres -- psql -U postgres appdb -c "SELECT * FROM items;"
```

### Масштабирование

```bash
# Масштабировать backend
kubectl scale deployment backend --replicas=5 -n multi-tier

# Масштабировать frontend
kubectl scale deployment frontend --replicas=3 -n multi-tier
```

### Обновление

```bash
# Обновить образ backend (если собрали свой)
kubectl set image deployment/backend backend=my-backend:v2 -n multi-tier

# Следить за rollout
kubectl rollout status deployment/backend -n multi-tier
```

## Мониторинг

```bash
# Все ресурсы в namespace
kubectl get all -n multi-tier

# Логи backend
kubectl logs -f -n multi-tier -l app=backend

# Логи PostgreSQL
kubectl logs -f -n multi-tier -l app=postgres

# Использование ресурсов
kubectl top pods -n multi-tier
```

## Troubleshooting

### Backend не подключается к БД

```bash
# Проверить PostgreSQL
kubectl get pods -n multi-tier -l app=postgres

# Проверить Secret
kubectl get secret db-secret -n multi-tier -o yaml

# Проверить подключение
kubectl exec -n multi-tier deploy/postgres -- psql -U postgres -c '\l'

# Проверить логи backend
kubectl logs -n multi-tier -l app=backend --tail=50
```

### Frontend не может обратиться к Backend

```bash
# Проверить Service backend
kubectl get svc backend -n multi-tier

# Проверить endpoints
kubectl get endpoints backend -n multi-tier

# Тест из frontend pod
kubectl exec -n multi-tier deploy/frontend -- wget -O- http://backend:3000/health
```

### Медленные запросы к БД

```bash
# Добавить индексы в PostgreSQL
kubectl exec -n multi-tier -it deploy/postgres -- psql -U postgres appdb

# В psql:
CREATE INDEX idx_items_created_at ON items(created_at);
CREATE INDEX idx_items_name ON items(name);
```

## Разработка локально

### 1. Запустить PostgreSQL

```bash
docker run -d --name postgres \
  -e POSTGRES_PASSWORD=postgres123 \
  -e POSTGRES_DB=appdb \
  -p 5432:5432 \
  postgres:15-alpine
```

### 2. Настроить переменные окружения

```bash
export DB_HOST=localhost
export DB_PORT=5432
export DB_NAME=appdb
export DB_USER=postgres
export DB_PASSWORD=postgres123
export PORT=3000
```

### 3. Запустить backend

```bash
cd backend
go run main.go
```

### 4. Тестировать API

```bash
# Health check
curl http://localhost:3000/health

# Получить items
curl http://localhost:3000/api/items

# Создать item
curl -X POST http://localhost:3000/api/items \
  -H "Content-Type: application/json" \
  -d '{"name":"Test Item"}'
```

## Production готовность

### 1. Собрать оптимизированный образ

```bash
cd backend
docker build -t backend-go:latest .

# Проверить размер
docker images backend-go:latest
# Alpine base: ~20 MB! 🎉
```

### 2. Добавить мониторинг

```go
import "github.com/prometheus/client_golang/prometheus/promhttp"

http.Handle("/metrics", promhttp.Handler())
```

### 3. Structured logging

```go
import "log/slog"

logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
logger.Info("Server starting", "port", port)
```

## Удаление

```bash
kubectl delete namespace multi-tier
```

## Связанные разделы курса

Это приложение использует концепции из разделов:
- [04-configmaps-secrets](../../04-configmaps-secrets/) - Secrets для паролей
- [05-volumes](../../05-volumes/) - PersistentVolume для PostgreSQL
- [06-ingress](../../06-ingress/) - HTTP маршрутизация
- [07-namespaces](../../07-namespaces/) - Изоляция ресурсов
- [08-resource-management](../../08-resource-management/) - Resource limits
- [09-health-checks](../../09-health-checks/) - Health endpoints
- [10-statefulsets](../../10-statefulsets/) - StatefulSet для PostgreSQL

## Следующие шаги

1. **Деплой в облако** - см. [setup/timeweb-cloud-setup.md](../../setup/timeweb-cloud-setup.md)
2. **Добавить кеширование** - Redis для кеша запросов
3. **Connection pooling** - оптимизировать подключения к БД
4. **Мониторинг** - Prometheus + Grafana
5. **CI/CD** - автоматический деплой через GitHub Actions
6. **Rate limiting** - защита API от DDoS
7. **JWT аутентификация** - защита endpoints

## Полезные ресурсы

- 📖 [Go Database/SQL Tutorial](https://go.dev/doc/database/sql-injection)
- 🎓 [Effective Go](https://go.dev/doc/effective_go)
- 📚 [pq (PostgreSQL driver)](https://github.com/lib/pq)
- 🚀 [Go Web Examples](https://gowebexamples.com/)

---

**🎉 Поздравляем! Вы развернули полноценное production-ready приложение на Go в Kubernetes!**

Теперь вы готовы к работе с реальными микросервисами на Go + Kubernetes! 💪
