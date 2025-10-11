# Практическое задание: Работа с Deployments

## Цель
Научиться создавать, масштабировать, обновлять и откатывать Deployments.

## Предварительные требования

- Запущен minikube
- Пройден раздел 01-basics

## Задание 1: Создание простого Deployment

1. Создайте deployment с именем `web-app`:
   - Образ: `nginx:1.24`
   - Реплики: 3
   - Метка: `app=web`

2. Проверьте, что создались:
   - 1 Deployment
   - 1 ReplicaSet
   - 3 Pods

3. Удалите один из pods и посмотрите, что произойдет

<details>
<summary>Подсказка</summary>

```bash
kubectl create deployment web-app --image=nginx:1.24 --replicas=3
kubectl get deployments
kubectl get rs
kubectl get pods
kubectl delete pod <pod-name>
kubectl get pods --watch
```
</details>

## Задание 2: Масштабирование

1. Увеличьте количество реплик до 5
2. Проверьте, что появились новые pods
3. Уменьшите до 2 реплик
4. Убедитесь, что лишние pods удалены

<details>
<summary>Подсказка</summary>

```bash
kubectl scale deployment web-app --replicas=5
kubectl get pods
kubectl scale deployment web-app --replicas=2
kubectl get pods
```
</details>

## Задание 3: Обновление приложения (Rolling Update)

1. Обновите образ на `nginx:1.25`
2. Следите за процессом обновления в реальном времени
3. Проверьте историю обновлений
4. Убедитесь, что все pods используют новый образ

<details>
<summary>Подсказка</summary>

```bash
kubectl set image deployment/web-app nginx=nginx:1.25
kubectl rollout status deployment/web-app
kubectl get pods --watch
kubectl rollout history deployment/web-app
kubectl get pods -o custom-columns=NAME:.metadata.name,IMAGE:.spec.containers[0].image
```
</details>

## Задание 4: Откат (Rollback)

1. Обновите образ на несуществующую версию `nginx:wrong`
2. Посмотрите, что произойдет с pods
3. Откатитесь к предыдущей рабочей версии
4. Убедитесь, что приложение снова работает

<details>
<summary>Подсказка</summary>

```bash
kubectl set image deployment/web-app nginx=nginx:wrong
kubectl get pods
kubectl describe pod <failing-pod-name>
kubectl rollout undo deployment/web-app
kubectl rollout status deployment/web-app
kubectl get pods
```
</details>

## Задание 5: Deployment из манифеста

Создайте файл `my-deployment.yaml` со следующими требованиями:

1. Имя: `backend-api`
2. Реплики: 4
3. Образ: `httpd:2.4`
4. Стратегия: RollingUpdate
   - maxSurge: 1
   - maxUnavailable: 0 (zero downtime)
5. Resource limits:
   - requests: cpu=100m, memory=128Mi
   - limits: cpu=200m, memory=256Mi
6. Метки:
   - app: api
   - tier: backend
   - environment: dev

<details>
<summary>Решение</summary>

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: backend-api
  labels:
    app: api
spec:
  replicas: 4
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxSurge: 1
      maxUnavailable: 0
  selector:
    matchLabels:
      app: api
      tier: backend
  template:
    metadata:
      labels:
        app: api
        tier: backend
        environment: dev
    spec:
      containers:
      - name: httpd
        image: httpd:2.4
        ports:
        - containerPort: 80
        resources:
          requests:
            cpu: "100m"
            memory: "128Mi"
          limits:
            cpu: "200m"
            memory: "256Mi"
