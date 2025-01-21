.DEFAULT_GOAL := run
help:
	@echo "Usage:"
	@echo "  make         - Run the Go application"
	@echo "  make help    - Show this help message"
	@echo "  make run     - Run the Go application"
	@echo "  make sql     - Generate sql functions"
run:
	@echo "Running the Go application"
	@go run . $(ARGS)
build:
	@echo "Building the Go application"
	@go build -o bin/app .
sql:
	@echo "Generating sql queries"
	@sqlc generate -f config/sqlc.yml

.PHONY: help run sql
