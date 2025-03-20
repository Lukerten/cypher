.RECIPEPREFIX := >
MAKEFLAGS += --silent
MAKEFLAGS += --no-print-directory

# Variables
BINARY_NAME = cypher
SRC_DIR = ./src
BIN_DIR = ./bin

# Default target
all: build

#HELP: Build the Go application
build:
> mkdir -p $(BIN_DIR)
> go build -o $(BIN_DIR)/$(BINARY_NAME) $(SRC_DIR)/...

#HELP: Run the Go application
run: build
> ./$(BIN_DIR)/$(BINARY_NAME) $(ARGS)

#HELP: Clean up build artifacts
clean:
> rm -f $(BIN_DIR)/$(BINARY_NAME)

#HELP: Clean up all build artifacts and directories
dist-clean: clean
> rm -rf $(BIN_DIR)

#HELP: prints this screen
help:
> @printf "Available targets\n\n"
> @awk '/^[a-zA-Z\-_0-9]+:/ { \
>   helpMessage = match(lastLine, /^#HELP: (.*)/); \
>   if (helpMessage) { \
>     helpCommand = substr($$1, 0, index($$1, ":")-1); \
>     helpMessage = substr(lastLine, RSTART + 6, RLENGTH); \
>     gsub(/\\n/, "\n", helpMessage); \
>     printf "\033[36m%-30s\033[0m %s\n", helpCommand, helpMessage; \
>   } \
> } \
> { lastLine = $$0 }' $(MAKEFILE_LIST)

