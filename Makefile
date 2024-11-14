.DEFAULT_GOAL := run
help:
	@echo "Usage:"
	@echo "  make         - Run the Go application"
	@echo "  make help    - Show this help message"
	@echo "  make run-Database - Run the Go Database Server"
	@echo "  make run-Web - Run the Go Web Application Server"
	@echo "  make run-Cli - Run the Go CLI Application"
run:
	@echo "Running the Go application"
	@go run . $(ARGS)
