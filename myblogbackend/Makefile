.PHONY: build run test clean

# Build the application
build:
	go build -o bin/myblogbackend ./cmd/myblogbackend

# Run the application
run:
	go run ./cmd/myblogbackend

# Run tests
test:
	go test -v ./...

# Clean build files
clean:
	rm -rf bin/
