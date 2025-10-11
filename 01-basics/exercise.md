# Практическое задание: Работа с Pods

## Цель
Закрепить знания о создании, управлении и отладке Pods в Kubernetes.

## Предварительные требования

- Запущен minikube: `./scripts/start-minikube.sh`
- kubectl настроен и работает: `kubectl get nodes`

## Задание 1: Создание простого Pod

1. Создайте pod с именем `my-nginx` используя образ `nginx:1.25`
2. Убедитесь, что pod запустился
3. Получите IP адрес пода
4. Пробросьте порт 8080 на локальной машине к порту 80 в поде
5. Откройте браузер и проверьте, что nginx работает

<details>
<summary>Подсказка</summary>

```bash
kubectl run my-nginx --image=nginx:1.25 --port=80
kubectl get pods
kubectl get pod my-nginx -o wide
kubectl port-forward my-nginx 8080:80
# В браузере: http://localhost:8080
```
</details>

## Задание 2: Работа с манифестами

1. Примените все манифесты из текущей директории
2. Проверьте, что все pods запустились
3. Найдите pod с меткой `app=web`
4. Посмотрите детальную информацию о `multi-container-pod`

<details>
<summary>Подсказка</summary>

```bash
kubectl apply -f .
kubectl get pods
kubectl get pods -l app=web
kubectl describe pod multi-container-pod
```
</details>

## Задание 3: Логи и выполнение команд

1. Получите логи из `pod-with-env`
2. Выполните команду `env` внутри этого пода и найдите переменную `APP_NAME`
3. В `multi-container-pod` выведите логи контейнера `log-sidecar`
4. Войдите в интерактивную оболочку контейнера `nginx` в `multi-container-pod`

<details>
<summary>Подсказка</summary>

```bash
kubectl logs pod-with-env
kubectl exec pod-with-env -- env | grep APP_NAME
kubectl logs multi-container-pod -c log-sidecar
kubectl exec -it multi-container-pod -c nginx -- /bin/bash
# Внутри: ls /var/log/nginx
# Выйти: exit
```
</details>

## Задание 4: Создание кастомного Pod

Создайте файл `my-custom-pod.yaml` со следующими требованиями:

1. Имя: `redis-cache`
2. Образ: `redis:7-alpine`
3. Метки:
   - `app`: `cache`
   - `tier`: `backend`
4. Переменная окружения: `REDIS_MAX_MEMORY=256mb`
5. Порт: `6379`
6. Resource requests: `memory: 128Mi, cpu: 100m`
7. Resource limits: `memory: 256Mi, cpu: 200m`

После создания:
- Примените манифест
- Проверьте, что pod работает
- Выполните команду `redis-cli ping` внутри пода
- Посмотрите использование ресурсов

<details>
<summary>Решение</summary>

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: redis-cache
  labels:
    app: cache
    tier: backend
spec:
  containers:
  - name: redis
    image: redis:7-alpine
    ports:
    - containerPort: 6379
    env:
    - name: REDIS_MAX_MEMORY
      value: "256mb"
    resources:
      requests:
        memory: "128Mi"
        cpu: "100m"
      limits:
        memory: "256Mi"
        cpu: "200m"
```

```bash
kubectl apply -f my-custom-pod.yaml
kubectl get pod redis-cache
kubectl exec redis-cache -- redis-cli ping
kubectl top pod redis-cache
```
</details>

## Задание 5: Отладка проблемного Pod

Создайте pod с ошибкой и научитесь её находить:

```bash
kubectl run broken-pod --image=nginx:wrong-tag
```

1. Посмотрите статус пода
2. Узнайте, почему под не запускается
3. Посмотрите события
4. Исправьте проблему

<details>
<summary>Подсказка</summary>

```bash
kubectl get pod broken-pod
# Статус будет: ErrImagePull или ImagePullBackOff

kubectl describe pod broken-pod
# Смотрим секцию Events

kubectl get events --sort-by='.lastTimestamp'

