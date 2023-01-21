# Running Containers

The most basic command is `docker run <image>`, which instantiates a running
container from a given image. If the image is not already present on your
machine, it is automatically pulled first.

Try running the following:

```bash
docker run hello-world
```

You should get an output starting with _Hello from Docker!_ indicating that
everything is setup correctly!

TODO: what actually happened (running untill exit)

TODO `docker ps -a`

## Flags

### `-d`

`-d` is short for `--detach`

### `--name`

### `-it`

`-it` is a combination of both `--interactive` and `--tty`

### `--rm`

### `-p`

`-p` is short for `--publish` (aka port)

### `-e`

`-e` is short for `--env`

### `-v`

`-v` is short for `--volume`

### `--network`

## Entering a running container

`exec -it`

