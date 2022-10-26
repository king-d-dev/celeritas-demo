BINARY_NAME=celeritasApp

build:
	@go mod vendor
	@echo "Building Celeritas ..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "Celeritas built!"

run: build
	@echo "Starting Celeritas..."
	./tmp/${BINARY_NAME} &
	@echo "Celeritas started!"
	

clean: 
	@echo "Cleaning ..."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned!"

start: run

stop: 
	@echo "Stopping Celeritas ..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped Celeritas!"