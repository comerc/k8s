# Kubectl команды для работы с Deployments

## Создание Deployments

```bash
# Создать deployment из манифеста
kubectl apply -f deployment-simple.yaml

# Создать deployment императивно
kubectl create deployment nginx --image=nginx:1.25 --replicas=3

# С портом
kubectl create deployment nginx --image=nginx:1.25 --port=80 --replicas=3

# Создать все deployments из директории
kubectl apply -f .
```

## Просмотр Deployments

```bash
# Список deployments
kubectl get deployments
kubectl get deploy

# С дополнительной информацией
kubectl get deployments -o wide

# Детальная информация
kubectl describe deployment nginx-deployment

# YAML манифест
kubectl get deployment nginx-deployment -o yaml

# Следить за изменениями
kubectl get deployments --watch
kubectl get deployments -w
```

## Просмотр ReplicaSets

```bash
# Список ReplicaSets
kubectl get replicasets
kubectl get rs

# ReplicaSets конкретного deployment
kubectl get rs -l app=nginx

# Детальная информация
kubectl describe rs <replicaset-name>
```

## Масштабирование

```bash
# Увеличить количество реплик
kubectl scale deployment nginx-deployment --replicas=5

# Уменьшить количество реплик
kubectl scale deployment nginx-deployment --replicas=2

# Автомасштабирование (HPA - будет в разделе 08)
kubectl autoscale deployment nginx-deployment --min=2 --max=10 --cpu-percent=80

# Проверить текущее количество реплик
kubectl get deployment nginx-deployment -o jsonpath='{.spec.replicas}'
```

## Обновление образа

```bash
# Обновить образ контейнера
kubectl set image deployment/nginx-deployment nginx=nginx:1.26

# Обновить несколько контейнеров
kubectl set image deployment/nginx-deployment nginx=nginx:1.26 sidecar=busybox:1.36

# Обновить через edit
kubectl edit deployment nginx-deployment

# Обновить через apply (после изменения манифеста)
kubectl apply -f deployment-simple.yaml
```

## Rollout (управление обновлениями)

```bash
# Статус текущего обновления
kubectl rollout status deployment/nginx-deployment

# История обновлений
kubectl rollout history deployment/nginx-deployment

# Детали конкретной ревизии
kubectl rollout history deployment/nginx-deployment --revision=2

# Откат к предыдущей версии
kubectl rollout undo deployment/nginx-deployment

# Откат к конкретной ревизии
kubectl rollout undo deployment/nginx-deployment --to-revision=2

# Пауза обновления (для canary deployments)
kubectl rollout pause deployment/nginx-deployment

# Возобновление обновления
kubectl rollout resume deployment/nginx-deployment

# Перезапуск deployment (пересоздает все pods)
kubectl rollout restart deployment/nginx-deployment
```

## Редактирование

```bash
# Редактировать в текстовом редакторе
kubectl edit deployment nginx-deployment

# Применить изменения из файла
kubectl apply -f deployment-simple.yaml

# Заменить полностью
kubectl replace -f deployment-simple.yaml

# Patch (частичное обновление)
kubectl patch deployment nginx-deployment -p '{"spec":{"replicas":5}}'

# Patch из файла
kubectl patch deployment nginx-deployment --patch-file patch.yaml
```

## Управление ресурсами

```bash
# Установить resource limits
kubectl set resources deployment nginx-deployment \
  --limits=cpu=200m,memory=256Mi \
  --requests=cpu=100m,memory=128Mi

# Установить для конкретного контейнера
kubectl set resources deployment nginx-deployment \
  -c=nginx \
  --limits=cpu=200m,memory=256Mi
```

## Управление переменными окружения

```bash
# Добавить переменную окружения
kubectl set env deployment/nginx-deployment APP_VERSION=v1.0

# Удалить переменную
kubectl set env deployment/nginx-deployment APP_VERSION-

# Установить несколько переменных
kubectl set env deployment/nginx-deployment \
  APP_NAME=myapp \
  APP_VERSION=v1.0 \
  ENVIRONMENT=production

# Из ConfigMap
kubectl set env deployment/nginx-deployment --from=configmap/my-config

# Из Secret
kubectl set env deployment/nginx-deployment --from=secret/my-secret
```

## Удаление

```bash
# Удалить deployment
kubectl delete deployment nginx-deployment

# Удалить из манифеста
kubectl delete -f deployment-simple.yaml

# Удалить все deployments с меткой
kubectl delete deployments -l app=nginx

# Удалить deployment, но оставить pods (orphan)
kubectl delete deployment nginx-deployment --cascade=orphan
```

## Отладка

```bash
# События deployment
kubectl describe deployment nginx-deployment | grep -A 10 Events

# Логи всех pods в deployment
kubectl logs -l app=nginx --all-containers=true

# Логи с follow
kubectl logs -l app=nginx -f

# Статус pods
kubectl get pods -l app=nginx

# Проверка, почему pods не запускаются
kubectl get events --sort-by='.lastTimestamp' | grep nginx

# Использование ресурсов
kubectl top pods -l app=nginx
```