# Исправление: удалить и создать с правильным образом
kubectl delete pod broken-pod
kubectl run broken-pod --image=nginx:1.25
```
</details>

## Задание 6: Многоконтейнерный Pod

Создайте pod с двумя контейнерами:

1. Основной контейнер:
   - Образ: `busybox:1.36`
   - Команда: `while true; do echo "$(date) - Hello from main" >> /data/app.log; sleep 5; done`
   - Volume mount: `/data`

2. Sidecar контейнер:
   - Образ: `busybox:1.36`
   - Команда: `tail -f /data/app.log`
   - Volume mount: `/data`

Общий volume типа `emptyDir`

После создания:
- Проверьте логи обоих контейнеров
- Убедитесь, что sidecar видит логи основного контейнера

<details>
<summary>Решение</summary>

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: app-with-logger
spec:
  containers:
  - name: app
    image: busybox:1.36
    command: ['sh', '-c', 'while true; do echo "$(date) - Hello from main" >> /data/app.log; sleep 5; done']
    volumeMounts:
    - name: data
      mountPath: /data
  
  - name: logger
    image: busybox:1.36
    command: ['sh', '-c', 'tail -f /data/app.log']
    volumeMounts:
    - name: data
      mountPath: /data
  
  volumes:
  - name: data
    emptyDir: {}
```

```bash
kubectl apply -f app-with-logger.yaml
kubectl logs app-with-logger -c app
kubectl logs app-with-logger -c logger -f
```
</details>

## Задание 7: Работа с метками

1. Создайте 3 пода с разными метками:
   ```bash
   kubectl run web-1 --image=nginx --labels="app=web,env=prod,version=v1"
   kubectl run web-2 --image=nginx --labels="app=web,env=dev,version=v2"
   kubectl run db-1 --image=postgres:15 --labels="app=database,env=prod,version=v1"
   ```

2. Найдите все pods с меткой `app=web`
3. Найдите все pods с метками `env=prod` И `app=web`
4. Найдите все pods с версией `v1`
5. Добавьте метку `monitored=true` ко всем pods
6. Удалите все pods с меткой `env=dev`

<details>
<summary>Подсказка</summary>

```bash
kubectl get pods -l app=web
kubectl get pods -l app=web,env=prod
kubectl get pods -l version=v1
kubectl label pods --all monitored=true
kubectl delete pods -l env=dev
```
</details>

## Задание 8: Проверка знаний

Ответьте на вопросы:

1. В чем разница между `requests` и `limits` для ресурсов?
2. Можно ли изменить образ в работающем поде?
3. Что произойдет, если контейнер в поде упадет?
4. Зачем нужны метки (labels)?
5. Какой командой можно войти в интерактивную оболочку пода?

<details>
<summary>Ответы</summary>

1. `requests` - гарантированные ресурсы, которые кластер резервирует. `limits` - максимальные ресурсы, которые контейнер может использовать.

2. Нет, нельзя. Нужно удалить pod и создать новый. Для этого используются Deployments (следующий раздел).

3. Kubernetes попытается перезапустить контейнер согласно `restartPolicy` (по умолчанию Always).

4. Метки используются для организации, идентификации и выбора ресурсов. Они критически важны для Services и Deployments.

5. `kubectl exec -it <pod-name> -- /bin/bash` или `/bin/sh`
</details>

## Очистка

После выполнения заданий удалите все созданные ресурсы:

```bash
kubectl delete pod my-nginx redis-cache app-with-logger web-1 web-2 db-1 broken-pod 2>/dev/null
kubectl delete -f .
```

Или используйте скрипт:

```bash
../../scripts/cleanup.sh
```

## Критерии успеха

- ✅ Вы понимаете, что такое Pod и для чего он нужен
- ✅ Умеете создавать Pods из манифестов и императивно
- ✅ Можете просматривать логи и выполнять команды в поде
- ✅ Понимаете концепцию многоконтейнерных pods
- ✅ Умеете работать с метками и селекторами
- ✅ Знаете основные команды kubectl для работы с pods

## Что дальше?

Переходите к разделу **02-deployments** для изучения управления множеством pods!

