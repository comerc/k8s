package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var (
	ctx         = context.Background()
	redisClient *redis.Client
	hostname    string
)

const htmlTemplate = `
<!DOCTYPE html>
<html lang="ru">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Counter App - K8s Demo (Go)</title>
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; }
        body {
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
            display: flex;
            justify-content: center;
            align-items: center;
            color: white;
        }
        .container {
            text-align: center;
            padding: 50px;
            background: rgba(255, 255, 255, 0.1);
            border-radius: 20px;
            backdrop-filter: blur(10px);
            box-shadow: 0 8px 32px 0 rgba(31, 38, 135, 0.37);
            max-width: 600px;
        }
        h1 { font-size: 2.5em; margin-bottom: 30px; }
        .counter {
            font-size: 5em;
            margin: 30px 0;
            font-weight: bold;
            color: #FFD700;
            text-shadow: 2px 2px 4px rgba(0,0,0,0.3);
        }
        .buttons {
            display: flex;
            gap: 20px;
            justify-content: center;
            margin: 30px 0;
        }
        button {
            padding: 15px 30px;
            font-size: 1.2em;
            border: none;
            border-radius: 10px;
            cursor: pointer;
            transition: transform 0.2s, box-shadow 0.2s;
            font-weight: bold;
        }
        button:hover { transform: translateY(-2px); box-shadow: 0 4px 12px rgba(0,0,0,0.3); }
        button:active { transform: translateY(0); }
        .increment { background: #4CAF50; color: white; }
        .reset { background: #f44336; color: white; }
        .info {
            margin-top: 30px;
            padding: 20px;
            background: rgba(255, 255, 255, 0.1);
            border-radius: 10px;
            font-size: 0.9em;
            text-align: left;
        }
        .info p { margin: 5px 0; }
        .go-badge {
            display: inline-block;
            background: #00ADD8;
            color: white;
            padding: 5px 15px;
            border-radius: 20px;
            font-size: 0.8em;
            margin: 10px 0;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>🔢 Counter App</h1>
        <div class="go-badge">⚡ Powered by Go</div>
        <p>Демонстрация работы с Redis в Kubernetes</p>
        <div class="counter">{{.Count}}</div>
        <div class="buttons">
            <button class="increment" onclick="increment()">➕ Увеличить</button>
            <button class="reset" onclick="reset()">🔄 Сброс</button>
        </div>
        <div class="info">
            <p><strong>Pod:</strong> {{.Hostname}}</p>
            <p><strong>Redis:</strong> {{.RedisHost}}:{{.RedisPort}}</p>
            <p><strong>Язык:</strong> Go {{.GoVersion}}</p>
            <p><strong>Версия:</strong> v1.0.0</p>
        </div>
    </div>
    <script>
        function increment() {
            fetch('/increment', { method: 'POST' })
                .then(() => location.reload());
        }
        function reset() {
            fetch('/reset', { method: 'POST' })
                .then(() => location.reload());
        }
    </script>
</body>
</html>
`

type PageData struct {
	Count     string
	Hostname  string
	RedisHost string
	RedisPort string
	GoVersion string
}

func init() {
	var err error
	hostname, err = os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = "redis"
	}

	redisPort := os.Getenv("REDIS_PORT")
	if redisPort == "" {
		redisPort = "6379"
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", redisHost, redisPort),
	})
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/increment", handleIncrement)
	http.HandleFunc("/reset", handleReset)
	http.HandleFunc("/health", handleHealth)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Counter App starting on port %s", port)
	log.Printf("📡 Redis connection: %s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	count, err := redisClient.Get(ctx, "counter").Result()
	if err == redis.Nil {
		count = "0"
	} else if err != nil {
		http.Error(w, "Redis connection error", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.New("index").Parse(htmlTemplate))
	data := PageData{
		Count:     count,
		Hostname:  hostname,
		RedisHost: os.Getenv("REDIS_HOST"),
		RedisPort: os.Getenv("REDIS_PORT"),
		GoVersion: "1.21",
	}

	tmpl.Execute(w, data)
}

func handleIncrement(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := redisClient.Incr(ctx, "counter").Err()
	if err != nil {
		http.Error(w, "Redis error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func handleReset(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := redisClient.Set(ctx, "counter", 0, 0).Err()
	if err != nil {
		http.Error(w, "Redis error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	err := redisClient.Ping(ctx).Err()
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, `{"status":"unhealthy","error":"%s"}`, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"status":"healthy"}`)
}

