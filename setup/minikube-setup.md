# Настройка локального окружения с Minikube

## Что такое Minikube?

Minikube - это инструмент для запуска локального Kubernetes кластера на вашей машине. Идеально подходит для обучения и разработки.

## Установка

### macOS

```bash
# Установка через Homebrew
brew install minikube
brew install kubectl

# Проверка установки
minikube version
kubectl version --client
```

### Linux

```bash
# Установка minikube
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube

# Установка kubectl
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
```

### Windows

```powershell
# Используйте Chocolatey
choco install minikube
choco install kubernetes-cli
```

## Запуск Minikube

### Базовый запуск

```bash
# Запуск с настройками по умолчанию
minikube start

# Запуск с кастомными параметрами (рекомендуется)
minikube start \
  --cpus=2 \
  --memory=4096 \
  --driver=docker \
  --kubernetes-version=v1.28.0
```

### Выбор драйвера

Minikube поддерживает несколько драйверов:

```bash
# Docker (рекомендуется для macOS/Linux)
minikube start --driver=docker

# VirtualBox
minikube start --driver=virtualbox

# Hyper-V (Windows)
minikube start --driver=hyperv

# Посмотреть текущий драйвер
minikube config get driver
```

## Проверка установки

```bash
# Статус minikube
minikube status

# Информация о кластере
kubectl cluster-info

# Список нод
kubectl get nodes

# Версия Kubernetes
kubectl version

# Dashboard (опционально)
minikube dashboard
```

Вы должны увидеть что-то вроде:

```
NAME       STATUS   ROLES           AGE   VERSION
minikube   Ready    control-plane   1m    v1.28.0
```

## Полезные команды Minikube

```bash
# Остановить кластер (сохраняет данные)
minikube stop

# Удалить кластер
minikube delete

# Перезапустить кластер
minikube stop && minikube start

# SSH в ноду
minikube ssh

# Получить IP адрес кластера
minikube ip

# Открыть Kubernetes Dashboard
minikube dashboard

# Список аддонов
minikube addons list

# Включить аддон
minikube addons enable ingress
minikube addons enable metrics-server
```

## Настройка контекста kubectl

```bash
# Текущий контекст
kubectl config current-context

# Список контекстов
kubectl config get-contexts

# Переключение на minikube
kubectl config use-context minikube
```

## Работа с образами Docker

Для использования локальных Docker образов:

```bash
# Использовать Docker daemon из minikube
eval $(minikube docker-env)

# Теперь docker build будет создавать образы внутри minikube
docker build -t my-app:v1 .

# Вернуться к локальному Docker
eval $(minikube docker-env -u)
```

## Проброс портов

```bash
# Доступ к сервису через NodePort
minikube service <service-name>

# Получить URL сервиса
minikube service <service-name> --url

# Туннель для LoadBalancer сервисов
minikube tunnel
```

## Рекомендуемые аддоны для обучения

```bash
# Ingress контроллер
minikube addons enable ingress

# Metrics Server (для HPA и мониторинга)
minikube addons enable metrics-server

# Dashboard
minikube addons enable dashboard

# Registry (локальный Docker registry)
minikube addons enable registry
```

## Устранение проблем

### Проблема: Недостаточно ресурсов

```bash
# Удалите текущий кластер
minikube delete

# Запустите с большими ресурсами
minikube start --cpus=4 --memory=8192
```

### Проблема: Конфликт драйверов

```bash
# Явно укажите драйвер
minikube start --driver=docker --force
```

### Проблема: Образы не скачиваются

```bash
# Проверьте логи
minikube logs

# Используйте другой registry mirror (для России)
minikube start --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers
```

### Очистка и переустановка

```bash
# Полная очистка
minikube delete --all --purge

# Удаление конфигурации
rm -rf ~/.minikube

# Свежая установка
minikube start
```

## Проверка готовности

После установки выполните:

```bash
# 1. Проверьте статус
minikube status

# 2. Создайте тестовый под
kubectl run test-nginx --image=nginx --port=80

# 3. Проверьте, что под запустился
kubectl get pods

# 4. Удалите тестовый под
kubectl delete pod test-nginx
```

Если все команды выполнились успешно - вы готовы к обучению! 🎉

## Что дальше?

Переходите к разделу **01-basics** и начинайте изучать Pods!

