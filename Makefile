# Define your Go executable name
GOEXEC = createYourPassword

# Define the Go build command
GOBUILD = go build -v

# Install dependencies (Go modules)
install-deps:
	@echo "Installing dependencies..."
	@go mod tidy

# Build the Go program
build: install-deps
	@echo "Building the program..."
	@$(GOBUILD) -o $(GOEXEC)

# Run the program
run: build
	@echo "Running the program..."
	@./$(GOEXEC)

# Clean the build
clean:
	@echo "Cleaning up..."
	@rm -f $(GOEXEC)

# Default target
all: build