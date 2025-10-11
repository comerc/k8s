# 04. ConfigMaps и Secrets - Управление конфигурацией

## ConfigMap

**ConfigMap** - объект для хранения неконфиденциальных данных конфигурации в виде пар ключ-значение.

### Создание ConfigMap

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: app-config
data:
  database_url: "postgres://db:5432"
  app_name: "MyApp"
  app_mode: "production"
```

### Использование в Pod

**1. Как переменные окружения:**

```yaml
env:
- name: DATABASE_URL
  valueFrom:
    configMapKeyRef:
      name: app-config
      key: database_url
```

**2. Все ключи как env vars:**

```yaml
envFrom:
- configMapRef:
    name: app-config
```

**3. Как файлы (volume):**

```yaml
volumes:
- name: config
  configMap:
    name: app-config
containers:
- volumeMounts:
  - name: config
    mountPath: /etc/config
```

## Secret

**Secret** - объект для хранения конфиденциальных данных (пароли, токены, ключи).

### Создание Secret

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: db-secret
type: Opaque
data:
  username: YWRtaW4=        # base64 encoded
  password: cGFzc3dvcmQ=    # base64 encoded
```

**Кодирование:**
```bash
echo -n 'admin' | base64
echo -n 'password' | base64
```

### Использование Secret

```yaml
env:
- name: DB_USERNAME
  valueFrom:
    secretKeyRef:
      name: db-secret
      key: username
- name: DB_PASSWORD
  valueFrom:
    secretKeyRef:
      name: db-secret
      key: password
```

## Типы Secrets

- **Opaque** - произвольные данные (по умолчанию)
- **kubernetes.io/tls** - TLS сертификаты
- **kubernetes.io/dockerconfigjson** - Docker registry credentials
- **kubernetes.io/basic-auth** - Basic аутентификация
- **kubernetes.io/ssh-auth** - SSH ключи

## ConfigMap vs Secret

| Аспект | ConfigMap | Secret |
|--------|-----------|--------|
| Данные | Неконфиденциальные | Конфиденциальные |
| Кодирование | Plaintext | Base64 |
| Безопасность | Нет особой защиты | Может быть зашифрован at rest |
| Использование | Конфиги, параметры | Пароли, ключи, токены |

## Лучшие практики

1. **Не коммитьте Secrets в Git** (используйте .gitignore)
2. **Используйте внешние системы** для production (Vault, AWS Secrets Manager)
3. **Ограничьте RBAC** доступ к Secrets
4. **Обновляйте Secrets** без перезапуска Pods (volume mount)
5. **Не используйте Secret для действительно чувствительных данных** в production без дополнительного шифрования

## Файлы в этой директории

- `configmap-env.yaml` - ConfigMap как env vars
- `configmap-volume.yaml` - ConfigMap как файлы
- `secret-env.yaml` - Secret как env vars
- `secret-tls.yaml` - TLS Secret

