# Практическое задание: Работа с Services

## Задание 1: ClusterIP Service

1. Создайте deployment `web` с nginx (3 реплики)
2. Создайте ClusterIP Service для него
3. Создайте test pod и проверьте доступность через DNS
4. Проверьте load balancing (запросите несколько раз)

<details>
<summary>Решение</summary>

```bash
kubectl create deployment web --image=nginx:1.25 --replicas=3
kubectl expose deployment web --port=80
kubectl run test --image=curlimages/curl -it --rm -- sh
# В test pod:
# curl http://web
# exit
```
</details>

## Задание 2: NodePort Service

1. Создайте NodePort Service для deployment web
2. Получите URL через minikube
3. Откройте в браузере

<details>
<summary>Решение</summary>

```bash
kubectl expose deployment web --type=NodePort --port=80 --name=web-nodeport
minikube service web-nodeport --url
minikube service web-nodeport  # Откроет браузер
```
</details>

## Задание 3: Multi-port Service

Создайте приложение с несколькими портами и Service для него.

## Очистка

```bash
kubectl delete deployment web
kubectl delete svc web web-nodeport
../../scripts/cleanup.sh
```

