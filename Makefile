.DEFAULT_GOAL := start
help:
	@echo "Usage:"
	@echo "  make         - Run the Go application"
	@echo "  make help    - Show this help message"
	@echo "  make run     - Run the Go application"
	@echo "  make sql     - Generate sql functions"
start:
	@echo "Running docker-compose"
	@docker-compose -f config/compose.yml build
	@docker-compose -f config/compose.yml up -d
stop:
	@echo "Stopping docker-compose"
	@docker-compose -f config/compose.yml down
run:
	@echo "Running the Go application"
	@go run .
web:
	@echo "Running the Go application"
	@go run . -web
build:
	@echo "Building the Go application"
	@go build -o bin/app .
sql:
	@echo "Generating sql queries"
	@sqlc generate -f config/sqlc.yml

.PHONY: help run sql start stop web build