## Информация о deployment

```bash
# Количество реплик
kubectl get deployment nginx-deployment -o jsonpath='{.spec.replicas}'

# Количество доступных реплик
kubectl get deployment nginx-deployment -o jsonpath='{.status.availableReplicas}'

# Стратегия обновления
kubectl get deployment nginx-deployment -o jsonpath='{.spec.strategy}'

# Селектор
kubectl get deployment nginx-deployment -o jsonpath='{.spec.selector}'

# Образ контейнера
kubectl get deployment nginx-deployment -o jsonpath='{.spec.template.spec.containers[0].image}'

# Все в одной команде
kubectl get deployment nginx-deployment -o custom-columns=\
NAME:.metadata.name,\
REPLICAS:.spec.replicas,\
AVAILABLE:.status.availableReplicas,\
IMAGE:.spec.template.spec.containers[0].image
```

## Работа с метками

```bash
# Добавить метку к deployment
kubectl label deployment nginx-deployment env=production

# Изменить метку
kubectl label deployment nginx-deployment env=staging --overwrite

# Удалить метку
kubectl label deployment nginx-deployment env-

# Добавить метку к pods через deployment
kubectl patch deployment nginx-deployment -p \
  '{"spec":{"template":{"metadata":{"labels":{"monitored":"true"}}}}}'
```

## Работа с селекторами

```bash
# Получить deployments по метке
kubectl get deployments -l app=nginx

# Сложные селекторы
kubectl get deployments -l 'environment in (production,staging)'
kubectl get deployments -l 'app=nginx,environment=production'
kubectl get deployments -l 'app!=nginx'

# Получить pods, управляемые deployment
kubectl get pods -l app=nginx
```

## Экспорт и бэкап

```bash
# Экспорт в YAML
kubectl get deployment nginx-deployment -o yaml > nginx-backup.yaml

# Экспорт без cluster-specific полей
kubectl get deployment nginx-deployment -o yaml --export > nginx-clean.yaml

# Экспорт всех deployments
kubectl get deployments -o yaml > all-deployments.yaml

# Экспорт в JSON
kubectl get deployment nginx-deployment -o json > nginx.json
```

## Dry-run (проверка без применения)

```bash
# Проверить манифест без создания
kubectl apply -f deployment-simple.yaml --dry-run=client

# С выводом в YAML
kubectl apply -f deployment-simple.yaml --dry-run=client -o yaml

# Проверить на сервере (с валидацией)
kubectl apply -f deployment-simple.yaml --dry-run=server

# Создать манифест императивно без применения
kubectl create deployment nginx --image=nginx:1.25 --dry-run=client -o yaml > deployment.yaml
```

## Расширенные сценарии

```bash
# Получить все pods deployment и их статусы
kubectl get pods -l app=nginx -o custom-columns=\
NAME:.metadata.name,\
STATUS:.status.phase,\
NODE:.spec.nodeName,\
IP:.status.podIP

# Перезапустить deployment после изменения ConfigMap
kubectl rollout restart deployment/nginx-deployment

# Проверить, что обновление завершено
kubectl rollout status deployment/nginx-deployment --timeout=5m

# Получить историю с причинами изменений
kubectl rollout history deployment/nginx-deployment

# Записать причину изменения (для истории)
kubectl annotate deployment/nginx-deployment kubernetes.io/change-cause="Updated to nginx 1.26"

# Масштабировать несколько deployments
kubectl scale deployment nginx-deployment web-deployment --replicas=3

# Удалить все deployments в namespace
kubectl delete deployments --all

# Получить deployment в другом namespace
kubectl get deployment nginx-deployment -n production

# Скопировать deployment в другой namespace
kubectl get deployment nginx-deployment -o yaml | \
  sed 's/namespace: default/namespace: production/' | \
  kubectl apply -f -
```

## Мониторинг обновления

```bash
# Следить за процессом обновления
watch kubectl get pods -l app=nginx

# Следить за rollout статусом
kubectl rollout status deployment/nginx-deployment --watch

# Проверить события в реальном времени
kubectl get events --watch | grep nginx

# Следить за количеством реплик
watch 'kubectl get deployment nginx-deployment -o custom-columns=DESIRED:.spec.replicas,CURRENT:.status.replicas,AVAILABLE:.status.availableReplicas'
```

## Советы и трюки

```bash
# Быстрое создание deployment с нужными параметрами
kubectl create deployment nginx --image=nginx:1.25 --replicas=3 --dry-run=client -o yaml | \
  kubectl apply -f -

# Обновить и сразу следить за статусом
kubectl set image deployment/nginx-deployment nginx=nginx:1.26 && \
  kubectl rollout status deployment/nginx-deployment

# Автоматический откат при ошибке (через progressDeadlineSeconds в манифесте)
# Deployment автоматически откатится, если обновление не завершится за указанное время

# Проверить, что deployment healthy
kubectl get deployment nginx-deployment -o jsonpath='{.status.conditions[?(@.type=="Available")].status}'
# Ответ должен быть: True
```

