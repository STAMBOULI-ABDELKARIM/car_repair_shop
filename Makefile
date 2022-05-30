postgres:
	 docker run --name postgres14car -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=root -d postgres:14-alpine
createdb:
	docker exec -it postgres14car createdb -U postgres --username=postgres --owner=postgres car_repair_shop

dropdb:
	docker exec -it postgres14car dropdb -U postgres bank
migrateup:
	migrate -path db/migration -database "postgresql://postgres:root@localhost:5432/car_repair_shop?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://postgres:root@localhost:5432/car_repair_shop?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go

.PHONY:	postgres createdb dropdb migrateup migratedown sqlc test server
