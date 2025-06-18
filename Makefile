.PHONY: build run test clean gen-swagger migrate-up migrate-down start-infra stop-infra

# Build the application
build:
	go build -o bin/app main.go

# Run the application
run:
	go run main.go

# Run tests
test:
	go test -v ./...

# Clean build files
clean:
	rm -rf bin/
	go clean

# Generate Swagger documentation
gen-swagger:
	swag init -g main.go -o docs

# Run database migrations up
migrate-up:
	migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/go_template?sslmode=disable" up

# Run database migrations down
migrate-down:
	migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/go_template?sslmode=disable" down

# Start infrastructure services
start-infra:
	docker-compose up -d

# Stop infrastructure services
stop-infra:
	docker-compose down

# Install dependencies
deps:
	go mod download

# Install development tools
install-tools:
	go install github.com/swaggo/swag/cmd/swag@latest
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Help
help:
	@echo "Available commands:"
	@echo "  make build         - Build the application"
	@echo "  make run          - Run the application"
	@echo "  make test         - Run tests"
	@echo "  make clean        - Clean build files"
	@echo "  make gen-swagger  - Generate Swagger documentation"
	@echo "  make migrate-up   - Run database migrations up"
	@echo "  make migrate-down - Run database migrations down"
	@echo "  make start-infra  - Start infrastructure services"
	@echo "  make stop-infra   - Stop infrastructure services"
	@echo "  make deps         - Install dependencies"
	@echo "  make install-tools - Install development tools"

