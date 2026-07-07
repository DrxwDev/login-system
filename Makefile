include .env

migrations:
	@goose -dir=database/migrations/schema create ${name} sql

mu: migrate/up
migrate/up:
	@echo "Running up migrations..."
	@goose -dir=database/migrations/schema postgres ${DSN} up

md: migrate/down
migrate/down:
	@echo "Rolling back migrations..."
	@goose -dir=database/migrations/schema postgres ${DSN} down

mf: migrate/fresh
migrate/fresh:
	@echo "Dropping..."
	@goose -dir=database/migrations/schema postgres ${DSN} reset
	@echo "Migrating..."
	@$(MAKE) --no-print-directory migrate/up
