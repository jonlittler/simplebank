# Makefile

# Go
server:
	go run main.go
	
test:
	go test -v -cover ./...

# Database
postgres: 
	docker run --name postgres13 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:13-alpine

createdb:
	docker exec -it postgres13 createdb --username=root --owner=root simplebank

dropdb:
	docker exec -it postgres13 dropdb simplebank

migrateup:
	migrate -path db/migration -database "postgres://root:password@localhost:5432/simplebank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgres://root:password@localhost:5432/simplebank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgres://root:password@localhost:5432/simplebank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgres://root:password@localhost:5432/simplebank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

mock:
	mockgen -package mockdb -destination ./db/mock/store.go github.com/jonlittler/ts/simplebank/db/sqlc Store 

.PHONY: server test postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc mock