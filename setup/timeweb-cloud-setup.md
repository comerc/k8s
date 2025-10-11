# Развертывание в Timeweb Cloud

## О Timeweb Cloud Kubernetes

Timeweb Cloud предоставляет управляемый Kubernetes кластер. Это хороший вариант для деплоя учебных проектов в реальное облако.

## Создание кластера

### 1. Регистрация и вход

1. Зарегистрируйтесь на https://timeweb.cloud/
2. Войдите в панель управления
3. Перейдите в раздел "Kubernetes"

### 2. Создание кластера

1. Нажмите "Создать кластер"
2. Выберите параметры:
   - **Регион**: выберите ближайший к вам
   - **Версия Kubernetes**: последняя стабильная (1.28+)
   - **Пресет**: для обучения достаточно минимального
   - **Количество нод**: 2 ноды (можно начать с 1)

3. Нажмите "Создать"

Создание кластера занимает 5-10 минут.

## Подключение к кластеру

### 1. Скачивание kubeconfig

После создания кластера:

1. Откройте страницу кластера
2. Найдите раздел "Доступ"
3. Скачайте файл `kubeconfig`

### 2. Настройка kubectl

```bash
# Сохраните файл в ~/.kube/
mkdir -p ~/.kube
cp ~/Downloads/kubeconfig-timeweb.yaml ~/.kube/config-timeweb

# Установите переменную окружения
export KUBECONFIG=~/.kube/config-timeweb

# Или объедините с существующим конфигом
KUBECONFIG=~/.kube/config:~/.kube/config-timeweb kubectl config view --flatten > ~/.kube/config.tmp
mv ~/.kube/config.tmp ~/.kube/config
```

### 3. Проверка подключения

```bash
# Список контекстов
kubectl config get-contexts

# Переключение на Timeweb контекст
kubectl config use-context <timeweb-context-name>

# Проверка нод
kubectl get nodes

# Информация о кластере
kubectl cluster-info
```

## Переключение между minikube и Timeweb

```bash
# Использовать локальный minikube
kubectl config use-context minikube

# Использовать Timeweb Cloud
kubectl config use-context <timeweb-context-name>

# Текущий контекст
kubectl config current-context
```

## Деплой приложений

После подключения все команды kubectl работают так же, как и с minikube:

```bash
# Применить манифесты
kubectl apply -f deployment.yaml

# Проверить статус
kubectl get all

# Посмотреть логи
kubectl logs <pod-name>
```

## Особенности Timeweb Cloud

### LoadBalancer

В отличие от minikube, в Timeweb Cloud LoadBalancer работает автоматически:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-service
spec:
  type: LoadBalancer  # Автоматически создаст внешний IP
  ports:
    - port: 80
      targetPort: 8080
  selector:
    app: my-app
```

Получите внешний IP:

```bash
kubectl get svc my-service
```

### Persistent Volumes

Timeweb Cloud автоматически предоставляет storage класс:

```bash
# Список storage классов
kubectl get storageclass

# Использование в PVC
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: my-pvc
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 10Gi
  storageClassName: timeweb-cloud-default  # Или другой доступный класс
```

### Ingress

Для Ingress нужно установить контроллер:

```bash
# Установка Nginx Ingress Controller
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.8.1/deploy/static/provider/cloud/deploy.yaml

# Проверка установки
kubectl get pods -n ingress-nginx

# Получить внешний IP Ingress
kubectl get svc -n ingress-nginx
```

## Мониторинг расходов

```bash
# Проверка использования ресурсов
kubectl top nodes
kubectl top pods --all-namespaces

# Для экономии можно:
# 1. Уменьшить количество реплик
kubectl scale deployment my-app --replicas=1

# 2. Удалить неиспользуемые ресурсы
kubectl delete all --all -n default

# 3. Остановить кластер через панель управления (когда не используется)
```

## Полезные команды для Timeweb

```bash
# Проверка квот
kubectl describe quota

# Проверка лимитов
kubectl describe limits

# Список всех ресурсов
kubectl get all --all-namespaces

# Очистка всех ресурсов в namespace
kubectl delete all --all -n default
```

## Безопасность

### Не коммитьте kubeconfig в Git!

Добавьте в `.gitignore`:

```
# Kubernetes configs
kubeconfig*
*.kubeconfig
.kube/
```

### Используйте Secrets для паролей

```bash
# Создание Secret
kubectl create secret generic db-password \
  --from-literal=password='super-secret-password'

# Не храните пароли в манифестах!
```

## Удаление кластера

Когда закончите обучение:

1. Войдите в панель Timeweb Cloud
2. Перейдите в "Kubernetes"
3. Выберите ваш кластер
4. Нажмите "Удалить"

Или через API (если настроено):

```bash
# Удалить все ресурсы перед удалением кластера
kubectl delete all --all --all-namespaces
```

## Сравнение: Minikube vs Timeweb Cloud

| Аспект | Minikube | Timeweb Cloud |
|--------|----------|---------------|
| **Стоимость** | Бесплатно | Платно (от ~500₽/мес) |
| **Доступ из интернета** | Нет (только localhost) | Да (внешний IP) |
| **LoadBalancer** | Требует `minikube tunnel` | Работает автоматически |
| **Persistent Storage** | Локальный диск | Облачное хранилище |
| **Масштабирование** | Ограничено вашим ПК | Легко масштабируется |
| **Использование** | Обучение, разработка | Тестирование, продакшн |

## Рекомендуемый подход

1. **Обучение (темы 01-06)**: используйте minikube
2. **Практика (темы 07-12)**: продолжайте на minikube
3. **Финальный проект**: разверните multi-tier приложение в Timeweb Cloud

Это даст вам опыт работы и с локальным окружением, и с реальным облачным кластером!

## Полезные ссылки

- [Документация Timeweb Cloud Kubernetes](https://timeweb.cloud/docs/k8s)
- [API документация](https://timeweb.cloud/api-docs)
- [Панель управления](https://timeweb.cloud/my)

