include .env
## build: Build binary
build:
	@echo "Building..."
	env CGO_ENABLED=0  go build -ldflags="-s -w" -o ${BINARY_NAME} ./cmd/file-sharing
	@echo "Built!"

## run: builds and runs the application
run: build
	@echo "Starting..."
	@env ./${BINARY_NAME} &

## debug: runs application in debug mode
debug:
	@echo "Application started on debugging mode..."
	@go run . -debug=true

## clean: runs go clean and deletes binaries
clean:
	@echo "Cleaning..."
	@go clean
	@rm ${BINARY_NAME}
	@echo "Cleaned!"

## start: an alias to run
start: run

## stop: stops the running application
stop:
	@echo "Stopping..."
	@-pkill -SIGTERM -f "./${BINARY_NAME}"
	@echo "Stopped!"

## restart: stops and starts the application
restart: stop start

## run test: runs all tests
test:
	go test -v ./...
