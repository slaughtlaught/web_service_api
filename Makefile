MIGRATIONS_DIR := "migrations"
PG_DSN := "postgres://postgres:postgres@postgres:5432/notes?sslmode=disable"

migrate-up:
  migrate -verbose -path $(MIGRATIONS_DIR) -database $(PG_DSN) up

migrate-down:
  migrate -verbose -path $(MIGRATIONS_DIR) -database $(PG_DSN) down 1

migrate-version:
  migrate -verbose -path $(MIGRATIONS_DIR) -database $(PG_DSN) version

# Usage: make migrate-create name=migration_name
migrate-create:
  migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq -digits 6 $(name)
  ls -t $(MIGRATIONS_DIR)/*.sql | head -2 | while read line; do \
    echo "BEGIN;\n\n\n\nCOMMIT;" > $$line; \
  done

migrate-install:
  go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest