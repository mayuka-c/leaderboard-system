postgres:
	docker run --name dev-postgres16 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mysecretpassword -e POSTGRES_DB=leaderboard-system -d -p 5432:5432 postgres:16-alpine

createdb:
	docker exec -it dev-postgres16 createdb --username=root --owner=root leaderboard-system

dropdb:
	docker exec -it dev-postgres16 dropdb leaderboard-system

migrateup:
	migrate -path internal/pkg/db/migration -database "postgresql://root:mysecretpassword@localhost:5432/leaderboard-system?sslmode=disable" --verbose up

migratedown:
	migrate -path internal/pkg/db/migration -database "postgresql://root:mysecretpassword@localhost:5432/leaderboard-system?sslmode=disable" --verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc