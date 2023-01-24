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

A docker-compose file often consist of multuple services. Think of a service as
an application or a container. In the example below you can see some of the
fiels one may configure:

```yaml
services:
  backend:
    container_name: backend
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

The container name is `backend`. It has no `image` defined, but is rather build
from the dockerfile within the `./backend` directory. It is configured to
restart if were to stop. The other keys should be known from previous sections.
In short, the port `8000` is forwarded to the container, and it is part of the
network `web`. Additionally, a volume is mounted so that the cache is kept even
though the container is restarted.

## Usage

Depending on how it is installed, use either `docker-compose` or `docker
compose` to run the cli tool. The most basic commands are:

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

📝 See whats going on with `docker ps` and `docker compose logs`.

One container name will be automatically deduced from the folder and container
id, as `container_name` was not specified.

📝 To stop all the containers run the following:

```bash
docker compose down
```

## Tasks

In this session we will build a realworld application consisting of the
following:

- Frontend written in React Redux
- Backend written in Node
- Mysql database

### Prerequisites

In the `06-docker-compose/realworld` folder, clone the following react frontend
[khaledosman/react-redux-realworld-example-app](https://github.com/khaledosman/react-redux-realworld-example-app)
and node
backend[ditsmod/realworld](https://github.com/ditsmod/realworld#getting-started):

```bash
git clone git@github.com:khaledosman/react-redux-realworld-example-app.git realworld-frontend

git clone git@github.com:ditsmod/realworld.git realworld-backend
```

📝 Both the frontend and backend are missing Dockerfiles, so you will have to
create these yourself, one inside the frontend folder and another in the
backend.

NOTE: Some dependency issues might occur in the frontend, so try using `--force`
when installing dependencies with `npm`.

### Task 1

Change directory to `06-docker-compose/realworld`:

```bash
cd 06-docker-compose/realworld
```

Take a look at the `docker-compose.yml` file and run the following command:

```bash
docker compose up -d --build
```

List the running containers with the `docker ps` command and see that a frontend
application is running with the port `4100` exposed.

### Task 2

When opening the browser at [localhost:4100](http://localhost:4100) we see that the
frontend application is running and it is using the production backend api. Lets
run the backend in a container instead.

We have already cloned the backend project in
`06-docker-compose/realworld/backend/node-realworld`. Follow the same structure
as the one describing the frontend-service and include these instructions:

```bash
container_name:
image:
build:
ports:
environment:
  - NODE_OPTIONS="--max-old-space-size=4096"
```

Run the following command again:

```bash
docker compose up -d --build
```

See that all containers are running with the `docker ps` command.

### Task 3

Now try to access the appliaction through [localhost:4100](localhost:4100)
again.

Woops. No connection?? Try adding an environment variable that defines the
backend-service with its port as the value.

```yaml
environment:
  - REACT_APP_BACKEND_URL=[backend-service:port]
```

Add this to the spec of the `conduit-frontend` service in the
`docker-compose.yml` file and run the same docker-compose command as above.

### Task 4

Now [localhost:4100](localhost:4100) should show the same interface as the first
time you built the docker-compose file. But something is missing...

There are noe tags and no posts showing. We need to add a database for this.

Add another service to your `docker-compose.yml` file:

```yaml
real_world:
  container_name: database
  image: mysql:latest
  ports:
    - "3306:3306"
  volumes:
    - ./realworld-backend/packages/server/sql/dump:/docker-entrypoint-initdb.d/
  environment:
    - MYSQL_ALLOW_EMPTY_PASSWORD=true
    - MYSQL_DATABASE=realworld
```
