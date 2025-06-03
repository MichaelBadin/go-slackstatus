# Binary name
BIN := slackstatus

# Default target
.PHONY: all
all: build

# Build the binary
.PHONY: build
build:
	go build -o $(BIN) main.go

# Run the CLI with optional ARGS
.PHONY: run
run: build
	./$(BIN) $(ARGS)

# Clean the build output
.PHONY: clean
clean:
	rm -f $(BIN)

# Run tests
.PHONY: test
test:
	go test ./...

# Lint using go vet (or customize to golangci-lint)
.PHONY: lint
lint:
	go vet ./...
