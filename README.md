# my-team-microservices

## Project

This is the root directory of a project that contains two microservices - an authentication service and a team service.

# Usage

The Makefile in the project directory provides several targets to build, start and stop the Docker containers for the two services, and manage their databases.

Start Docker containers
To start the Docker containers for both services, run:

make up
This will start the containers in the background without forcing a build. If the containers have already been built, this target will simply start them.

If you want to force a rebuild of the containers, you can run:

make up_build
This will stop any running containers, rebuild the services' binaries, and then start the containers again.

# Stop Docker containers
To stop the Docker containers for both services, run:

make down
This will stop the running containers.

# Manage databases
The Makefile also provides targets to manage the databases for the services.

To drop the authentication service's database, run:

make drop_auth_db
To create the authentication service's database, run:

make createdb
To run migrations on the authentication service's database, run:

make migrateup
To roll back migrations on the authentication service's database, run:

make migratedown
Similarly, to drop, create, migrate up or migrate down the team service's database, you can use the following targets:

make drop_teams_db
make create_teams_db
make teams_migrateup
make teams_migratedown
Note that the authentication service's database runs on port 5433 and the team service's database runs on port 5434.

# Other targets

The Makefile also includes targets to build the binaries for the authentication and team services. These targets are used internally by the up_build target, but can also be run independently.

To build the binary for the authentication service, run:

make build_auth
To build the binary for the team service, run:

make build_team
These targets create a Linux executable for each service in their respective directories, and name them authApp and teamsApp.