AUTH_BINARY=authApp
TEAMS_BINARY=teamsApp

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_team build_auth
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

## build_team: builds the broker binary as a linux executable
build_team:
	@echo "Building team binary..."
	cd ../teams-service && env GOOS=linux CGO_ENABLED=0 go build -o ${TEAMS_BINARY} .
	@echo "Done!"

## build_auth: builds the auth binary as a linux executable
build_auth:
	@echo "Building auth binary..."
	cd ../authentication-service && env GOOS=linux CGO_ENABLED=0 go build -o ${AUTH_BINARY} ./cmd/api
	@echo "Done!"

## drop_db: drops the users-authentication database
drop_auth_db:
	docker exec project-postgres-1 dropdb -U postgres users

createdb:
	docker exec -it project-postgres-1 createdb --username=postgres --owner=postgres users

migrateup:
	migrate -path ./../authentication-service/db/migration -database "postgres://postgres:password@localhost:5433/users?sslmode=disable" -verbose up

migratedown:
	migrate -path ./../authentication-service/db/migration -database "postgres://postgres:password@localhost:5433/users?sslmode=disable" -verbose down

## drop_db: drops the team database
drop_teams_db:
	docker exec project-teams-postgres-1 dropdb -U postgres team

create_teams_db:
	docker exec -it project-teams-postgres-1 createdb --username=postgres --owner=postgres team

teams_migrateup:
	migrate -path ./../teams-service/db/migration -database "postgres://postgres:password@localhost:5434/team?sslmode=disable" -verbose up

teams_migratedown:
	migrate -path ./../teams-service/db/migration -database "postgres://postgres:password@localhost:5434/team?sslmode=disable" -verbose down

# mock_teams:
# 	mockgen -destination=db/mocks/db_client.go teams/db DBClientInterface

# auth_server:
# 	go run ./../authentication-service/cmd/api/main.go

.PHONY: drop_auth_db createdb migrateup migratedown drop_teams_db create_teams_db teams_migrateup teams_migratedown

# in teams-microservices/authentication-service/cmd/api go run main.go  is running, but in teams-microservices/project go run ./../authentication-service/cmd/api/main.go is giving errors