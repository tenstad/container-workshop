# Container Workshop

TODO: Write an introduction - essence of containers

Containers have been popularized by _Docker, Inc._ and their tooling. The
concepts are thus often refered to as _docker_ things, e.g. _docker files_,
_docker images_, _docker containers_. However, container technology is much more
than just Docker.

This is a hands-on workshop where you will learn about container technology
though actual use, along with some of the tooling that improves upon the
experience of developing and working with containers. Any container technology
tools may be used to build images and run containers, but Docker has been chosen
in the examples as it likely the most accessible for the majority of people.

## Prerequisites

- Install a _container image build tool_ and _container runtime_, the simplest
  beeing the all-in-one solution [Docker
  Desktop](https://www.docker.com/products/docker-desktop/). See [Docker
  FAQs](https://www.docker.com/pricing/faq/?utm_campaign=2022-08-31-desktop-update)
  for any questions regarding Docker Desktop. If using any other container image
  build tool or runtime, be sure to adapt the workshop's `docker` commands to
  fit your tools.
- [Git](https://git-scm.com/) - as you probably want to `git clone https://github.com/tenstad/container-workshop.git`
- Create a user on [DockerHub](https://hub.docker.com/) or any other container
  image registry.

## Content

- [Docker Desktop](./01-docker-desktop)
- [Container Images](./02-container-images)
- [Running Containers](./03-running-containers)
- [Dockerfile](./04-dockerfile)
- [Build and Publish Container Images](./05-build-and-publish)
- [Docker Compose](./06-docker-compose)
- [Vulnerability Scanning](./07-vulnerability-scanning)
- [Extras](./08-extras)
