MAIN_SRC := server.go

ifeq ($(shell echo "check_quotes"),"check_quotes")
OUTFILE := bin/server-windows-386
else
OUTFILE := bin/server-linux-386
endif

.PHONY: help clean build build-all docs run start all

clean: ## Deletes all compiled / executable files
	@find bin -type f -not -name '.gitkeep' -print0 | xargs -0 rm --
	@echo "Deleted all build files"

build: ## Compile the go files
	@echo "Building go file..."
	@go build -o $(OUTFILE) $(MAIN_SRC)

build-all: ## Compile the go files for multiple OS
	@echo "Building go files for multiple OS..."
	@GOOS=linux GOARCH=arm go build -o bin/server-linux-arm $(MAIN_SRC)
	@GOOS=linux GOARCH=arm64 go build -o bin/server-linux-arm64 $(MAIN_SRC)
	@GOOS=linux GOARCH=386 go build -o bin/server-linux-386 $(MAIN_SRC)
	@GOOS=freebsd GOARCH=386 go build -o bin/server-freebsd-386 $(MAIN_SRC)
	@GOOS=windows GOARCH=386 go build -o bin/server-windows-386 $(MAIN_SRC)
	@echo "Finished building"

run: ## Runs the server
	go run $(MAIN_SRC)

start: ## Runs the compiled server in production mode
	$(OUTFILE)

all: build start ## Build and Run the server

.DEFAULT_GOAL := help