.PHONY: run
run:
	go run main.go

# https://github.com/golang-migrate/migrate
# create a migration file for versions of the database.
# db changes
.PHONY: migrate
migrate:
	migrate create -ext sql -dir db/migration -seq init_schema

postgres:
	docker run --name postgres-container -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:12-alpine 

createdb:
	docker exec -it postgres-container createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres-container dropdb simple_bank

# apply migration to the postgres database
migrateup:
	migrate -path db/migration/ -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration/ -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown