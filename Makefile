.EXPORT_ALL_VARIABLES:

.PHONY: build

#Enviroment variables used in application
ENV=dev
PORT=9011
DATABASE_URL=postgres://api:api@localhost:5433/api
DATABASE_URL_STACK=postgres://api:api@api-db:5432/api
NEGATIVE_TYPES=compra,saQue


#Migration folder
MIGRATION_FOLDER=db/migrations

build:
	go build cmd/server.go

run:
	CompileDaemon -exclude-dir ".git" -exclude-dir "postgres-data" --build="go build cmd/rest/server.go" --command=./server

stack:
	go test ./... -v
	go run cmd/rest/server.go


migration:
	docker-compose -f docker-compose.all.yml up db_migrations

test :
	go test ./... -race -coverprofile cp.out
	go tool cover -html=./cp.out -o cover.html

