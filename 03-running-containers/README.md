# Running Containers

The most basic command is `docker run <image>`, which instantiates a running
container from a given image. If the image is not already present on your
machine, it is automatically pulled.

📝 Run the following:

```bash
docker run hello-world
```

You should get an output starting with _Hello from Docker!_ indicating that
everything is setup correctly!

What just happened? First, the `hello-world` container image was not found on
your local machiene, and was therefore downloaded from DockerHub. A container
based on the image was then started. It ran for a very brief time before its
main process exited succesfully. As it completed, the container then terminated.
You can see the terminated container with the command `docker ps -a`. The `-a`
stands for _all_, not only containers that are currently running.

📝 List all your containers (including the stopped), then remove the
`hello-world` container:

```bash
docker rm <id of container>
```

NOTE: you do not need to write the entire ID, just enough characters to make ite
ID unique.

## Flags

The docker run command has a lot of flags that configure who the container is
run. These are placed inbetween `docker run` and the the name of the image.

E.g. `docker run --rm -it -p 8080:80 -v $PWD/05-*/nginx/html:/usr/share/nginx/html nginx`

### `--rm`

`--rm` may be specified to automatically delete a container once it is stopped.
Comes in handy when you start a lot of containers you know you will no longer
need once stopped.

📝 Run `hello-world` with `--rm` and verify that the container is automatically
deleted afterwards.

### `-d`

`-d` is short for `--detach`. For long-running containers such as a web server,
you might want to run the container in the background instead of it occupying a
terminal.

📝 Run the webserver image `nginx` in detached mode.

📝 List the running containers.

📝 Show the logs of the nginx contaner with the `docker logs` command.

📝 Delete the container. You will need to use the force (`--force`) if not
gracefully `docker stop`ping it first.

### `-it`

`-it` is a combination of both `--interactive` and `--tty`. `-t` essentially
creates a virtual terminal session and `-i` forwards what you write to the
container. These are most often used together, in the case when you want to
interact with the container.

📝 Run for example:

```bash
docker run --rm -it ubuntu
```

You are now in a terminal of an ubuntu distribution of Linux and may play around
without causing harm to you actual system. Such an image may be extended with a
custom Dockerfile, to create a playground with a lot of tooling installed - that
may be totally reset at will.

Type `exit` to quit and destroy (because of `--rm`) the container btw.

### `-v` and `--workdir`

`-v` is short for `--volume`. In the case with ubuntu above, you may want to
make files from your own system accessible inside the container. `-v` lets you
_mount_ an (absolute) path from your host into a path of the container. By
additionally specifying `--workdir`, that path may be chosen as the starting
directory.

📝 Run the following:

```bash
docker run --rm -it -v $HOME:/hostuser --workdir /hostuser ubuntu
```

If you now run `ls` inside the container, you should see the content of your
computers home directory. Use with care though, you can delete stuff aswell.

Sidenote: `docker volume` may be used to create additional volumes that are
intially empty and may be mounted into one or more containers.

### `-p`

`-p` is short for `--publish` (aka port). It is used to forward a port on the
local machine to another port of the container. The syntax is `-p <local
port>:<container port>`.

📝 Run the nginx container and port-forward any local port to `80` inside the
container. Then open `localhost:` followed by the chosen port in a browser and
see if you are able to communicate with the container.

📝 Inspect the container logs. If you ran the container with `-it` you will see
them instantly. If in detached mode (`-d`), use `docker ps` and `docker logs`.

### `-e`

`-e` is short for `--env`. It is often used to pass the container secrets that
one does not want to build into the image (e.g `--env KEY=123`). Environent
values are not yet determined at buildtime or the values vary between
deployments. It may be the url to the api, that differ in dev and production, or
the logging format to use.

### `--network` and `--name`

`docker network` is used to create and manage networks. Containers may be
entered into a network when started, and can then communicate with other
containers by names rather than IP addresses (e.g. the url to backend is
`http://backend:3000` because it was started with `--name backend`).

📝 Create a network. Use `docker network --help` to guide you.

📝 Start a nginx container in the network container you just created. Remember
to give it a name.

📝 Start an ubuntu container in interactive mode on the same network and connect
to the nginx one with `curl <container name>`.

NOTE: Install curl with `apt update && apt install -y curl`.

Sidenote: `--network host` is a _special_ network that will make the containers
avaible from the host machine without port-forwarding. You are not able to
choose ports tho (as with `8080:80`), meaning multiple containers may fight over
a given port.

📝 If you wish: run `docker run --rm -it --network host nginx` in two terminals
and see what happes. Additionally, open `localhost` in a browser.

## Entering a running container

📝 With an nginx container running and a port forwarded, do the following:

```bash
docker exec <container name/id> ls
docker exec -it <container name/id> bash
```

`docker exec` enables you to execute another process within the container. In
the first case, a simple `ls` (listing files and directories) is executed and
the output returned. In the second one an interactive shell process is started.

📝 You may now modify the webserver while it runs instide the container, by
installing a text exitor and updating the html files:

```bash
apt update && apt install nano
nano /usr/share/nginx/html/index.html
nginx -s reload
```

`nginx -s reload` restarts the webserver process. If a port is exposed, you
should be able to see the changes in your browser.
