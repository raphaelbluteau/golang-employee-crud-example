.PHONY: migrate run down

migrate:
	docker-compose run --rm flyway /wait-for-postgres.sh postgres flyway migrate

run:
	docker-compose up -d postgres
	make migrate
	export $(grep -v '^#' .env | xargs) && go run api/main.go

down:
	docker-compose down
