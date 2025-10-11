#!/bin/bash

# Скрипт для запуска Minikube с оптимальными настройками для обучения

set -e

echo "🚀 Запуск Minikube для K8s обучения..."

# Проверка установки minikube
if ! command -v minikube &> /dev/null; then
    echo "❌ Minikube не установлен!"
    echo "Установите его: brew install minikube"
    exit 1
fi

# Проверка установки kubectl
if ! command -v kubectl &> /dev/null; then
    echo "❌ kubectl не установлен!"
    echo "Установите его: brew install kubectl"
    exit 1
fi

# Проверка, запущен ли уже minikube
if minikube status &> /dev/null; then
    echo "✅ Minikube уже запущен"
    minikube status
    exit 0
fi

# Запуск minikube с оптимальными параметрами
echo "⏳ Запуск minikube (это может занять 2-3 минуты)..."
minikube start \
    --cpus=2 \
    --memory=4096 \
    --driver=docker \
    --kubernetes-version=stable

echo ""
echo "✅ Minikube успешно запущен!"
echo ""

# Включение полезных аддонов
echo "📦 Включение полезных аддонов..."

# Metrics Server - для мониторинга и HPA
echo "  - metrics-server (для kubectl top и HPA)"
minikube addons enable metrics-server

# Ingress - для работы с Ingress в разделе 06
echo "  - ingress (для раздела 06-ingress)"
minikube addons enable ingress

# Dashboard - опционально, для визуального интерфейса
echo "  - dashboard (опционально, для GUI)"
minikube addons enable dashboard

echo ""
echo "✅ Аддоны включены!"
echo ""

# Информация о кластере
echo "📊 Информация о кластере:"
echo ""
kubectl cluster-info
echo ""
kubectl get nodes
echo ""

# Полезные команды
echo "💡 Полезные команды:"
echo ""
echo "  kubectl get nodes              # Список нод"
echo "  kubectl get pods               # Список подов"
echo "  kubectl get all                # Все ресурсы"
echo "  minikube dashboard             # Открыть Dashboard"
echo "  minikube stop                  # Остановить кластер"
echo "  minikube delete                # Удалить кластер"
echo ""

echo "🎉 Всё готово! Можно начинать обучение!"
echo "Перейдите в директорию 01-basics и следуйте инструкциям."

