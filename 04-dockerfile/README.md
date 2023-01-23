# Dockerfile

If at any point something is unclear, please refer to Docker's [Dockerfile
reference](https://docs.docker.com/engine/reference/builder/) for a more thorough
explanation. Use the _Contents_ on the right side to navigate to specific
sections or Dockerfile instructions.

Developers _always_ disagree about how to best write Dockerfiles. To avoid
arguments, try to allways follow Docker's [Best practices for writing
Dockerfiles](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/).
Bonus points if you send the link to teammates that don't.

## Instructions

### `FROM`

`FROM` is used to specify an existing container image on which you build upon -
the starting point for your image. It is typically an operating system (e.g
`debian` or `alpine`), a programming language or collection of build tools (e.g.
`node`, `python`, or some sdk), or some sort of runtime or program (e.g.
`dotnet` or `nginx`). It may be wise to also specify a specific tag (e.g.
`python:3.10`) so that the base don't suddely change when some third party
pushes a newer image.

```Dockerfile
FROM python:3
```

You may rarely encounter `FROM scratch` in the wild, which indicates that the
container image is build from scratch - a totally empty starting point.

### `WORKDIR`

`WORKDIR` sets the working directory within the file-system. The folder will be
created if it does not already exist. If a relative path is provided (not
starting with `/`), the resulting workdir will be relative to one previously
specified.

```Dockerfile
WORKDIR /app
```

### `COPY`

TODO: COPY

### `RUN`

TODO: RUN

### `ARG`

TODO: ARG

### `ENV`

TODO: ENV

### `CMD`

TODO: CMD

### `ENTRYPOINT`
Very similar to `cmd` but these commands will _not_ be overwridden by arguments stated through the CLI when running a container.

## Example Dockerfile

TODO: Basic FROM COPY RUN CMD example

## Linting with Hadolint

[Hadolint](https://github.com/hadolint/hadolint) will help you follow [best
practices for writing
Dockerfiles](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/),
as mentioned in the introduction. It is strongly recommended to make it part of
your CI/CD pipeline. In addition to avoiding pointless arguments, it will make
sure the Dockerfiles are (more likely) error free and that the resulting
container image is more reliable and smaller in size.

TODO: How to Hadolint (local or https://hadolint.github.io/hadolint/)

TODO: Task: Lint your Dockerfile
