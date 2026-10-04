include .env
export

service-run:
	@go run .

service-deploy:
	docker compose up -d --build application 

service-undeploy:
	docker compose down application

service-logs:
	docker compose logs application

migrate-up:
	migrate -path migrations -database ${DB_URL} up

migrate-down:
	migrate -path migrations -database ${DB_URL} down