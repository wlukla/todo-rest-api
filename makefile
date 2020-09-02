dev:
	go run cmd/apiserver/main.go

test:
	go test -cover ./...