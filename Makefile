run-local:
	go run ./main.go -env local

build:
	go build ./cmd/api/main.go

test:
	go test -cover ./...

run-pods:
	echo "Start local environment"
	docker-compose -f docker-compose.local.yml up --build
