include .env
migrateup:
	migrate  -path db/migration -database "mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(127.0.0.1:${MYSQL_PORT})/${MYSQL_DATABASE}" -verbose up
migratedown:
	migrate  -path db/migration -database "mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(127.0.0.1:${MYSQL_PORT})/${MYSQL_DATABASE}" -verbose down
.PHONY: migrationup migratedown