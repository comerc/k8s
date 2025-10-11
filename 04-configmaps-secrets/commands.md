# Команды для ConfigMaps и Secrets

## ConfigMap

```bash
# Создать из литералов
kubectl create configmap app-config --from-literal=key1=value1 --from-literal=key2=value2

# Создать из файла
kubectl create configmap app-config --from-file=config.properties

# Создать из директории
kubectl create configmap app-config --from-file=config-dir/

# Просмотр
kubectl get configmaps
kubectl describe configmap app-config
kubectl get configmap app-config -o yaml

# Редактирование
kubectl edit configmap app-config

# Удаление
kubectl delete configmap app-config
```

## Secret

```bash
# Создать из литералов
kubectl create secret generic db-secret --from-literal=username=admin --from-literal=password=secret

# Создать из файла
kubectl create secret generic ssh-key --from-file=~/.ssh/id_rsa

# TLS Secret
kubectl create secret tls my-tls --cert=cert.pem --key=key.pem

# Docker registry Secret
kubectl create secret docker-registry my-registry \
  --docker-server=registry.example.com \
  --docker-username=user \
  --docker-password=password \
  --docker-email=user@example.com

# Просмотр
kubectl get secrets
kubectl describe secret db-secret
kubectl get secret db-secret -o yaml

# Декодирование
kubectl get secret db-secret -o jsonpath='{.data.password}' | base64 --decode

# Редактирование
kubectl edit secret db-secret

# Удаление
kubectl delete secret db-secret
```

