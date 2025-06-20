md

# Concurrent API with Goroutines in Go
This project demonstrates how to use Goroutines in Go to create a concurrent API. Goroutines are lightweight threads that can run concurrently with


## Features

- 🧵 Concurrency using Goroutines
- 📡 API endpoint that simulates calling multiple services in parallel
- 📦 Go Modules support
- 🐳 Dockerized for easy deployment
- ⚙️ CI/CD with GitHub Actions

## How It Works

When you hit the `/concurrent` endpoint, two fake services are executed concurrently using Goroutines, and results are returned faster.

## Run Locally

```bash
git clone https://github.com/yourusername/concurrent-api.git
cd concurrent-api
go mod tidy
go run main.go

Visit: http://localhost:8080/concurrent

Run with Docker

docker build -t concurrent-api .
docker run -p 8080:8080 concurrent-api