# Counter App - Приложение со счетчиком на Go

Go приложение с Redis для демонстрации работы с stateful данными в Kubernetes.

**Уровень сложности**: 🟡 Средний

## Что изучите

- ✅ **Go веб-сервер** - net/http, html/template
- ✅ **Multi-container pods** - приложение + Redis
- ✅ **Service discovery** - подключение к Redis через DNS
- ✅ **ConfigMap** - конфигурация приложения
- ✅ **Health checks** - /health endpoint
- ✅ **Работа с Redis** - go-redis клиент
- ✅ **Multi-stage Docker build** - оптимизация образа

## Технологии

- **Backend**: Go 1.21 (net/http)
- **Cache**: Redis 7
- **Kubernetes**: ConfigMap, Deployment, Service

## Производительность

| Метрика | Значение |
|---------|----------|
| Образ Docker | ~15 MB |
| Память (runtime) | ~15 MB |
| Startup время | <2 сек |
| Запросов/сек | 10,000+ |

## Быстрый старт

### Вариант 1: Запуск в Kubernetes (рекомендуется)

```bash
# Применить манифесты
kubectl apply -f deployment.yaml

# Проверить статус
kubectl get pods -l app=counter-app-go
kubectl get svc counter-app-go

# Открыть в браузере
minikube service counter-app-go
```

### Вариант 2: Локальная разработка

```bash
# Установить зависимости
go mod download

# Запустить Redis локально
docker run -d -p 6379:6379 redis:7-alpine

# Установить переменные окружения
export REDIS_HOST=localhost
export REDIS_PORT=6379
export PORT=8080

# Запустить приложение
go run main.go

# Открыть браузер
open http://localhost:8080
```

### Вариант 3: Сборка Docker образа

```bash
# Собрать образ
docker build -t counter-app-go:latest .

# Запустить контейнер
docker run -d -p 8080:8080 \
  -e REDIS_HOST=redis \
  -e REDIS_PORT=6379 \
  counter-app-go:latest
```

## Структура проекта

```
counter-app-go/
├── main.go           # Основной код приложения
├── go.mod            # Go модули
├── go.sum            # Checksums зависимостей
├── Dockerfile        # Multi-stage Docker build
├── deployment.yaml   # Kubernetes манифесты
└── README.md         # Документация
```

## Основной код (main.go)

Приложение использует стандартную библиотеку Go:

```go
// net/http для веб-сервера
// html/template для шаблонов
// github.com/go-redis/redis/v8 для Redis
```

### Endpoints

- `GET /` - главная страница с счетчиком
- `POST /increment` - увеличить счетчик
- `POST /reset` - сбросить счетчик
- `GET /health` - health check для K8s

## Эксперименты

### Тест производительности

```bash
# Установить hey (HTTP load generator)
go install github.com/rakyll/hey@latest

# Тест нагрузки
hey -n 10000 -c 100 http://$(minikube service counter-app-go --url)

# Go должен показать отличные результаты! 🚀
```

### Сравнение с Python версией

```bash
# Запустить обе версии
kubectl apply -f ../counter-app/deployment.yaml
kubectl apply -f deployment.yaml

# Сравнить использование памяти
kubectl top pods -l app=counter-app      # Python
kubectl top pods -l app=counter-app-go   # Go

# Go обычно использует в 5-10 раз меньше памяти!
```

### Масштабирование

```bash
# Увеличить до 10 реплик
kubectl scale deployment counter-app-go --replicas=10

# Проверить, как быстро поднялись
kubectl get pods -l app=counter-app-go --watch

# Go pods стартуют почти мгновенно!
```

## Разработка

### Добавить новый endpoint

```go
func main() {
    // ... existing code ...
    
    http.HandleFunc("/api/stats", handleStats)
}

func handleStats(w http.ResponseWriter, r *http.Request) {
    count, _ := redisClient.Get(ctx, "counter").Result()
    fmt.Fprintf(w, `{"count": %s, "hostname": "%s"}`, count, hostname)
}
```

### Hot reload для разработки

```bash
# Установить air (hot reload для Go)
go install github.com/cosmtrek/air@latest

# Запустить с hot reload
air
```

## Тестирование

```bash
# Запустить тесты (если добавите)
go test ./...

# Запустить тесты с coverage
go test -cover ./...

# Benchmark тесты
go test -bench=. ./...
```

## Сборка оптимизированного бинарника

```bash
# Статически скомпилированный бинарник
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-s -w' -o counter-app .

# Проверить размер (должен быть ~8-12 MB)
ls -lh counter-app

# Для еще меньшего размера используйте UPX
upx --best --lzma counter-app
```

## Production best practices

### 1. Graceful shutdown

```go
// Добавить обработку сигналов
import "os/signal"

func main() {
    // ... setup ...
    
    server := &http.Server{Addr: ":8080"}
    
    go func() {
        if err := server.ListenAndServe(); err != nil {
            log.Println(err)
        }
    }()
    
    // Wait for interrupt signal
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    <-c
    
    ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
    defer cancel()
    server.Shutdown(ctx)
    log.Println("Shutting down gracefully")
}
```

### 2. Structured logging

```go
import "log/slog"

logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
logger.Info("Server starting", "port", port)
```

### 3. Metrics для Prometheus

```go
import "github.com/prometheus/client_golang/prometheus/promhttp"

http.Handle("/metrics", promhttp.Handler())
```

## Troubleshooting

### Приложение не подключается к Redis

```bash
# Проверить, что Redis работает
kubectl exec -it deploy/redis -- redis-cli ping

# Проверить логи Go приложения
kubectl logs -l app=counter-app-go

# Проверить переменные окружения
kubectl exec -it deploy/counter-app-go -- env | grep REDIS
```

### Go модули не скачиваются

```bash
# Проверить прокси
go env GOPROXY

# Использовать direct mode
GOPROXY=direct go mod download
```

## Удаление

```bash
kubectl delete -f deployment.yaml
```

## Полезные ресурсы

- 📖 [Effective Go](https://go.dev/doc/effective_go)
- 🎓 [Go by Example](https://gobyexample.com/)
- 📚 [Go Redis Client](https://github.com/go-redis/redis)
- 🐳 [Go Docker Best Practices](https://docs.docker.com/language/golang/)

## Связанные разделы курса

Это приложение использует концепции из разделов:
- [02-deployments](../../02-deployments/) - Deployments и ReplicaSets
- [03-services](../../03-services/) - Service Discovery
- [04-configmaps-secrets](../../04-configmaps-secrets/) - ConfigMaps
- [09-health-checks](../../09-health-checks/) - Health Probes

## Что дальше?

1. Добавьте unit тесты для Go кода
2. Реализуйте middleware для логирования (slog)
3. Добавьте Prometheus metrics
4. Настройте HPA (Horizontal Pod Autoscaler)
5. Попробуйте использовать PersistentVolume для Redis

После освоения counter-app переходите к **[multi-tier](../multi-tier/)** для полноценного production приложения с PostgreSQL.

