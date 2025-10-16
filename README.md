# 🚀 Kubernetes на Golang - Полный курс

Структурированный курс для освоения Kubernetes с практическими примерами на Go.

**Для кого**: Go разработчики  
**Длительность**: 4-5 недель  
**Формат**: Теория → Практика → Проект

---

## 📖 Содержание

- [Введение](#введение)
- [Прогресс обучения](#-прогресс-обучения)
- [Быстрый старт](#-быстрый-старт)
- [Программа курса](#-программа-курса)
- [Best Practices Go](#-best-practices-go--kubernetes)
- [Справочник](#-справочник)
- [Ресурсы](#-дополнительные-ресурсы)

---

## Введение

### Почему Go + Kubernetes?

| Аспект | Преимущество |
|--------|--------------|
| **Производительность** | В 10-50x быстрее Python/Node.js |
| **Память** | В 5-10x меньше потребление |
| **Startup** | <2 секунды (vs 10-30 сек) |
| **Размер образа** | 10-20 MB (vs 200+ MB) |
| **Интеграция** | Kubernetes написан на Go |

### Структура проекта

```
k8s/
├── 01-basics/              # Pods
├── 02-deployments/         # Управление приложениями
├── 03-services/            # Сетевое взаимодействие
├── 04-configmaps-secrets/  # Конфигурация
├── 05-volumes/             # Хранение данных
├── 06-ingress/             # HTTP маршрутизация
├── 07-namespaces/          # Изоляция ресурсов
├── 08-resource-management/ # Лимиты и HPA
├── 09-health-checks/       # Health probes
├── 10-statefulsets/        # Stateful приложения
├── 11-jobs-cronjobs/       # Задачи по расписанию
├── 12-rbac/                # Безопасность
│
├── demo-apps/
│   ├── hello-web/          # 🟢 Простое
│   ├── counter-app/        # 🟡 Среднее (Go + Redis)
│   └── multi-tier/         # 🔴 Сложное (Go + PostgreSQL)
│
├── scripts/
│   ├── start-minikube.sh
│   └── cleanup.sh
│
└── setup/
    ├── minikube-setup.md
    └── timeweb-cloud-setup.md
```

---

## ✅ Прогресс обучения

Отмечайте пройденные этапы:

### Неделя 1-2: Основы
- [ ] Установка и настройка окружения
- [ ] **01-basics** - Pods и Labels
- [ ] **02-deployments** - Deployments и Scaling
- [ ] **03-services** - Services и DNS
- [ ] **🎯 Проект**: hello-web

### Неделя 2-3: Конфигурация и данные
- [ ] **04-configmaps-secrets** - ConfigMaps и Secrets
- [ ] **05-volumes** - PersistentVolume и PVC
- [ ] **07-namespaces** - Namespaces и ResourceQuota
- [ ] **🎯 Проект**: counter-app (Go + Redis)

### Неделя 3-4: Продвинутые темы
- [ ] **06-ingress** - Ingress и маршрутизация
- [ ] **08-resource-management** - Limits и HPA
- [ ] **09-health-checks** - Health probes
- [ ] **10-statefulsets** - StatefulSets

### Неделя 4-5: Автоматизация и production
- [ ] **11-jobs-cronjobs** - Jobs и CronJobs
- [ ] **12-rbac** - RBAC и безопасность
- [ ] **🎯 Проект**: multi-tier (Go + PostgreSQL)
- [ ] Деплой в облако (Timeweb Cloud)

### Бонус (опционально)
- [ ] CI/CD с GitHub Actions
- [ ] Мониторинг (Prometheus + Grafana)
- [ ] Service Mesh (Istio)

---

## 🚀 Быстрый старт

### 1. Установка

```bash
# macOS
brew install kubectl minikube

# Проверка
kubectl version --client
minikube version
```

📖 **Подробная инструкция**: [setup/minikube-setup.md](setup/minikube-setup.md)

### 2. Запуск кластера

```bash
./scripts/start-minikube.sh
```

### 3. Первый Pod

```bash
cd 01-basics
kubectl apply -f pod-simple.yaml
kubectl get pods
```

**🎉 Готово! Вы запустили первый Pod в Kubernetes!**

---

## 📚 Программа курса

### Неделя 1-2: Основы

#### 01. Pods - базовая единица

```bash
cd 01-basics
kubectl apply -f pod-simple.yaml
kubectl get pods
kubectl logs pod-simple
kubectl exec -it pod-simple -- /bin/sh
```

**Что изучите**: Pod, Labels, Multi-container, Logs, Exec  
**Время**: 2-3 часа  
**📖 Материалы**: [01-basics/README.md](01-basics/) | [commands.md](01-basics/commands.md) | [exercise.md](01-basics/exercise.md)

#### 02. Deployments - управление приложениями

```bash
cd 02-deployments
kubectl apply -f deployment-simple.yaml
kubectl scale deployment nginx-deployment --replicas=5
kubectl set image deployment/nginx-deployment nginx=nginx:1.26
kubectl rollout undo deployment/nginx-deployment
```

**Что изучите**: ReplicaSet, Scaling, Rolling updates, Rollback  
**Время**: 3-4 часа  
**📖 Материалы**: [02-deployments/README.md](02-deployments/)

#### 03. Services - сетевое взаимодействие

```bash
cd 03-services
kubectl apply -f service-clusterip.yaml
kubectl get svc
kubectl run test --image=curlimages/curl -it --rm -- curl http://nginx-service
```

**Что изучите**: ClusterIP, NodePort, LoadBalancer, DNS  
**Время**: 2-3 часа  
**📖 Материалы**: [03-services/README.md](03-services/)

#### 🎯 Практика: hello-web

Примените знания на реальном приложении:

```bash
cd demo-apps/hello-web
kubectl apply -f deployment.yaml
minikube service hello-web
```

**📖 Инструкции**: [demo-apps/hello-web/README.md](demo-apps/hello-web/)

---

### Неделя 2-3: Конфигурация и данные

#### 04. ConfigMaps и Secrets

```bash
cd 04-configmaps-secrets
kubectl create configmap app-config --from-literal=key=value
kubectl create secret generic db-secret --from-literal=password=secret
kubectl apply -f configmap-env.yaml
```

**Что изучите**: Управление конфигурацией, Secrets, Environment  
**Время**: 2 часа  
**📖 Материалы**: [04-configmaps-secrets/README.md](04-configmaps-secrets/)

#### 05. Volumes - хранение данных

```bash
cd 05-volumes
kubectl apply -f pvc-simple.yaml
kubectl get pvc
kubectl get pv
```

**Что изучите**: PersistentVolume, PVC, StorageClass  
**Время**: 2-3 часа  
**📖 Материалы**: [05-volumes/README.md](05-volumes/)

#### 07. Namespaces

```bash
cd 07-namespaces
kubectl create namespace development
kubectl get namespaces
kubectl config set-context --current --namespace=development
```

**Что изучите**: Изоляция, ResourceQuota, LimitRange  
**Время**: 1-2 часа  
**📖 Материалы**: [07-namespaces/README.md](07-namespaces/)

#### 🎯 Практика: counter-app

Go приложение с Redis:

```bash
cd demo-apps/counter-app
kubectl apply -f deployment.yaml
minikube service counter-app
```

**📖 Инструкции**: [demo-apps/counter-app/README.md](demo-apps/counter-app/)

---

### Неделя 3-4: Продвинутые темы

#### 06. Ingress

```bash
minikube addons enable ingress
cd 06-ingress
kubectl apply -f ingress-simple.yaml
```

**Что изучите**: Ingress Controller, Routing, TLS  
**Время**: 2-3 часа  
**📖 Материалы**: [06-ingress/README.md](06-ingress/)

#### 08. Resource Management

```bash
cd 08-resource-management
kubectl apply -f hpa-cpu.yaml
kubectl top pods
```

**Что изучите**: Requests/Limits, HPA, QoS  
**Время**: 2 часа  
**📖 Материалы**: [08-resource-management/README.md](08-resource-management/)

#### 09. Health Checks

```bash
cd 09-health-checks
kubectl apply -f deployment-with-probes.yaml
kubectl describe pod <pod-name>
```

**Что изучите**: Liveness, Readiness, Startup probes  
**Время**: 2 часа  
**📖 Материалы**: [09-health-checks/README.md](09-health-checks/)

#### 10. StatefulSets

```bash
cd 10-statefulsets
kubectl apply -f statefulset-nginx.yaml
kubectl get statefulsets
```

**Что изучите**: StatefulSet, Headless Service, Volumes  
**Время**: 3 часа  
**📖 Материалы**: [10-statefulsets/README.md](10-statefulsets/)

---

### Неделя 4-5: Автоматизация и production

#### 11. Jobs и CronJobs

```bash
cd 11-jobs-cronjobs
kubectl apply -f job-simple.yaml
kubectl apply -f cronjob-simple.yaml
kubectl get jobs
kubectl get cronjobs
```

**Что изучите**: Batch задачи, Расписания  
**Время**: 2 часа  
**📖 Материалы**: [11-jobs-cronjobs/README.md](11-jobs-cronjobs/)

#### 12. RBAC

```bash
cd 12-rbac
kubectl apply -f serviceaccount-with-role.yaml
kubectl auth can-i list pods
```

**Что изучите**: Roles, ServiceAccounts, Security  
**Время**: 2-3 часа  
**📖 Материалы**: [12-rbac/README.md](12-rbac/)

#### 🎯 Финальный проект: multi-tier

**📖 Инструкции**: [demo-apps/multi-tier/README.md](demo-apps/multi-tier/)

---

## 💡 Best Practices: Go + Kubernetes

### 1. Оптимальный Dockerfile

```dockerfile
# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o app .

# Runtime stage  
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app"]
```

**Результат**: образ ~10-15 MB (vs 300+ MB без multi-stage)

### 2. Health Check Endpoint

```go
func healthHandler(w http.ResponseWriter, r *http.Request) {
    // Проверить зависимости (БД, Redis и т.д.)
    if err := db.Ping(); err != nil {
        w.WriteHeader(http.StatusServiceUnavailable)
        json.NewEncoder(w).Encode(map[string]string{
            "status": "unhealthy",
            "error": err.Error(),
        })
        return
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}
```

Kubernetes манифест:

```yaml
livenessProbe:
  httpGet:
    path: /health
    port: 8080
  initialDelaySeconds: 10
readinessProbe:
  httpGet:
    path: /health
    port: 8080
  initialDelaySeconds: 5
```

### 3. Graceful Shutdown

```go
func main() {
    server := &http.Server{Addr: ":8080", Handler: router}
    
    go func() {
        if err := server.ListenAndServe(); err != nil {
            log.Printf("Server error: %v", err)
        }
    }()
    
    // Ждём сигнал завершения
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    
    log.Println("Shutting down...")
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    if err := server.Shutdown(ctx); err != nil {
        log.Fatal("Forced shutdown:", err)
    }
    log.Println("Server stopped")
}
```

### 4. Конфигурация через ENV

```go
type Config struct {
    Port        string
    DatabaseURL string
    RedisHost   string
}

func loadConfig() *Config {
    return &Config{
        Port:        getEnv("PORT", "8080"),
        DatabaseURL: os.Getenv("DATABASE_URL"),
        RedisHost:   getEnv("REDIS_HOST", "localhost"),
    }
}

func getEnv(key, fallback string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return fallback
}
```

### 5. Structured Logging

```go
import "log/slog"

func main() {
    logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
        Level: slog.LevelInfo,
    }))
    
    logger.Info("server starting",
        "port", port,
        "version", version,
        "env", environment,
    )
}
```

---

## 📖 Справочник

### Основные команды kubectl

```bash
# Просмотр
kubectl get pods
kubectl get deployments
kubectl get services
kubectl get all

# Детали
kubectl describe pod <name>
kubectl logs <pod-name>
kubectl logs -f <pod-name>

# Выполнение
kubectl exec -it <pod-name> -- /bin/sh

# Применение
kubectl apply -f file.yaml
kubectl delete -f file.yaml

# Масштабирование
kubectl scale deployment <name> --replicas=3

# Обновление
kubectl set image deployment/<name> container=image:tag
kubectl rollout status deployment/<name>
kubectl rollout undo deployment/<name>

# Отладка
kubectl get events
kubectl top pods
kubectl port-forward <pod> 8080:80
```

### Очистка

```bash
# Очистить учебные ресурсы
./scripts/cleanup.sh

# Остановить кластер
minikube stop

# Удалить кластер
minikube delete
```

---

## 📚 Дополнительные ресурсы

### Официальная документация
- [Kubernetes Docs](https://kubernetes.io/docs/)
- [Kubectl Cheat Sheet](https://kubernetes.io/docs/reference/kubectl/cheatsheet/)
- [Go Documentation](https://go.dev/doc/)
- [Effective Go](https://go.dev/doc/effective_go)

### Интерактивное обучение
- [Play with Kubernetes](https://labs.play-with-k8s.com/)
- [Katacoda Kubernetes](https://www.katacoda.com/courses/kubernetes)

### Go + Kubernetes
- [client-go](https://github.com/kubernetes/client-go)
- [Operator SDK](https://sdk.operatorframework.io/)

---

## 🎯 Что дальше?

После завершения курса:

1. **Деплой в облако** - [Timeweb Cloud](setup/timeweb-cloud-setup.md)
2. **CI/CD** - GitHub Actions, ArgoCD
3. **Мониторинг** - Prometheus + Grafana
4. **Service Mesh** - Istio, Linkerd
5. **Helm Charts** (Terraform + Ansible + Helm)
6. **Сертификация** - CKA, CKAD

---

## 📝 Лицензия

MIT License - см. [LICENSE](LICENSE)

---

**🎉 Удачи в изучении!**

*Kubernetes написан на Go не случайно - это идеальная пара для облачных приложений!*

---

TODO: Flannel