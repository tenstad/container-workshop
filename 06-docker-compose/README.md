# Docker Compose

Docker Compose is a tool used to configure and run applications consisting of
multiple containers. With compose you can run multiple cooperating contianers
with only a single command, that builds your containers and sorrunding
environment based on the configurations described in a `docker-compose.yaml`
file.

It is in essence two parts to it. The first beeing the docker compose file's
standard for defining the commandline flags of `docker run`, for multiple
containers. The second is the tool that runs the containers with the
configuration defiend in the file.

## YAML-configuration

A docker-compose file often consist of multuple services. Think of a service as an application or a container. In the example below you can see some of the fiels one may configure:

```yaml
services:
  backend:
    build: ./backend
    restart: always
    ports:
      - 8000:8000
    environment:
      - DATABASE_HOST=database:5432
    network:
      - web
    volumes: ./backend/cache:/app/cache/
```

The container name is `backend`, the maps item's key within `services`. It has
no `image` defined, but is rather build from the dockerfile within the
`./backend` directory. It is configured to restart if were to stop. The other
keys should be known from previous sections. In short, the port `8000` is
forwarded to the container, and it is part of the network `web`. Additionally, a
volume is mounted so that the cache is kept even though the container is
restarted.

## Usage

Depending on how it is installed, use either `docker-compose` or `docker compose` to run the cli tool. The most basic commands are:

```bash
docker compose up -d --build
docker compose down
```

`up` brings up all the services configured in the dockerfile and outputs the
logs for all of them. As with `docker run`, `-d` may be specified to run it
detached (in the background). If using `build`, it is also wise to use `--build`
to force the images to be rebuilt before the container starts. Otherwise,
changes to source files may not be reflected in the new container.

`down` stops all the containers, just like `docker stop`.

## Trial run

📝 Run the following to start the containers defined in `docker-compose.yml`:

```bash
docker compose up -d
```

The containers are now running as if you started them manually with `docker run`.

📝 See whats happening with `docker ps` and `docker compose logs`

📝 To stop all the containers run the following:

```bash
docker compose down
```

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
