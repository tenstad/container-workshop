# Build and Publish Container Images

## Create your own website

📝 Make a few modifications to `nginx/html/index.html`.

📝 Add a line to `nginx/Dockerfile` so that the `html` folder exists in the
correct place within the container, as specified in the `nginx.conf`.

📝 Build and run it! See your magnificent html skills in the browser!

## Real (tiny) example application

📝 Complete the Dockerfile for the golang webserver located in the `go` folder.
Copy the `src` file into the container image and run `go build` within the
folder to build a `app` executable. Finally specify that the `app` executable
should be run when the container starts.

📝 Build an image from the Dockerfile, name it `goweb`,

📝 Run the image with `--env MESSAGE=<whatever you want>`.

## Publish an image to DockerHub

Sign up for an account over at <https://hub.docker.com/>, if you don't have one
already.

📝 `docker login`

📝 Either build an image with the prefix `docker.io/<your username>/`, or tag an
existing image. Preferably your website from earlier.

```bash
docker tag <existing image> <new image name>
```

`docker tag` basically creates a new copy of the image, with a new name. It is
not actually copied though, but accessible by both names afterwards.

📝 `docker push <image>` and goto `https://hub.docker.com/u/<your username>` to
see it.

Your image is now shared with the world. Anyone (with the same processor
architecture) can pull and run it!
