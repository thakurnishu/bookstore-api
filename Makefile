
db-start:
	@docker run -p 5432:5432 -d --name bookstore-postgres -e POSTGRES_DB=bookstore -e POSTGRES_USER=bookstore-user -e POSTGRES_PASSWORD=bookstore-password  postgres:16.1-alpine3.19
db-stop:
	@docker stop bookstore-postgres
	@docker rm bookstore-postgres 

migrate-up:
	@migrate -path storage/migration -database "postgresql://bookstore-user:bookstore-password@localhost:5432/bookstore?sslmode=disable" -verbose up

migrate-down:
	@migrate -path storage/migration -database "postgresql://bookstore-user:bookstore-password@localhost:5432/bookstore?sslmode=disable" -verbose down 

build:
	@go build -o bin/bookstore-api

run:  build
	@./bin/bookstore-api

test: 
	@go test -v ./...