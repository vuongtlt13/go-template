.PHONY: build run test clean proto gen-proto migrate-up migrate-down start-infra stop-infra

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

remove-old-pb:
	rm -rf pb && mkdir pb

# Generate protobuf code
gen-proto: remove-old-pb
	protoc \
	  -I=proto \
	  -I=third_party \
	  --proto_path=proto \
	  --go_out=pb \
	  --go_opt paths=source_relative \
	  --go-grpc_out=pb \
	  --go-grpc_opt paths=source_relative \
	  --connect-go_out=pb \
	  --connect-go_opt paths=source_relative \
	  --validate_out="lang=go:pb" \
      --validate_opt paths=source_relative \
	  --openapiv2_out=api \
      --openapiv2_opt logtostderr=true \
      --openapiv2_opt generate_unbound_methods=true,allow_merge=true,merge_file_name=api \
	  proto/auth/*.proto proto/health/*.proto proto/admin/*.proto proto/i18n/*.proto

# Run database migrations up
migrate-up:
	migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/go_web?sslmode=disable" up

# Run database migrations down
migrate-down:
	migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/go_web?sslmode=disable" down

# Start infrastructure services (PostgreSQL)
start-infra:
	docker-compose up -d

# Stop infrastructure services
stop-infra:
	docker-compose down

