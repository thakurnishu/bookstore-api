postgres-start:
	@docker run -p 5432:5432 -d --name bookstore-postgres -e POSTGRES_DB=bookstore -e POSTGRES_USER=bookstore-user -e POSTGRES_PASSWORD=bookstore-password  postgres:16.1-alpine3.19
	
postgres-stop:
	@docker stop bookstore-postgres
	@docker rm bookstore-postgres 

mysql-start:
	@docker run -p 3306:3306 -d --name bookstore-mysql -e MYSQL_ROOT_PASSWORD=bookstore-password -e MYSQL_DATABASE=bookstore -e MYSQL_USER=bookstore-user -e MYSQL_PASSWORD=bookstore-password mysql:latest

mysql-stop:
	@docker stop bookstore-mysql
	@docker rm bookstore-mysql

build:
	@go build -o bin/bookstore-api

run:  build
	@./bin/bookstore-api

test: 
	@go test -v ./...