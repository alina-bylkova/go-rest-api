dev:
	go run main.go

dev-env:
	env REDIS=1 GIN_MODE=release go run main.go
