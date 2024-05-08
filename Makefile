build:
	@go build -o bin/concursos-api cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/concursos-api
