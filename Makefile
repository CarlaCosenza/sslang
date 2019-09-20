test:
	go test ./...
build:
	@echo "Building compiler ..."
	@go build -o ssl main.go
	@echo "Done!"
