# 05. Volumes - Хранение данных

## Проблема с контейнерами

Контейнеры эфемерны - данные теряются при перезапуске. Volumes решают эту проблему.

## Типы Volumes

### 1. emptyDir
Временное хранилище, существует пока жив Pod.

```yaml
volumes:
- name: cache
  emptyDir: {}
```

**Использование**: кеш, временные файлы, обмен данными между контейнерами

### 2. hostPath
Монтирует директорию с ноды.

```yaml
volumes:
- name: host-data
  hostPath:
    path: /data
    type: Directory
```

**Использование**: доступ к файлам ноды, логирование

### 3. PersistentVolume (PV) и PersistentVolumeClaim (PVC)

**PersistentVolume** - ресурс хранилища в кластере
**PersistentVolumeClaim** - запрос на хранилище от пользователя

```yaml
# PersistentVolumeClaim
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: my-pvc
spec:
  accessModes:
  - ReadWriteOnce
  resources:
    requests:
      storage: 1Gi
```

```yaml
# Использование в Pod
volumes:
- name: data
  persistentVolumeClaim:
    claimName: my-pvc
```

## Access Modes

- **ReadWriteOnce (RWO)** - чтение/запись одной нодой
- **ReadOnlyMany (ROX)** - чтение многими нодами
- **ReadWriteMany (RWX)** - чтение/запись многими нодами

## StorageClass

Динамическое создание PV:

```yaml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: fast-storage
provisioner: kubernetes.io/aws-ebs
parameters:
  type: gp2
```

## Лучшие практики

1. Используйте **PVC** вместо hostPath в production
2. Для stateful приложений используйте **StatefulSet** с PVC
3. Настраивайте **backup** для важных данных
4. Используйте **StorageClass** для автоматизации

## Примеры в этой директории

- `volume-emptydir.yaml` - временное хранилище
- `volume-hostpath.yaml` - монтирование с ноды
- `pvc-simple.yaml` - PersistentVolumeClaim
- `pod-with-pvc.yaml` - Pod с PVC

