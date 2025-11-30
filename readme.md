# Golang Boilerplate API with Clean Architecture

This is a boilerplate for building a RESTful API with Golang using Clean Architecture. It includes authentication, authorization, and other best practices.

## Features

- Clean Architecture
- Gin Framework
- GORM for database
- PostgreSQL
- Authentication and Authorization
- Validation
- Logging
- Testing
- Swagger documentation
- Air for live reload
- Docker support

## Getting Started

To run this project, you need to have Go and Docker installed on your machine.

1. Clone this repository
2. Copy `.env.example` to `.env` and fill in the required variables

3. Run `docker-compose up -d` to start the database
4. Run `go run main.go` to start the server
5. Run `go test ./...` to run the tests
6. Run `swag init` to generate the swagger documentation
7. Run `air` to start the server with live reload
8. Run `make build` to build the project
9. Run `make run` to run the project
10. Run `make test` to run the tests
11. Run `make lint` to run the linter
12. Run `make fmt` to format the code
13. Run `make migrate` to run the migrations
14. Run `make seed` to seed the database
15. Run `make docs` to generate the swagger documentation
16. Run `make docker-build` to build the docker image
17. Run `make docker-run` to run the docker container
18. Run `make docker-stop` to stop the docker container
19. Run `make docker-rm` to remove the docker container
20. Run `make docker-rmi` to remove the docker image
21. Run `make docker-prune` to prune the docker system
