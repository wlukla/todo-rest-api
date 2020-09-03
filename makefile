dev:
	go run cmd/apiserver/main.go

test:
	go test -cover ./...

migrate-up:
	migrate -path migrations -database "postgres://localhost/todo-db-dev?sslmode=disable" up

migrate-down:
	migrate -path migrations -database "postgres://localhost/todo-db-dev?sslmode=disable" down