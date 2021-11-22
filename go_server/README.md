REDIS_URL=localhost REDIS_PORT=6379 RPC_PORT=50051 HTTP_PORT=8080 go run main.go
REDIS_URL=localhost REDIS_PORT=6379 RPC_PORT=50051 HTTP_PORT=8080 ./main
docker build -f Dockerfile .


curl localhost:8080/fib -d '{"number": 13}'
