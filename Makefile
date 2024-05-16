run:
	go run main.go




# https://github.com/golang-migrate/migrate
# create a migration file for versions of the database.
# db changes
migrate:
	migrate create -ext sql -dir db/migration -seq init_schema


postgres:
	docker run --name postgres-container -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:12-alpine 

enter-shell:
	docker exec -it postgres-container psql -U root

createdb:
	docker exec -it postgres-container createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres-container dropdb simple_bank

# apply migration to the postgres database itself (000001_init_schema.up.sql migration file)
migrateup:
	migrate -path db/migration/ -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration/ -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

# apply migration to the postgres database itself (000001_init_schema.down.sql migration file)
migratedown:
	migrate -path db/migration/ -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" -verbose down

# 'down 1' states that we only want to roll back one last migration that was applied before
migratedown1:
	migrate -path db/migration/ -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" -verbose down 1 

sqlc:
	sqlc generate

test:
	go test -v ./... -coverprofile=coverage.out 
	go tool cover -html=coverage.out

mock:
	mockgen -package mockdb -destination db/mock/store.go RyanFin/GoSimpleBank/db/sqlc Store

.PHONY: run migrate postgres createdb dropdb migrateup migratedown sqlc test mock migrateup1 migratedown1

# multi-curl command, replace URL with amd.tar.gz present at this URL: https://github.com/golang-migrate/migrate/releases
# https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md

