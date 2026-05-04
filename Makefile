# Makefile

# Go
server:
	go run main.go
	
test:
	go test -v -cover ./...

# Database
postgres: 
	docker run --name postgres13 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:13-alpine

createdb:
	docker exec -it postgres13 createdb --username=root --owner=root simplebank

dropdb:
	docker exec -it postgres13 dropdb simplebank

migrateup:
	migrate -path db/migration -database "postgres://root:password@localhost:5432/simplebank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://root:password@localhost:5432/simplebank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

mock:
	mockgen -package mockdb -destination ./db/mock/store.go github.com/jonlittler/ts/simplebank/db/sqlc Store 

.PHONY: server test postgres createdb dropdb migrateup migratedown sqlc mock