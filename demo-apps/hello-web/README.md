# Hello Web - Простое веб-приложение

Простое веб-приложение на nginx для демонстрации базовых концепций Kubernetes.

**Уровень сложности**: 🟢 Начальный

## Что изучите

- ✅ **Pods** - минимальная единица развертывания
- ✅ **Deployment** - управление репликами
- ✅ **Service (NodePort)** - доступ к приложению извне
- ✅ **ConfigMap** - хранение конфигурации (HTML)
- ✅ **Resource limits/requests** - управление ресурсами
- ✅ **Health checks** - liveness и readiness probes

## Технологии

- Nginx 1.25
- Kubernetes manifests (YAML)

## Быстрый старт

### 1. Запустить приложение

```bash
kubectl apply -f deployment.yaml
```

### 2. Проверить статус

```bash
kubectl get deployments hello-web
kubectl get pods -l app=hello-web
kubectl get service hello-web
```

### 3. Открыть в браузере

```bash
# В minikube
minikube service hello-web

# Или получить URL
minikube service hello-web --url
```

### 4. Проверить load balancing

Обновите страницу несколько раз - запросы будут распределяться между репликами.

## Эксперименты

### Масштабирование

```bash
# Увеличить до 5 реплик
kubectl scale deployment hello-web --replicas=5

# Проверить
kubectl get pods -l app=hello-web
```

### Обновление контента

```bash
# Отредактировать ConfigMap
kubectl edit configmap hello-web-html

# Перезапустить pods для применения изменений
kubectl rollout restart deployment hello-web
```

### Тест устойчивости

```bash
# Удалить один pod
kubectl delete pod -l app=hello-web --field-selector=status.phase=Running | head -1

# Kubernetes автоматически создаст новый
kubectl get pods -l app=hello-web --watch
```

## Доступ через Ingress

Если включен Ingress:

```bash
# Включить Ingress в minikube
minikube addons enable ingress

# Применить Ingress
kubectl apply -f ingress.yaml

# Добавить в /etc/hosts
echo "$(minikube ip) hello.local" | sudo tee -a /etc/hosts

# Открыть в браузере
open http://hello.local
```

## Удаление

```bash
kubectl delete -f deployment.yaml
```

## Структура

```
hello-web/
├── Dockerfile           # Docker образ (опционально)
├── index.html          # HTML страница
├── deployment.yaml     # Kubernetes манифесты
├── ingress.yaml       # Ingress (опционально)
└── README.md          # Эта документация
```

## Связанные разделы курса

Это приложение использует концепции из разделов:
- [01-basics](../../01-basics/) - Pods
- [02-deployments](../../02-deployments/) - Deployments
- [03-services](../../03-services/) - Services
- [04-configmaps-secrets](../../04-configmaps-secrets/) - ConfigMaps

## Что дальше?

После освоения hello-web переходите к **[counter-app](../counter-app/)** для изучения работы с Go, Redis и multi-container pods.

