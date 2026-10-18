include .env
export

service-run:
	@go run ./cmd/learning-hub

service-deploy:
	docker compose up -d --build application 

service-undeploy:
	docker compose down application

service-logs:
	docker compose logs application

docker-create:
	docker run -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) -p 5432:5432 -v ./out/pgdata:/var/lib/postgresql -d postgres:18-bookworm

db-start:
	docker start learning-hub-db

db-stop:
	docker stop learning-hub-db

migrate-up:
	migrate -path migrations -database ${DB_URL} up

migrate-down:
	migrate -path migrations -database ${DB_URL} down