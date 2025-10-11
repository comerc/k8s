# Kubectl команды для работы с Services

## Создание Services

```bash
# Создать Service из манифеста
kubectl apply -f service-clusterip.yaml

# Создать ClusterIP Service императивно
kubectl expose deployment nginx-deployment --port=80 --target-port=80

# Создать NodePort Service
kubectl expose deployment nginx-deployment --type=NodePort --port=80

# Создать LoadBalancer Service
kubectl expose deployment nginx-deployment --type=LoadBalancer --port=80

# С именем
kubectl expose deployment nginx-deployment --name=my-service --port=80
```

## Просмотр Services

```bash
# Список Services
kubectl get services
kubectl get svc

# С дополнительной информацией
kubectl get svc -o wide

# Детальная информация
kubectl describe service nginx-service

# YAML манифест
kubectl get svc nginx-service -o yaml

# Следить за изменениями
kubectl get svc --watch
```

## Просмотр Endpoints

```bash
# Список Endpoints
kubectl get endpoints
kubectl get ep

# Endpoints конкретного Service
kubectl get endpoints nginx-service

# Детальная информация
kubectl describe endpoints nginx-service
```

## Тестирование Service

```bash
# Получить ClusterIP
kubectl get svc nginx-service -o jsonpath='{.spec.clusterIP}'

# Получить порт
kubectl get svc nginx-service -o jsonpath='{.spec.ports[0].port}'

# Получить NodePort
kubectl get svc nginx-nodeport -o jsonpath='{.spec.ports[0].nodePort}'

# Тест из другого Pod
kubectl run test-pod --image=curlimages/curl -it --rm -- curl http://nginx-service

# Port-forward для локального доступа
kubectl port-forward service/nginx-service 8080:80
# Теперь: http://localhost:8080
```

## Доступ к NodePort Service

```bash
# Получить IP minikube
minikube ip

# Получить URL NodePort Service
minikube service nginx-nodeport --url

# Открыть в браузере
minikube service nginx-nodeport
```

## Доступ к LoadBalancer Service

```bash
# В minikube нужен tunnel
minikube tunnel
# (В другом терминале)

# Получить External IP
kubectl get svc nginx-loadbalancer

# Когда External IP назначен
curl http://<EXTERNAL-IP>
```

## DNS

```bash
# Тест DNS из Pod
kubectl run dns-test --image=busybox:1.36 -it --rm -- nslookup nginx-service

# Полное DNS имя
# <service-name>.<namespace>.svc.cluster.local
kubectl run dns-test --image=busybox:1.36 -it --rm -- nslookup nginx-service.default.svc.cluster.local

# Получить все DNS записи сервисов
kubectl run dns-test --image=busybox:1.36 -it --rm -- nslookup kubernetes.default.svc.cluster.local
```

## Редактирование

```bash
# Редактировать Service
kubectl edit service nginx-service

# Применить изменения из файла
kubectl apply -f service-clusterip.yaml

# Patch
kubectl patch service nginx-service -p '{"spec":{"type":"NodePort"}}'
```

## Удаление

```bash
# Удалить Service
kubectl delete service nginx-service

# Удалить из манифеста
kubectl delete -f service-clusterip.yaml

# Удалить все Services с меткой
kubectl delete services -l app=nginx
```

## Отладка

```bash
# Проверить, какие Pods за Service
kubectl get endpoints nginx-service

# Если Endpoints пустые:
# 1. Проверить селектор Service
kubectl get svc nginx-service -o jsonpath='{.spec.selector}'

# 2. Проверить метки Pods
kubectl get pods --show-labels

# 3. Проверить, что Pods готовы
kubectl get pods

# События Service
kubectl describe service nginx-service | grep Events

# Логи Pods за Service
kubectl logs -l app=nginx --all-containers=true
```

## Информация о Service

```bash
# ClusterIP
kubectl get svc nginx-service -o jsonpath='{.spec.clusterIP}'

# Type
kubectl get svc nginx-service -o jsonpath='{.spec.type}'

# Ports
kubectl get svc nginx-service -o jsonpath='{.spec.ports}'

# Selector
kubectl get svc nginx-service -o jsonpath='{.spec.selector}'

# External IP (LoadBalancer)
kubectl get svc nginx-loadbalancer -o jsonpath='{.status.loadBalancer.ingress[0].ip}'

# Все в одной команде
kubectl get svc -o custom-columns=\
NAME:.metadata.name,\
TYPE:.spec.type,\
CLUSTER-IP:.spec.clusterIP,\
EXTERNAL-IP:.status.loadBalancer.ingress[0].ip,\
PORT:.spec.ports[0].port
```

## Тестирование связности

```bash
# Создать test pod
kubectl run test-connectivity --image=nicolaka/netshoot -it --rm -- bash

# Внутри Pod:
# Тест ClusterIP
curl http://<cluster-ip>

# Тест DNS
curl http://nginx-service
curl http://nginx-service.default.svc.cluster.local

# Тест конкретного порта
curl http://nginx-service:80

# Telnet
telnet nginx-service 80

# Dig DNS
dig nginx-service.default.svc.cluster.local
```

## Работа с метками

```bash
# Получить Services по метке
kubectl get svc -l app=nginx

# Добавить метку
kubectl label service nginx-service env=production

# Удалить метку
kubectl label service nginx-service env-
```

## Session Affinity

```bash
# Добавить Session Affinity
kubectl patch service nginx-service -p '{"spec":{"sessionAffinity":"ClientIP"}}'

# Настроить timeout
kubectl patch service nginx-service -p \
  '{"spec":{"sessionAffinityConfig":{"clientIP":{"timeoutSeconds":10800}}}}'

# Удалить Session Affinity
kubectl patch service nginx-service -p '{"spec":{"sessionAffinity":"None"}}'
```

## Экспорт

```bash
# Экспорт в YAML
kubectl get svc nginx-service -o yaml > service-backup.yaml

# Экспорт всех Services
kubectl get svc -o yaml > all-services.yaml
```

## Полезные комбинации

```bash
# Получить все Pods за Service
kubectl get pods -l $(kubectl get svc nginx-service -o jsonpath='{.spec.selector}' | jq -r 'to_entries[] | "\(.key)=\(.value)"' | tr '\n' ',')

# Или проще через Endpoints
kubectl get endpoints nginx-service -o jsonpath='{.subsets[*].addresses[*].targetRef.name}'

# Проверить доступность всех Endpoints
for ip in $(kubectl get endpoints nginx-service -o jsonpath='{.subsets[*].addresses[*].ip}'); do
  echo "Testing $ip"
  curl -s http://$ip
done

# Следить за External IP LoadBalancer
kubectl get svc nginx-loadbalancer --watch -o jsonpath='{.status.loadBalancer.ingress[0].ip}'

# Быстрый тест Service из временного Pod
kubectl run tmp --image=curlimages/curl --rm -it --restart=Never -- curl http://nginx-service
```

