# 10. StatefulSets - Stateful приложения

## Что такое StatefulSet?

**StatefulSet** - контроллер для управления stateful приложениями, которым требуется:
- Стабильные, уникальные идентификаторы
- Стабильное сетевое имя
- Постоянное хранилище
- Упорядоченное развертывание и масштабирование

## StatefulSet vs Deployment

| Аспект | Deployment | StatefulSet |
|--------|-----------|-------------|
| Имена Pods | Случайные (hash) | Предсказуемые (0,1,2...) |
| DNS имена | Нестабильные | Стабильные |
| Storage | Shared или без | Уникальный PVC для каждого Pod |
| Порядок создания | Параллельно | Последовательно (0→1→2) |
| Порядок удаления | Параллельно | Обратный (2→1→0) |
| Использование | Stateless apps | Stateful apps (БД, очереди) |

## Структура StatefulSet

```yaml
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: mysql
spec:
  serviceName: mysql          # Headless Service
  replicas: 3
  selector:
    matchLabels:
      app: mysql
  template:
    metadata:
      labels:
        app: mysql
    spec:
      containers:
      - name: mysql
        image: mysql:8
        ports:
        - containerPort: 3306
        volumeMounts:
        - name: data
          mountPath: /var/lib/mysql
  volumeClaimTemplates:       # Автоматически создает PVC
  - metadata:
      name: data
    spec:
      accessModes: [ "ReadWriteOnce" ]
      resources:
        requests:
          storage: 10Gi
```

## Headless Service

StatefulSet требует Headless Service для стабильных DNS имен:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: mysql
spec:
  clusterIP: None    # Headless
  selector:
    app: mysql
  ports:
  - port: 3306
```

## DNS имена Pods

Формат: `<pod-name>.<service-name>.<namespace>.svc.cluster.local`

Примеры:
- `mysql-0.mysql.default.svc.cluster.local`
- `mysql-1.mysql.default.svc.cluster.local`
- `mysql-2.mysql.default.svc.cluster.local`

## VolumeClaimTemplates

Автоматически создает PVC для каждого Pod:

```yaml
volumeClaimTemplates:
- metadata:
    name: data
  spec:
    accessModes: [ "ReadWriteOnce" ]
    storageClassName: fast
    resources:
      requests:
        storage: 10Gi
```

Создаст: `data-mysql-0`, `data-mysql-1`, `data-mysql-2`

## Порядок развертывания

### Создание
```
mysql-0 → wait ready → mysql-1 → wait ready → mysql-2
```

### Удаление
```
mysql-2 → mysql-1 → mysql-0
```

## Масштабирование

```bash
# Увеличить до 5 реплик
kubectl scale statefulset mysql --replicas=5
# Создаст mysql-3, затем mysql-4

# Уменьшить до 2
kubectl scale statefulset mysql --replicas=2
# Удалит mysql-4, затем mysql-3
```

## Обновление

По умолчанию: `RollingUpdate` в обратном порядке (2→1→0)

```yaml
spec:
  updateStrategy:
    type: RollingUpdate
    rollingUpdate:
      partition: 0  # Обновить все Pods с индексом >= partition
```

## Примеры использования

- **Базы данных**: MySQL, PostgreSQL, MongoDB
- **Кеши**: Redis, Memcached
- **Очереди сообщений**: RabbitMQ, Kafka
- **Распределенные системы**: Zookeeper, etcd, Consul

## Лучшие практики

1. Используйте **Headless Service**
2. Настраивайте **PodManagementPolicy** при необходимости
3. Используйте **PodDisruptionBudget** для защиты от disruptions
4. **Backup данных** из PVCs
5. Тестируйте **failover scenarios**

## Примеры

- `statefulset-nginx.yaml` - Простой StatefulSet
- `statefulset-mysql.yaml` - MySQL кластер

