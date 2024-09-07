

APP_NAME := frank

# Default target
.DEFAULT_GOAL := run

# Build the Go application
build:
	@echo "Building $(APP_NAME)..."
	rm -rf .build
	mkdir .build
	go build -o ./.build/$(APP_NAME) .

# Run the application with arguments
run: build
	./.build/$(APP_NAME) $(filter-out $@,$(MAKECMDGOALS))

# Catch-all target to allow arguments to be passed
%:
	@:

.PHONY: build run
