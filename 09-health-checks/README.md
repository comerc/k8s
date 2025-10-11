# 09. Health Checks - Проверки здоровья

## Типы Probes

Kubernetes использует probes для проверки состояния контейнеров:

### 1. Liveness Probe

Проверяет, **жив ли контейнер**. Если проверка fails, Kubernetes перезапустит контейнер.

**Когда использовать**: для обнаружения deadlock, зависаний

### 2. Readiness Probe

Проверяет, **готов ли контейнер** принимать трафик. Если fails, Pod удаляется из Service endpoints.

**Когда использовать**: для инициализации, прогрева кеша, подключения к БД

### 3. Startup Probe

Проверяет, **запустился ли контейнер**. Пока не пройдет, другие probes не запускаются.

**Когда использовать**: для медленно стартующих приложений

## Типы проверок

### 1. HTTP GET

```yaml
livenessProbe:
  httpGet:
    path: /healthz
    port: 8080
    httpHeaders:
    - name: Custom-Header
      value: Awesome
  initialDelaySeconds: 3
  periodSeconds: 3
```

### 2. TCP Socket

```yaml
livenessProbe:
  tcpSocket:
    port: 8080
  initialDelaySeconds: 15
  periodSeconds: 20
```

### 3. Exec Command

```yaml
livenessProbe:
  exec:
    command:
    - cat
    - /tmp/healthy
  initialDelaySeconds: 5
  periodSeconds: 5
```

## Параметры конфигурации

```yaml
livenessProbe:
  httpGet:
    path: /healthz
    port: 8080
  initialDelaySeconds: 10  # Задержка перед первой проверкой
  periodSeconds: 5         # Интервал между проверками
  timeoutSeconds: 1        # Таймаут на проверку
  successThreshold: 1      # Успешных проверок для success
  failureThreshold: 3      # Неуспешных проверок для failure
```

## Пример полной конфигурации

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: app-with-probes
spec:
  containers:
  - name: app
    image: nginx:1.25
    ports:
    - containerPort: 80
    
    # Liveness: проверяем, что nginx жив
    livenessProbe:
      httpGet:
        path: /
        port: 80
      initialDelaySeconds: 10
      periodSeconds: 5
      failureThreshold: 3
    
    # Readiness: проверяем, что готов принимать трафик
    readinessProbe:
      httpGet:
        path: /
        port: 80
      initialDelaySeconds: 5
      periodSeconds: 5
      failureThreshold: 1
    
    # Startup: для медленного старта
    startupProbe:
      httpGet:
        path: /
        port: 80
      failureThreshold: 30
      periodSeconds: 10
```

## Лучшие практики

1. **Всегда** используйте readiness probe для production
2. **Используйте liveness probe** осторожно (слишком агрессивные могут вызвать restart loops)
3. **Создайте специальный endpoint** для health checks (не используйте `/`)
4. **Делайте проверки легковесными** (быстрые, без heavy operations)
5. **Startup probe** для приложений с долгой инициализацией
6. **initialDelaySeconds** должен быть больше времени старта приложения

## Rolling Updates и Probes

Readiness probe критически важна для zero-downtime deployments:
- Новый Pod не получает трафик, пока readiness не пройдет
- Старый Pod не удаляется, пока новый не станет ready

## Примеры

- `pod-with-liveness.yaml` - Liveness probe
- `pod-with-readiness.yaml` - Readiness probe
- `deployment-with-probes.yaml` - Полная конфигурация с probes

