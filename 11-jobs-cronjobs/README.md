# 11. Jobs и CronJobs - Задачи и расписания

## Job

**Job** - создает один или несколько Pods и гарантирует, что указанное количество успешно завершится.

### Простой Job

```yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: pi-calculation
spec:
  template:
    spec:
      containers:
      - name: pi
        image: perl:5.34
        command: ["perl",  "-Mbignum=bpi", "-wle", "print bpi(2000)"]
      restartPolicy: Never
  backoffLimit: 4
```

### Параллельные Jobs

```yaml
spec:
  completions: 5      # Всего успешных завершений
  parallelism: 2      # Параллельно запущенных Pods
```

## CronJob

**CronJob** - создает Jobs по расписанию (как cron в Linux).

### Структура

```yaml
apiVersion: batch/v1
kind: CronJob
metadata:
  name: backup-job
spec:
  schedule: "0 2 * * *"    # Каждый день в 2:00 AM
  jobTemplate:
    spec:
      template:
        spec:
          containers:
          - name: backup
            image: backup-tool:latest
            command: ["/bin/sh", "-c", "backup.sh"]
          restartPolicy: OnFailure
```

### Cron формат

```
# ┌───────────── минута (0 - 59)
# │ ┌───────────── час (0 - 23)
# │ │ ┌───────────── день месяца (1 - 31)
# │ │ │ ┌───────────── месяц (1 - 12)
# │ │ │ │ ┌───────────── день недели (0 - 6) (Sunday=0)
# │ │ │ │ │
# * * * * *
```

Примеры:
- `0 * * * *` - каждый час
- `*/5 * * * *` - каждые 5 минут
- `0 2 * * *` - каждый день в 2:00
- `0 0 * * 0` - каждое воскресенье в полночь
- `0 0 1 * *` - первого числа каждого месяца

## Параметры конфигурации

### Job

```yaml
spec:
  backoffLimit: 4              # Количество retry при failure
  completions: 1               # Сколько успешных завершений нужно
  parallelism: 1               # Параллельных Pods
  activeDeadlineSeconds: 600   # Макс время выполнения
  ttlSecondsAfterFinished: 100 # Удалить Job через N секунд после завершения
```

### CronJob

```yaml
spec:
  schedule: "*/5 * * * *"
  successfulJobsHistoryLimit: 3   # Сколько успешных Jobs хранить
  failedJobsHistoryLimit: 1       # Сколько failed Jobs хранить
  concurrencyPolicy: Allow         # Allow, Forbid, Replace
  startingDeadlineSeconds: 200    # Дедлайн для старта Job
  suspend: false                   # Приостановить CronJob
```

## Concurrency Policy

- **Allow** (по умолчанию) - разрешить одновременное выполнение
- **Forbid** - запретить, если предыдущий Job еще работает
- **Replace** - заменить текущий Job новым

## Лучшие практики

1. Используйте **ttlSecondsAfterFinished** для автоочистки
2. Устанавливайте **activeDeadlineSeconds** для предотвращения зависаний
3. Для CronJob используйте **concurrencyPolicy: Forbid** если нельзя запускать параллельно
4. Делайте Jobs **идемпотентными** (безопасно перезапускать)
5. Используйте **initContainers** для подготовки окружения

## Примеры использования

**Jobs**:
- Batch обработка данных
- Database migrations
- Backup операции
- Одноразовые задачи

**CronJobs**:
- Регулярные бэкапы
- Очистка старых данных
- Отправка отчетов
- Периодическая синхронизация

## Примеры

- `job-simple.yaml` - Простой Job
- `job-parallel.yaml` - Параллельный Job
- `cronjob-backup.yaml` - CronJob для backup
- `cronjob-cleanup.yaml` - Периодическая очистка

