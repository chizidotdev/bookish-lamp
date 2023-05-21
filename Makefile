postgres:
	sudo docker run --name copia -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	sudo docker exec -it copia createdb --username=root --owner=root copia

dropdb:
	sudo docker exec -it copia dropdb copia

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/copia?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/copia?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test
