# make ${target} DB_CONTAINER=${your_db_container} DB_USERNAME=${your_db_username} DB_NAME=${your_db_name}

postgres:
	docker run --name postgres-udemy -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:alpine

createdb:
	docker exec -it postgres-udemy createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres-udemy dropdb simple_bank

# Path: Makefile
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: createdb dropdb postgres migrateup migratedown sqlc test