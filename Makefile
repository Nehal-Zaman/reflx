build:
	@go build -o bin/reflx -ldflags="-s -w"

run:build
	@./bin/reflx