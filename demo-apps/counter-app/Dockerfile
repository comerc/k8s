# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Копируем go mod файлы
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходники
COPY main.go ./

# Компилируем
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o counter-app .

# Runtime stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируем бинарник из builder
COPY --from=builder /app/counter-app .

EXPOSE 8080

CMD ["./counter-app"]

