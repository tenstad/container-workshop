# Docker Compose

## What is Docker Compose?
Docker Compose is a tool for defining and running multi-container Docker applications. With compose you can run multiple cooperating contianers with only a single command that builds your containers and sorrunding environment based on the configurations described in a `docker-compose.yaml` file.

TODO: elaborate

## YAML-configuration

### `services`
A service is equivalent to a container or an application. 

- Service properties:

```yaml
image: <image_name>
name: <service_name>
volumes:
  - <volume_this_service_uses>
  - <another_volume_this_service_uses>
networks: 
  - <network_this_service_should_use>
  - <another_network_this_service_should_use>
ports:
  - <port_mapping_0>
  - <port_mapping_1>
  - <port_mapping_2>
environment:
  - <environment_variable_name>: <value_of_the_env_variable>
  - <another_environment_variable_name>: <value_of_the_env_variable>
```

### `volumes`
TODO: elaborate


## Complete Example
TODO: add example yml

## Usage??
```bash
$ docker compose up [OPTIONS] [SERVICE...]
```
The most basic command to create and start containers in a compose-file is only by adding the `-d` flag which starts the containers in a detached mode running in the background. 

Run the below command in a directory where your docker-compose.yml file resides:
```bash
$ docker compose up -d --build
```
`-- build` is for building or rebuilding the services listed in the compose-file.

When stopping all the containers run the following command:
```bash
$ docker compose down
```

It is also possible to stop only specific containers:
```bash
$ docker stop [container-name or contianer ID]
```
_Tip: you only need to specify the number of characters or digits (from the beginning) in the container ID that makes it unique among your running continaers._

## Tasks
In this session we will build a realworld application put together by 
- a frontend written in React Redux
- a backend written in 
- a database

### Task 1
Change directory to `06-docker-compose/src`

```bash
cd 06-docker-compose/src
```

Take a look at det docker-compose.yml file and run the following command:

```bash
docker compose up -d --build
```

List the running containers with the `docker ps` command and see that a frontend application is available at [localhost:4100](localhost:4100). 

### Task 2
This frontend application is using the production backend api. This can be overridden with a environment variable

```yaml
environment:
      - REACT_APP_BACKEND_URL=conduit-backend:3000
```

add this to the spec of the `conduit-frontend` service in the `docker-compose.yml` file and run the same docker-compose command as in Task 1.

### Task 3
When opening the browser at [localhost:4100](localhost:4100) now we see that the appliaction will not load because it does not get any response form the locally running backend api - lets create one. 

We already have a backend code in `06-docker-compose/src/backend/node-realworld`. Follow the same structure as the one describing the frontend-service and include these instructions:

```bash
container_name:
build:
ports:
```

### Task 4
The backend project includes a `.env` file which defines some environment variables with values used by the mysql database which is included in the project under `06-docker-compose/src/backend/node-realworld/packages/server/sql/`. 

This database can also run as a service defined through the docker-compose.yml file. 

TODO: Figure out if .env is overridable??

Add another service to your `docker-compose.yml` file:

```yaml
real_world:
    image: mysql:latest
    ports: 
      - "3306:3306"
    volumes:
      - realworld:/var/lib/mysql
```

Find the environment variables in `06-docker-compose/src/backend/node-realworld/packages/server/.env` and include these as environment variables in the backend-application by using the `environment:` description.