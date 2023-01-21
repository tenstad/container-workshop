# Build and Publish Container Images

## Task 1

Modify `Dockerfile` so that the `html` folder exists in the correct place within the container, as specified in the `nginx.conf`.

## Task 2

Complete the Dockerfile for the golang webserver located in the `go` folder.
Copy the `src` file into the container and run `go build` within the folder to
build a `app` executable. Finally specify that the `app` should be run when the
container starts.

Build an image from the Dockerfile and name it `goweb`.

## Task 3

Publish an image to DockerHub.
