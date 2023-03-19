.PHONY: migrate run down loadtest

migrate:
	docker-compose run --rm flyway /wait-for-postgres.sh postgres flyway migrate

run:
	docker-compose up -d postgres
	make migrate
	#export $(grep -v '^#' .env | xargs) && go run api/main.go
	docker-compose up --build app

down:
	docker-compose down

loadtest:
	docker-compose up --build k6
