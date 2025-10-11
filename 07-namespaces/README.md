# 07. Namespaces - Изоляция ресурсов

## Что такое Namespace?

**Namespace** - это виртуальный кластер внутри физического Kubernetes кластера. Способ разделения ресурсов между пользователями/командами/проектами.

## Зачем нужны Namespaces?

- ✅ Изоляция ресурсов между проектами/командами
- ✅ Организация ресурсов (dev, staging, prod)
- ✅ Применение квот и лимитов на уровне namespace
- ✅ RBAC политики на уровне namespace

## Системные Namespaces

Kubernetes создает несколько namespaces по умолчанию:

- **default** - namespace по умолчанию
- **kube-system** - системные компоненты Kubernetes
- **kube-public** - публичные ресурсы
- **kube-node-lease** - heartbeats нод

## Создание Namespace

```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: development
```

```bash
kubectl create namespace development
```

## Использование Namespace

```bash
# Создать ресурс в namespace
kubectl apply -f deployment.yaml -n development

# Получить ресурсы из namespace
kubectl get pods -n development

# Установить namespace по умолчанию
kubectl config set-context --current --namespace=development

# Все namespaces
kubectl get pods --all-namespaces
kubectl get pods -A
```

## DNS и Namespaces

Формат DNS внутри кластера:
```
<service-name>.<namespace>.svc.cluster.local
```

Примеры:
- `nginx-service.development.svc.cluster.local`
- `database.production.svc.cluster.local`

## Resource Quotas

Ограничение ресурсов в namespace:

```yaml
apiVersion: v1
kind: ResourceQuota
metadata:
  name: compute-quota
  namespace: development
spec:
  hard:
    requests.cpu: "10"
    requests.memory: 20Gi
    limits.cpu: "20"
    limits.memory: 40Gi
    pods: "10"
```

## LimitRange

Установка дефолтных лимитов для ресурсов:

```yaml
apiVersion: v1
kind: LimitRange
metadata:
  name: limit-range
  namespace: development
spec:
  limits:
  - default:
      cpu: 500m
      memory: 512Mi
    defaultRequest:
      cpu: 100m
      memory: 128Mi
    type: Container
```

## Лучшие практики

1. Используйте namespaces для **разделения окружений** (dev, staging, prod)
2. Применяйте **ResourceQuotas** для контроля использования ресурсов
3. Настраивайте **RBAC** на уровне namespace
4. Используйте **labels** для дополнительной организации внутри namespace
5. **Не используйте** namespace `default` в production

## Примеры

- `namespace-dev.yaml` - development namespace с квотами
- `namespace-prod.yaml` - production namespace

