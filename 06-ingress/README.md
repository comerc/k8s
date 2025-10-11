# 06. Ingress - Маршрутизация HTTP/HTTPS трафика

## Что такое Ingress?

**Ingress** - объект API, который управляет внешним доступом к сервисам в кластере, обычно HTTP/HTTPS.

### Возможности Ingress

- ✅ Маршрутизация по URL (path-based routing)
- ✅ Маршрутизация по hostname (host-based routing)
- ✅ SSL/TLS терминация
- ✅ Единая точка входа для множества сервисов
- ✅ Load balancing

## Структура Ingress

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: my-ingress
spec:
  rules:
  - host: myapp.example.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: my-service
            port:
              number: 80
```

## Ingress Controller

Ingress сам по себе ничего не делает - нужен **Ingress Controller**:
- **Nginx Ingress Controller** (самый популярный)
- **Traefik**
- **HAProxy**
- **Istio**

### Установка в minikube

```bash
minikube addons enable ingress
```

## Типы маршрутизации

### 1. По hostname

```yaml
rules:
- host: app1.example.com
  http:
    paths:
    - path: /
      backend:
        service:
          name: app1-service
- host: app2.example.com
  http:
    paths:
    - path: /
      backend:
        service:
          name: app2-service
```

### 2. По path

```yaml
rules:
- host: example.com
  http:
    paths:
    - path: /api
      backend:
        service:
          name: api-service
    - path: /web
      backend:
        service:
          name: web-service
```

## TLS/SSL

```yaml
spec:
  tls:
  - hosts:
    - myapp.example.com
    secretName: tls-secret
  rules:
  - host: myapp.example.com
    http:
      paths:
      - path: /
        backend:
          service:
            name: my-service
```

## Примеры

- `ingress-simple.yaml` - базовый Ingress
- `ingress-multiple-hosts.yaml` - несколько хостов
- `ingress-tls.yaml` - с SSL/TLS

