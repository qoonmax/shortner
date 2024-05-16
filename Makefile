create_migration:
	goose -dir ./internal/storage/sqlite/migrations sqlite3 ./sqlite_data/sqlite create $(name) sql

migrate_up:
	goose -dir ./internal/storage/sqlite/migrations sqlite3 ./sqlite_data/sqlite up

migrate_down:
	goose -dir ./internal/storage/sqlite/migrations sqlite3 ./sqlite_data/sqlite down