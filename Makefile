build:
	@go build -o bin/rest-api cmd/main.go

test:
	@go test -v ./...

run-build: build
	@./bin/rest-api

run: 
	@go run cmd/main.go

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down
