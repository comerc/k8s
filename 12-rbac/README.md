# 12. RBAC - Role-Based Access Control

## Что такое RBAC?

**RBAC** - система управления доступом на основе ролей. Определяет, кто (Subject) может делать что (Verb) с какими ресурсами (Resource).

## Основные компоненты

### 1. Role и ClusterRole

**Role** - права доступа в рамках namespace
**ClusterRole** - права доступа на уровне кластера

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: pod-reader
  namespace: default
rules:
- apiGroups: [""]
  resources: ["pods"]
  verbs: ["get", "list", "watch"]
```

### 2. RoleBinding и ClusterRoleBinding

Привязывает Role к пользователям/группам/service accounts.

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: read-pods
  namespace: default
subjects:
- kind: User
  name: jane
  apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: Role
  name: pod-reader
  apiGroup: rbac.authorization.k8s.io
```

### 3. ServiceAccount

Идентификатор для Pods (не для людей).

```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: my-service-account
  namespace: default
```

## Verbs (действия)

- **get** - получить один ресурс
- **list** - получить список ресурсов
- **watch** - следить за изменениями
- **create** - создать ресурс
- **update** - обновить ресурс
- **patch** - частично обновить
- **delete** - удалить ресурс
- **deletecollection** - удалить коллекцию

## API Groups

- **""** (core) - pods, services, configmaps, secrets
- **apps** - deployments, statefulsets, daemonsets
- **batch** - jobs, cronjobs
- **networking.k8s.io** - ingresses
- **rbac.authorization.k8s.io** - roles, rolebindings

## Примеры ролей

### Read-only доступ к Pods

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: pod-reader
rules:
- apiGroups: [""]
  resources: ["pods", "pods/log"]
  verbs: ["get", "list", "watch"]
```

### Управление Deployments

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: deployment-manager
rules:
- apiGroups: ["apps"]
  resources: ["deployments"]
  verbs: ["get", "list", "watch", "create", "update", "patch", "delete"]
```

### Полный доступ в namespace

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: namespace-admin
rules:
- apiGroups: ["*"]
  resources: ["*"]
  verbs: ["*"]
```

## ServiceAccount в Pod

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: my-pod
spec:
  serviceAccountName: my-service-account
  containers:
  - name: app
    image: nginx:1.25
```

## Проверка прав доступа

```bash
# Проверить, может ли текущий пользователь создавать pods
kubectl auth can-i create pods

# Проверить для другого пользователя
kubectl auth can-i create deployments --as=jane

# Проверить для service account
kubectl auth can-i list secrets --as=system:serviceaccount:default:my-sa

# Список всех прав
kubectl auth can-i --list
```

## Встроенные ClusterRoles

Kubernetes имеет встроенные роли:
- **cluster-admin** - полный доступ ко всему
- **admin** - полный доступ в namespace
- **edit** - чтение/запись в namespace
- **view** - только чтение в namespace

```bash
# Посмотреть встроенные роли
kubectl get clusterroles
```

## Лучшие практики

1. **Principle of Least Privilege** - минимальные необходимые права
2. Используйте **ServiceAccounts** для Pods, а не default
3. Не используйте **cluster-admin** без необходимости
4. Создавайте **отдельные роли** для разных команд/приложений
5. Регулярно **аудируйте** права доступа
6. Используйте **namespace** для изоляции

## Примеры

- `role-pod-reader.yaml` - Read-only роль для pods
- `role-deployment-manager.yaml` - Управление deployments
- `serviceaccount-with-role.yaml` - ServiceAccount с привязкой роли

