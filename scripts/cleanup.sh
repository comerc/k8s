#!/bin/bash

# Скрипт для очистки всех учебных ресурсов из кластера

set -e

echo "🧹 Очистка учебных ресурсов из Kubernetes..."
echo ""

# Проверка подключения к кластеру
if ! kubectl cluster-info &> /dev/null; then
    echo "❌ Нет подключения к Kubernetes кластеру!"
    echo "Запустите minikube: ./scripts/start-minikube.sh"
    exit 1
fi

echo "Текущий контекст: $(kubectl config current-context)"
echo ""

# Предупреждение
read -p "⚠️  Удалить ВСЕ ресурсы из namespace 'default'? (y/N): " -n 1 -r
echo ""

if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo "Отменено."
    exit 0
fi

echo ""
echo "🗑️  Удаление ресурсов..."
echo ""

# Удаление всех ресурсов в namespace default
echo "  - Deployments"
kubectl delete deployments --all -n default 2>/dev/null || echo "    Нет deployments"

echo "  - Services (кроме kubernetes)"
kubectl delete services --all -n default --field-selector metadata.name!=kubernetes 2>/dev/null || echo "    Нет services"

echo "  - Pods"
kubectl delete pods --all -n default 2>/dev/null || echo "    Нет pods"

echo "  - ConfigMaps"
kubectl delete configmaps --all -n default 2>/dev/null || echo "    Нет configmaps"

echo "  - Secrets (кроме default)"
kubectl delete secrets --all -n default --field-selector type!=kubernetes.io/service-account-token 2>/dev/null || echo "    Нет secrets"

echo "  - PersistentVolumeClaims"
kubectl delete pvc --all -n default 2>/dev/null || echo "    Нет PVCs"

echo "  - Ingress"
kubectl delete ingress --all -n default 2>/dev/null || echo "    Нет ingress"

echo "  - StatefulSets"
kubectl delete statefulsets --all -n default 2>/dev/null || echo "    Нет statefulsets"

echo "  - Jobs"
kubectl delete jobs --all -n default 2>/dev/null || echo "    Нет jobs"

echo "  - CronJobs"
kubectl delete cronjobs --all -n default 2>/dev/null || echo "    Нет cronjobs"

echo ""
echo "✅ Очистка завершена!"
echo ""

# Показать оставшиеся ресурсы
echo "📊 Оставшиеся ресурсы:"
kubectl get all -n default

echo ""
echo "💡 Для полного удаления кластера выполните: minikube delete"

