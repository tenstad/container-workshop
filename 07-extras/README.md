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

## CNCF

https://landscape.cncf.io/

## Open Container Initiative

https://opencontainers.org/
