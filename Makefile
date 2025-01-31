build:
	@go build -o bin/Dbgo cmd/main.go

run: build
	@./bin/Dbgo
test:
	@go test -v ./...