.DEFAULT_GOAL := run-Database
help:
	@echo "Usage:"
	@echo "  make         - Run the Go application"
	@echo "  make help    - Show this help message"
	@echo "  make run-Database - Run the Go Database Server"
	@echo "  make run-Web - Run the Go Web Application Server"
	@echo "  make run-Cli - Run the Go CLI Application"
run-Database:
	@echo "Running the Go Database Server"
	cd Database && go run -mod=vendor main.go && cd ..
run-Web:
	@echo "Running the Go Web Application Server"
	cd Web && go run -mod=vendor main.go && cd ..
run-Cli:
	@echo "Running the Go CLI Application"
	cd Cli && go run -mod=vendor main.go && cd ..
