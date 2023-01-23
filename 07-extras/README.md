# Extras

## Inspecting image layers

You can use `docker image history <image>` to list the instructions of a
container image, even if you do not possess the Dockerfile used to build it. It
can come in handy when debugging, but does not show the file-changes in each
layer.

[Dive](https://github.com/wagoodman/dive) is a tool that may be used to _dive_
deeper into the image and the changes of each layer. Install it and inspect the
layers in the `goweb` image creted earlier (or another image of choice).

Use the UP and DOWN arrow-key to step through each layer within the image. Press
TAB to switch focus to the file tree. There you may press CTRL+U to toggle
unmodified files. Make sure only changed files are shown. Press TAB to go back
to the layers and see what files each layer adds. Is anything redundant in the
`go build` layer?

## Distroless

Checkout ["Distroless" Container
Images](https://github.com/GoogleContainerTools/distroless#distroless-container-images)
and follow [Examples with
Docker](https://github.com/GoogleContainerTools/distroless#examples-with-docker)
to minimize the `goweb` image built earlier. How large was the original image
and how small are you able to get it?

## Buildah

Buildah is another container image building tool that does things a bit differently.

Check out the tutorials: <https://github.com/containers/buildah/tree/main/docs/tutorials>

## Podman

If you like the docker experience but not Docker Inc, maybe checkout Podman:
<https://podman.io/>. It has the same commands, but is not as locked into Docker
Inc. and their echosystem.

## Open Container Initiative

OCI defines the spec for container images: <https://opencontainers.org/>

## Cloud Native Computing Foundation

CNCF has a lot of cool open source projects revolving around containers:
<https://www.cncf.io/projects/>. Also checkout their full landscape:
<https://landscape.cncf.io/>.
