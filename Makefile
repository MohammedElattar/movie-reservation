run: build
	@./bin/api

build: 
	@go build -o ./bin/ ./cmd/api/
