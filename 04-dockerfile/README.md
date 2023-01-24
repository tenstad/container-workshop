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

`COPY` copies files from your local filesystem into the container image. Only
files within the _build context_ (a folder specified on build) may be copied. It
is quite similar to `cp` on unix systems. Both single files and directories may
be copied, and the target path will automatically be created if absent. If the
destination is relative (not staring with `/`), it will be relative to the
`WORKDIR`.

```Dockerfile
# Copy all files within the build context into the working directory of the
# container image
COPY . .

# Copy only the src folder within the build context into src within the root
# of the container filesystem
COPY src /src

# Copy the file index.html into /var/www/html/index.html, even though the
# folder /var/www/html/ does not already exist
COPY html/index.html /var/www/html/index.html
```

### `RUN`

`RUN` runs a command within the shell of the container beeing built. It is often
used to update or install packages, and compile or build the project. Commands
are often written on one line each, with `\` at the end of all lines but the
last.

```Dockerfile
RUN npm install \
    && npm build
```

### `ARG`

`ARG` is used to accept build arguments passed when building an image. It may be
the version of a piece of software or any other value that you wish to specify
when building.

```Dockerfile
ARG JQ_VERSION=1.6
RUN apt-get update && apt-get install -y \
    curl \
    && rm -rf /var/lib/apt/lists/* \
    && curl -sSLo /usr/local/bin/jq \
    https://github.com/stedolan/jq/releases/download/jq-${JQ_VERSION}/jq-linux64 \
    && chmod +x /usr/local/bin/jq
```

### `ENV`

`ENV` is used to define environment variables within the container. They are
variables that may be used by any process within the container, such as `PATH`
or `JAVA_HOME`. Note that you may define additional environment variables when
running a container based on the image. It may be used in conjunction with
`ARG`.

```Dockerfile
ARG LOG_LEVEL=info
ENV LOG_LEVEL=$LOG_LEVEL
```

Here, one may override the log level when building the application, but by
default it is `info`.

### `CMD`

`CMD` is used to specify the command or executable that runs on startup.

```Dockerfile
CMD [ "python3", "-m" , "flask", "run", "--host=0.0.0.0"]
```

### `ENTRYPOINT`

Similar to `CMD`. If you want, see [Docker's
description](https://docs.docker.com/engine/reference/builder/#entrypoint), as
it is a bit tricky to explain.

## Linting with Hadolint

[Hadolint](https://github.com/hadolint/hadolint) will help you follow [best
practices for writing
Dockerfiles](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/),
as mentioned in the introduction. It is strongly recommended to make it part of
your CI/CD pipeline. In addition to avoiding pointless arguments, it will make
sure the Dockerfiles are (more likely) error free and that the resulting
container image is more reliable and smaller in size.

📝 See some examples of what Hadolint may assist with over at:
<https://hadolint.github.io/hadolint>
