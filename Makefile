.DEFAULT_GOAL := run
help:
	@echo "Usage:"
	@echo "  make         - Run the Go application"
	@echo "  make help    - Show this help message"
run:
	go run main.go
