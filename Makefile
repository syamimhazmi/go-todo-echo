GO := go

BINARY_NAME := todo-echo

MAIN_PACKAGE := .

build: 
	$(GO) build -o $(BINARY_NAME) $(MAIN_PACKAGE)

run: build
		./$(BINARY_NAME)

clean:
	rm -rf $(BINARY_NAME)

test:
	$(GO) test -v ./...

.PHONY: build run clean test