```

```bash
kubectl apply -f my-deployment.yaml
kubectl get deployment backend-api
kubectl describe deployment backend-api
```
</details>

## Задание 6: Эксперимент с стратегиями

1. Создайте deployment `app-recreate` со стратегией `Recreate`:
   ```bash
   kubectl apply -f deployment-recreate.yaml
   ```

2. Обновите образ с `nginx:1.24` на `nginx:1.25`:
   ```bash
   kubectl set image deployment/app-recreate app=nginx:1.25
   ```

3. Следите за процессом обновления:
   ```bash
   kubectl get pods --watch
   ```

4. Обратите внимание на разницу с RollingUpdate:
   - Все pods удаляются сразу
   - Потом создаются новые
   - Есть период, когда нет ни одного running pod (downtime!)

## Задание 7: Паттерн Canary Deployment

Canary deployment - это техника, когда новая версия сначала развертывается для небольшого процента пользователей.

1. Создайте deployment с v1:
   ```bash
   kubectl create deployment app-v1 --image=nginx:1.24 --replicas=9
   kubectl label deployment app-v1 version=v1
   ```

2. Создайте deployment с v2 (canary):
   ```bash
   kubectl create deployment app-v2 --image=nginx:1.25 --replicas=1
   kubectl label deployment app-v2 version=v2
   ```

3. Оба deployment должны иметь одинаковые метки для Service (будет в следующем разделе):
   ```bash
   kubectl patch deployment app-v1 -p '{"spec":{"template":{"metadata":{"labels":{"app":"myapp"}}}}}'
   kubectl patch deployment app-v2 -p '{"spec":{"template":{"metadata":{"labels":{"app":"myapp"}}}}}'
   ```

4. Теперь 90% трафика идет на v1, 10% на v2 (canary)

5. Если v2 работает хорошо, масштабируйте:
   ```bash
   kubectl scale deployment app-v1 --replicas=5
   kubectl scale deployment app-v2 --replicas=5
   ```

6. Затем полностью переключитесь:
   ```bash
   kubectl scale deployment app-v1 --replicas=0
   kubectl scale deployment app-v2 --replicas=10
   ```

## Задание 8: Работа с историей

1. Создайте deployment:
   ```bash
   kubectl create deployment history-app --image=nginx:1.23
   ```

2. Сделайте несколько обновлений с комментариями:
   ```bash
   kubectl set image deployment/history-app nginx=nginx:1.24
   kubectl annotate deployment/history-app kubernetes.io/change-cause="Updated to 1.24"
   
   kubectl set image deployment/history-app nginx=nginx:1.25
   kubectl annotate deployment/history-app kubernetes.io/change-cause="Updated to 1.25"
   
   kubectl set image deployment/history-app nginx=nginx:1.26
   kubectl annotate deployment/history-app kubernetes.io/change-cause="Updated to 1.26"
   ```

3. Посмотрите историю:
   ```bash
   kubectl rollout history deployment/history-app
   ```

4. Откатитесь к версии 1.24:
   ```bash
   kubectl rollout history deployment/history-app --revision=1
   kubectl rollout undo deployment/history-app --to-revision=1
   ```

## Задание 9: Пауза и возобновление обновления

1. Создайте deployment с несколькими репликами:
   ```bash
   kubectl create deployment pause-demo --image=nginx:1.24 --replicas=5
   ```

2. Начните обновление и сразу поставьте на паузу:
   ```bash
   kubectl set image deployment/pause-demo nginx=nginx:1.25
   kubectl rollout pause deployment/pause-demo
   ```

3. Проверьте состояние (часть pods будет обновлена, часть нет):
   ```bash
   kubectl get pods -l app=pause-demo -o custom-columns=NAME:.metadata.name,IMAGE:.spec.containers[0].image
   ```

4. Возобновите обновление:
   ```bash
   kubectl rollout resume deployment/pause-demo
   kubectl rollout status deployment/pause-demo
   ```

## Задание 10: Диагностика проблем

Создайте deployment с проблемами и научитесь их находить:

```bash
kubectl create deployment broken-app --image=nonexistent/image:latest --replicas=3
```

1. Посмотрите статус deployment
2. Проверьте pods
3. Найдите причину проблемы
4. Исправьте образ
5. Проверьте, что все заработало

<details>
<summary>Подсказка</summary>

```bash
kubectl get deployment broken-app
kubectl get pods -l app=broken-app
kubectl describe pod <pod-name>
# Увидите: ErrImagePull или ImagePullBackOff

# Исправление
kubectl set image deployment/broken-app image=nginx:1.25
kubectl rollout status deployment/broken-app
kubectl get pods -l app=broken-app
```
</details>

## Задание 11: Мониторинг ресурсов

1. Создайте deployment с ограничениями ресурсов:
   ```bash
   kubectl apply -f deployment-nginx.yaml
   ```

2. Проверьте использование ресурсов:
   ```bash
   kubectl top pods -l app=nginx
   ```

3. Измените resource limits:
   ```bash
   kubectl set resources deployment nginx-full \
     --limits=cpu=500m,memory=512Mi \
     --requests=cpu=200m,memory=256Mi
   ```

4. Убедитесь, что pods пересозданы с новыми лимитами:
   ```bash
   kubectl get pods -l app=nginx --watch
   kubectl describe pod <pod-name> | grep -A 5 "Limits\|Requests"
   ```

## Бонус: Blue-Green Deployment

Создайте blue-green deployment паттерн:

1. Blue (текущая версия):
   ```bash
   kubectl create deployment blue --image=nginx:1.24 --replicas=3
   kubectl label deployment blue version=blue
   ```

2. Green (новая версия):
   ```bash
   kubectl create deployment green --image=nginx:1.25 --replicas=3
   kubectl label deployment green version=green
   ```

3. Service указывает на blue (будет в следующем разделе)

4. Протестируйте green

5. Переключите Service на green

6. Если всё ОК, удалите blue:
   ```bash
   kubectl delete deployment blue
   ```

## Очистка

```bash
kubectl delete deployment web-app backend-api app-recreate history-app pause-demo broken-app
kubectl delete deployment app-v1 app-v2 blue green
../../scripts/cleanup.sh
```

## Критерии успеха

- ✅ Понимаете разницу между Pod, ReplicaSet и Deployment
- ✅ Умеете создавать Deployments из манифестов и императивно
- ✅ Можете масштабировать приложения
- ✅ Понимаете Rolling Update и знаете, как его настроить
- ✅ Умеете откатывать обновления
- ✅ Знаете, как работать с историей изменений
- ✅ Понимаете разные стратегии деплоя (Canary, Blue-Green)

## Что дальше?

Теперь вы умеете управлять приложениями! Переходите к разделу **03-services** для изучения сетевого взаимодействия между pods.

