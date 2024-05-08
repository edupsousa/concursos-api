build:
	@go build -o bin/concursos-api cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/concursos-api

start-db:
	@docker run --detach --name concursos-api-db --env MARIADB_ROOT_PASSWORD=root --env MARIADB_DATABASE=concursos-api --rm -p 127.0.0.1:3306:3306/tcp mariadb:latest

stop-db:
	@docker stop concursos-api-db

connect-db:
	@mariadb -h 127.0.0.1 -u root -proot concursos-api