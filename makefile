# run test coverage
test:
	go test -v -cover ./...
# run linter run it to avoid waste actions minuts
lint:
	golangci-lint run
# start the application in docker
start_docker:
	docker-compose up
# develop start
start:
	go run cmd/main.go
# start local dev for develop
init_db:
	docker run --name leal -e POSTGRES_DB=db-leal -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=root -p 5434:5432 -d postgres