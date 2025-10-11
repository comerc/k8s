# 03. Services - Сетевое взаимодействие

## Что такое Service?

**Service** - это абстракция, которая определяет логический набор Pods и политику доступа к ним.

### Проблема без Services

- ❌ IP адреса Pods постоянно меняются (при пересоздании)
- ❌ Нужен способ load balancing между репликами
- ❌ Сложно организовать service discovery
- ❌ Нет единой точки входа для группы Pods

### Что дает Service?

- ✅ Стабильный IP адрес (ClusterIP)
- ✅ Стабильное DNS имя
- ✅ Load balancing между Pods
- ✅ Service discovery внутри кластера
- ✅ Доступ к приложению извне (LoadBalancer, NodePort)

## Типы Services

### 1. ClusterIP (по умолчанию)

Создает внутренний IP адрес, доступный только внутри кластера.

```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-service
spec:
  type: ClusterIP
  selector:
    app: nginx
  ports:
  - port: 80          # Порт Service
    targetPort: 80    # Порт контейнера
```

**Когда использовать**: внутренние сервисы (БД, кеш, внутренние API)

### 2. NodePort

Открывает порт на каждой ноде кластера (30000-32767).

```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-service
spec:
  type: NodePort
  selector:
    app: nginx
  ports:
  - port: 80
    targetPort: 80
    nodePort: 30080   # Опционально (будет выбран автоматически)
```

**Доступ**: `<NodeIP>:30080`
**Когда использовать**: для тестирования, быстрого доступа извне

### 3. LoadBalancer

Создает внешний load balancer (в облаке).

```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-service
spec:
  type: LoadBalancer
  selector:
    app: nginx
  ports:
  - port: 80
    targetPort: 80
```

**Когда использовать**: production приложения, доступные из интернета
**Примечание**: в minikube требуется `minikube tunnel`

### 4. ExternalName

Создает CNAME запись для внешнего сервиса.

```yaml
apiVersion: v1
kind: Service
metadata:
  name: external-api
spec:
  type: ExternalName
  externalName: api.example.com
```

**Когда использовать**: для доступа к внешним сервисам через DNS

## Как Service находит Pods?

Service использует **селекторы** для выбора Pods:

```yaml
# Service
spec:
  selector:
    app: nginx        # Ищет все Pods с меткой app=nginx
    tier: frontend

# Pod/Deployment
metadata:
  labels:
    app: nginx        # Эта метка связывает Pod с Service
    tier: frontend
```

## Service Discovery

### 1. Через DNS (рекомендуется)

Kubernetes автоматически создает DNS записи:

```bash
# Формат: <service-name>.<namespace>.svc.cluster.local
my-service.default.svc.cluster.local

# Или короткая форма (в том же namespace)
my-service
```

Из любого Pod в кластере:

```bash
kubectl exec -it my-pod -- curl http://my-service
```

### 2. Через переменные окружения

Kubernetes автоматически создает переменные:

```bash
MY_SERVICE_SERVICE_HOST=10.96.0.10
MY_SERVICE_SERVICE_PORT=80
```

## Endpoints

Kubernetes автоматически создает объект **Endpoints**, который содержит IP адреса всех подходящих Pods:

```bash
kubectl get endpoints my-service
```

Если Pods нет или они не готовы, Endpoints будет пустым.

## Порты

```yaml
ports:
- name: http          # Имя порта (опционально, но рекомендуется)
  protocol: TCP       # TCP или UDP
  port: 80            # Порт, на котором Service слушает
  targetPort: 8080    # Порт контейнера (куда перенаправляется трафик)
  nodePort: 30080     # Только для NodePort
```

**port vs targetPort**:
- `port`: порт Service (к чему обращаются клиенты)
- `targetPort`: порт контейнера (куда идет трафик)
- Могут быть одинаковыми или разными

## Session Affinity

По умолчанию Service балансирует случайным образом. Для "залипания" клиента к одному Pod:

```yaml
spec:
  sessionAffinity: ClientIP
  sessionAffinityConfig:
    clientIP:
      timeoutSeconds: 10800  # 3 часа
```

## Headless Service

Service без ClusterIP - для прямого доступа к Pods:

```yaml
spec:
  clusterIP: None    # Headless
  selector:
    app: database
```

DNS вернет IP адреса всех Pods напрямую.

**Когда использовать**: StatefulSets, прямое подключение к конкретному Pod

## Multi-Port Services

Service может expose несколько портов:

```yaml
ports:
- name: http
  port: 80
  targetPort: 8080
- name: https
  port: 443
  targetPort: 8443
- name: metrics
  port: 9090
  targetPort: 9090
```

## Лучшие практики

1. **Используйте ClusterIP для внутренних сервисов**
2. **Используйте LoadBalancer для production** (не NodePort)
3. **Называйте порты** (name: http, name: https)
4. **Проверяйте Endpoints** если Service не работает
5. **Используйте DNS** для service discovery (не env vars)
6. **Для stateful приложений** используйте Headless Service

## Практические примеры

В этой директории:
- `service-clusterip.yaml` - внутренний сервис
- `service-nodeport.yaml` - доступ через ноду
- `service-loadbalancer.yaml` - external load balancer
- `service-multi-port.yaml` - несколько портов
- `service-headless.yaml` - headless service

## Основные команды

См. `commands.md` для полного списка команд.

## Практическое задание

См. `exercise.md` для практики.

## Диаграмма

```
Internet
   ↓
LoadBalancer Service (external IP)
   ↓
NodePort Service (node:30080)
   ↓
ClusterIP Service (10.96.0.10:80)
   ↓ (selector: app=nginx)
   ↓ (load balancing)
   ├→ Pod 1 (10.244.0.5:8080)
   ├→ Pod 2 (10.244.0.6:8080)
   └→ Pod 3 (10.244.0.7:8080)
```

## Что дальше?

После освоения Services переходите к **04-configmaps-secrets** для изучения управления конфигурацией.

## Полезные ссылки

- [Официальная документация Services](https://kubernetes.io/docs/concepts/services-networking/service/)
- [Service Types](https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types)
- [DNS for Services](https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/)

