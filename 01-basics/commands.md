# Kubectl команды для работы с Pods

## Создание Pods

```bash
# Создать pod из манифеста
kubectl apply -f pod-simple.yaml

# Создать pod императивно (без манифеста)
kubectl run nginx --image=nginx:1.25 --port=80

# Создать pod с переменными окружения
kubectl run busybox --image=busybox --env="APP_NAME=myapp" --env="VERSION=1.0"

# Создать несколько pods
kubectl apply -f .
```

## Просмотр Pods

```bash
# Список всех pods
kubectl get pods

# Список pods с дополнительной информацией
kubectl get pods -o wide

# Список pods с метками
kubectl get pods --show-labels

# Список pods по селектору
kubectl get pods -l app=nginx
kubectl get pods -l environment=dev,tier=frontend

# Детальная информация о поде
kubectl describe pod nginx-simple

# Информация в формате YAML
kubectl get pod nginx-simple -o yaml

# Информация в формате JSON
kubectl get pod nginx-simple -o json

# Только имена pods
kubectl get pods -o name

# Следить за изменениями в реальном времени
kubectl get pods --watch
kubectl get pods -w
```

## Логи

```bash
# Показать логи пода
kubectl logs nginx-simple

# Логи конкретного контейнера (в многоконтейнерном поде)
kubectl logs multi-container-pod -c nginx
kubectl logs multi-container-pod -c log-sidecar

# Логи в реальном времени
kubectl logs -f nginx-simple

# Последние N строк логов
kubectl logs --tail=50 nginx-simple

# Логи за последний час
kubectl logs --since=1h nginx-simple

# Логи с временными метками
kubectl logs --timestamps nginx-simple

# Логи предыдущего запуска пода (если был рестарт)
kubectl logs --previous nginx-simple
```

## Выполнение команд в поде

```bash
# Выполнить команду в поде
kubectl exec nginx-simple -- ls /usr/share/nginx/html

# Интерактивная оболочка
kubectl exec -it nginx-simple -- /bin/bash
kubectl exec -it nginx-simple -- /bin/sh

# В конкретном контейнере
kubectl exec -it multi-container-pod -c nginx -- /bin/bash

# Выполнить несколько команд
kubectl exec nginx-simple -- sh -c "ls -la && pwd"
```

## Проброс портов

```bash
# Пробросить порт с пода на локальную машину
kubectl port-forward nginx-simple 8080:80

# Теперь можно открыть http://localhost:8080 в браузере

# Пробросить на определенный адрес
kubectl port-forward nginx-simple 8080:80 --address 0.0.0.0

# В фоновом режиме
kubectl port-forward nginx-simple 8080:80 &
```

## Копирование файлов

```bash
# Скопировать файл ИЗ пода
kubectl cp nginx-simple:/etc/nginx/nginx.conf ./nginx.conf

# Скопировать файл В под
kubectl cp ./index.html nginx-simple:/usr/share/nginx/html/index.html

# Для многоконтейнерного пода
kubectl cp multi-container-pod:/logs/access.log ./access.log -c log-sidecar
```

## Редактирование

```bash
# Редактировать pod (откроется редактор)
kubectl edit pod nginx-simple

# Применить изменения из файла
kubectl apply -f pod-simple.yaml

# Заменить ресурс
kubectl replace -f pod-simple.yaml --force
```

## Удаление Pods

```bash
# Удалить pod
kubectl delete pod nginx-simple

# Удалить из манифеста
kubectl delete -f pod-simple.yaml

# Удалить все pods с меткой
kubectl delete pods -l app=nginx

# Удалить все pods в namespace
kubectl delete pods --all

# Принудительное удаление (не рекомендуется)
kubectl delete pod nginx-simple --force --grace-period=0
```

## Отладка

```bash
# События в кластере
kubectl get events

# События, связанные с подом
kubectl get events --field-selector involvedObject.name=nginx-simple

# Статус и причина проблем
kubectl describe pod nginx-simple | grep -A 10 Events

# Проверка использования ресурсов (требует metrics-server)
kubectl top pod nginx-simple
kubectl top pods

# Подробная информация о контейнерах
kubectl get pod nginx-simple -o jsonpath='{.spec.containers[*].name}'

# Статус контейнеров
kubectl get pod nginx-simple -o jsonpath='{.status.containerStatuses[*].state}'
```

## Фильтрация и сортировка

```bash
# По статусу
kubectl get pods --field-selector=status.phase=Running
kubectl get pods --field-selector=status.phase!=Running

# Сортировка по времени создания
kubectl get pods --sort-by=.metadata.creationTimestamp

# Сортировка по имени
kubectl get pods --sort-by=.metadata.name

# Кастомные столбцы
kubectl get pods -o custom-columns=NAME:.metadata.name,STATUS:.status.phase,IP:.status.podIP
```

## Работа с метками

```bash
# Добавить метку
kubectl label pod nginx-simple env=production

# Изменить метку
kubectl label pod nginx-simple env=staging --overwrite

# Удалить метку
kubectl label pod nginx-simple env-

# Показать все метки
kubectl get pods --show-labels
```

## Работа с аннотациями

```bash
# Добавить аннотацию
kubectl annotate pod nginx-simple description="Nginx web server"

# Удалить аннотацию
kubectl annotate pod nginx-simple description-

# Посмотреть аннотации
kubectl describe pod nginx-simple | grep Annotations
```

## Полезные комбинации

```bash
# Получить IP всех pods
kubectl get pods -o wide | awk '{print $6}'

# Получить имена всех running pods
kubectl get pods --field-selector=status.phase=Running -o name

# Перезапустить pod (удалить и создать заново)
kubectl delete pod nginx-simple && kubectl apply -f pod-simple.yaml

# Проверить, что pod работает
kubectl wait --for=condition=ready pod/nginx-simple --timeout=60s

# Получить логи всех pods с определенной меткой
for pod in $(kubectl get pods -l app=nginx -o name); do
  echo "=== $pod ==="
  kubectl logs $pod
done
```

## Шпаргалка по форматам вывода

```bash
-o wide           # Дополнительные столбцы (IP, Node)
-o yaml           # YAML формат
-o json           # JSON формат
-o name           # Только имена
-o jsonpath       # JSONPath выражения
-o custom-columns # Кастомные столбцы
--show-labels     # Показать метки
--watch           # Следить за изменениями
```

## Контекст и Namespace

```bash
# Текущий namespace
kubectl config view --minify --output 'jsonpath={..namespace}'

# Указать namespace явно
kubectl get pods -n kube-system

# Установить namespace по умолчанию
kubectl config set-context --current --namespace=default

# Pods во всех namespaces
kubectl get pods --all-namespaces
kubectl get pods -A
```

