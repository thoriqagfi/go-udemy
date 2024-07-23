# make ${target} DB_CONTAINER=${your_db_container} DB_USERNAME=${your_db_username} DB_NAME=${your_db_name}

postgres:
	docker run --name postgres-udemy --network bank-network -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=123456789 -p 5432:5432 -d postgres:alpine

createdb:
	docker exec -it postgresql createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it postgresql dropdb --username=postgres simple_bank

# Path: Mafkefile
migrateup:
	migrate -path db/migration -database "postgresql://postgres:123456789@127.0.0.1/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://postgres:123456789@127.0.0.1/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://postgres:123456789@127.0.0.1:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://postgres:123456789@127.0.0.1:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go go-udemy.sqlc.dev/app/db/sqlc Store

docker-run:
	docker run --name simplebank --network bank-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://postgres:123456789@postgresql/simple_bank?sslmode=disable" simplebank:latest

.PHONY: createdb dropdb postgres migrateup migratedown sqlc test server mock migrateup1 migratedown1