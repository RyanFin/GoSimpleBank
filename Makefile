
DB_URL=postgresql://root:root@localhost:5432/simple_bank?sslmode=disable

run:
	go run main.go

# https://github.com/golang-migrate/migrate
# create a migration file for versions of the database.
# db changes
migrate:
	migrate create -ext sql -dir db/migration -seq init_schema

# create a new tables. Firstly by generating new migration files
#  Change the table name at the end of the command
migrate-create:
	migrate create -ext sql -dir db/migration -seq add_sessions

postgres-start:
	docker network create bank-network
	docker run --name postgres-container --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:16-alpine

enter-shell:
	docker exec -it postgres-container psql -U root

createdb:
	docker exec -it postgres-container createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres-container dropdb simple_bank

# apply migration to the postgres database itself (000001_init_schema.up.sql migration file)
migrateup:
	migrate -path db/migration/ -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration/ -database "$(DB_URL)" -verbose up 1

# apply migration to the postgres database itself (000001_init_schema.down.sql migration file)
migratedown:
	migrate -path db/migration/ -database "$(DB_URL)" -verbose down

# 'down 1' states that we only want to roll back one last migration that was applied before
migratedown1:
	migrate -path db/migration/ -database "$(DB_URL)" -verbose down 1 

sqlc:
	sqlc generate

test:
	go test -v ./... -coverprofile=coverage.out 
	go tool cover -html=coverage.out

mock:
	mockgen -package mockdb -destination db/mock/store.go RyanFin/GoSimpleBank/db/sqlc Store

# build API documentation using DBDocs at the location of the specified db.dbml file
docs-build:
	dbdocs build doc/db.dbml

dbml-schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

proto:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto


.PHONY: run migrate postgres-start createdb dropdb migrateup migratedown sqlc test mock migrateup1 migratedown1 migrate-create docs-build dbml-schema proto

# multi-curl command, replace URL with amd.tar.gz present at this URL: https://github.com/golang-migrate/migrate/releases
# https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md