# 02. Deployments - Управление приложениями

## Что такое Deployment?

**Deployment** - это декларативный способ управления Pods и ReplicaSets. Это основной инструмент для развертывания stateless приложений в Kubernetes.

### Проблема с обычными Pods

- ❌ Если pod удален, он не пересоздается автоматически
- ❌ Нет автоматического масштабирования
- ❌ Сложно обновлять приложения без downtime
- ❌ Нельзя откатиться к предыдущей версии

### Что дает Deployment?

- ✅ Автоматическое пересоздание pods
- ✅ Декларативное масштабирование
- ✅ Rolling updates (обновление без downtime)
- ✅ Rollback к предыдущим версиям
- ✅ Управление жизненным циклом приложения

## Иерархия

```
Deployment
    ↓ управляет
ReplicaSet
    ↓ управляет
Pods
```

**ReplicaSet** - обеспечивает, что запущено нужное количество реплик pods. Обычно вы не работаете с ReplicaSet напрямую, а используете Deployment.

## Структура Deployment манифеста

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx
spec:
  replicas: 3                    # Количество реплик pods
  selector:                       # Как найти pods для управления
    matchLabels:
      app: nginx
  template:                       # Шаблон для создания pods
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.25
        ports:
        - containerPort: 80
```

## Стратегии обновления

### 1. RollingUpdate (по умолчанию)

Постепенная замена старых pods на новые:

```yaml
spec:
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxSurge: 1        # Макс. количество новых pods сверх replicas
      maxUnavailable: 0  # Макс. количество unavailable pods
```

**Преимущества**: zero downtime, постепенное обновление
**Когда использовать**: почти всегда (по умолчанию)

### 2. Recreate

Сначала удаляет все старые pods, потом создает новые:

```yaml
spec:
  strategy:
    type: Recreate
```

**Преимущества**: простота
**Недостатки**: downtime во время обновления
**Когда использовать**: когда невозможно иметь две версии одновременно

## Масштабирование

### Ручное масштабирование

```bash
# Изменить количество реплик
kubectl scale deployment nginx-deployment --replicas=5

# В манифесте
spec:
  replicas: 5
```

### Автоматическое масштабирование (HPA)

Будет рассмотрено в разделе **08-resource-management**.

## История изменений и откат

Deployment сохраняет историю изменений (revision history):

```bash
# Посмотреть историю
kubectl rollout history deployment/nginx-deployment

# Откатиться к предыдущей версии
kubectl rollout undo deployment/nginx-deployment

# Откатиться к конкретной версии
kubectl rollout undo deployment/nginx-deployment --to-revision=2

# Статус обновления
kubectl rollout status deployment/nginx-deployment

# Пауза обновления
kubectl rollout pause deployment/nginx-deployment

# Возобновление обновления
kubectl rollout resume deployment/nginx-deployment
```

## Жизненный цикл обновления

```
1. Изменение манифеста (новый image)
       ↓
2. Создание нового ReplicaSet
       ↓
3. Постепенное создание новых pods
       ↓
4. Удаление старых pods
       ↓
5. Завершение rollout
```

## Когда использовать Deployment?

✅ **Используйте Deployment для**:
- Stateless приложений (web серверы, API)
- Приложений, которые можно легко масштабировать
- Приложений, требующих rolling updates

❌ **НЕ используйте Deployment для**:
- Stateful приложений с persistent data (используйте StatefulSet)
- Одноразовых задач (используйте Job)
- Задач по расписанию (используйте CronJob)

## Лучшие практики

1. **Всегда используйте labels**: для связи Deployment → Pods
2. **Указывайте версии образов**: избегайте тега `latest`
3. **Настраивайте readiness/liveness probes**: для правильного rolling update
4. **Устанавливайте resource requests/limits**: для стабильности
5. **Сохраняйте историю**: `revisionHistoryLimit: 10`
6. **Используйте maxUnavailable: 0**: для zero-downtime обновлений

## Практические примеры

В этой директории:
- `deployment-simple.yaml` - базовый deployment
- `deployment-with-strategy.yaml` - с настроенной стратегией
- `deployment-nginx.yaml` - полноценный nginx deployment
- `deployment-multiple-replicas.yaml` - deployment с несколькими репликами

## Основные команды

См. `commands.md` для полного списка команд.

## Практическое задание

См. `exercise.md` для практики.

## Что дальше?

После освоения Deployments переходите к **03-services** для изучения сетевого взаимодействия между pods.

## Полезные ссылки

- [Официальная документация Deployments](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)
- [Deployment Strategies](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#strategy)

