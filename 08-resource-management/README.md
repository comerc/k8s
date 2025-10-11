# 08. Resource Management - Управление ресурсами

## Resource Requests и Limits

### Requests
Гарантированное количество ресурсов, которое кластер резервирует для контейнера.

### Limits
Максимальное количество ресурсов, которое контейнер может использовать.

```yaml
resources:
  requests:
    memory: "64Mi"
    cpu: "250m"      # 0.25 CPU
  limits:
    memory: "128Mi"
    cpu: "500m"      # 0.5 CPU
```

## CPU

- Измеряется в cores (ядрах)
- 1 CPU = 1000 millicores (m)
- 1 core = 1 vCPU (AWS) = 1 Core (GCP) = 1 vCore (Azure)

Примеры:
- `100m` = 0.1 CPU
- `500m` = 0.5 CPU
- `1` = 1 CPU
- `2000m` = 2 CPU

## Memory

Измеряется в байтах:
- `128Mi` = 128 мебибайт
- `1Gi` = 1 гибибайт
- `1G` = 1 гигабайт

**Mi vs M**: Mi = 1024^2 bytes, M = 1000^2 bytes

## QoS (Quality of Service) классы

Kubernetes автоматически присваивает QoS класс:

### 1. Guaranteed
- Requests = Limits для всех контейнеров
- Наивысший приоритет
- Удаляются последними при нехватке ресурсов

### 2. Burstable
- Requests < Limits
- Средний приоритет

### 3. BestEffort
- Нет requests и limits
- Lowest priority
- Удаляются первыми

## Horizontal Pod Autoscaler (HPA)

Автоматическое масштабирование на основе метрик:

```yaml
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: app-hpa
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: my-app
  minReplicas: 2
  maxReplicas: 10
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 70
```

### Требования для HPA

1. Установлен **Metrics Server**:
   ```bash
   minikube addons enable metrics-server
   ```

2. В Deployment указаны **resource requests**

## Vertical Pod Autoscaler (VPA)

Автоматическая настройка requests и limits (не рассматриваем подробно).

## LimitRange

Устанавливает дефолтные значения и ограничения на уровне namespace:

```yaml
apiVersion: v1
kind: LimitRange
metadata:
  name: limit-range
spec:
  limits:
  - default:
      cpu: 500m
      memory: 512Mi
    defaultRequest:
      cpu: 100m
      memory: 128Mi
    max:
      cpu: "2"
      memory: 2Gi
    min:
      cpu: 50m
      memory: 32Mi
    type: Container
```

## Лучшие практики

1. **Всегда** указывайте requests для production
2. **Указывайте limits** для предотвращения noisy neighbor
3. Используйте **HPA** для автомасштабирования
4. Мониторьте фактическое использование (`kubectl top`)
5. Настраивайте **LimitRange** на namespace level

## Примеры

- `deployment-with-resources.yaml` - Deployment с requests/limits
- `hpa-cpu.yaml` - HPA на основе CPU
- `limitrange.yaml` - LimitRange для namespace

