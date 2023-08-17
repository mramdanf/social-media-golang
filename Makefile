USER_BINARY=userApp

down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

up_build: build_user
	@echo "Stopping docker iamges (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker iamges..."
	docker-compose up --build -d
	@echo "Docker iamges built and started!"

build_user:
	@echo "Building user binary..."
	cd user-service && env GOOS=linux CGO_ENABLED=0 go build -o ${USER_BINARY} ./cmd/restserver
	@echo "Done!"