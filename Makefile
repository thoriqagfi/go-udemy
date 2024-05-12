# make ${target} DB_CONTAINER=${your_db_container} DB_USERNAME=${your_db_username} DB_NAME=${your_db_name}

postgres:
	docker run --name postgres-udemy -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=123456789 -p 5432:5432 -d postgres:alpine

createdb:
	docker exec -it postgresql createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it postgresql dropdb --username=postgres simple_bank

# Path: Mafkefile
migrateup:
	migrate -path db/migration -database "postgresql://postgres:123456789@127.0.0.1/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:123456789@127.0.0.1:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: createdb dropdb postgres migrateup migratedown sqlc test server