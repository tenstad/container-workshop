# Container Images

## Image names

Image names often have four parts:

1. The url of the _registry_ (e.g. `docker.io`)
2. The org or user that owns the image (e.g. `docker`)
3. The name of the image (e.g. `getting-started`)
4. The tag or version of the image (e.g. `latest`)

All put together, the full name is `docker.io/docker/getting-started:latest`.

Try to pull (download) the following images:

```bash
docker pull docker.io/docker/getting-started:latest
docker pull docker.io/library/nginx:1.22
```

Then view your list of local images in the _Images_ tab of the Docker Desktop
UI, or with the following command:

```bash
docker images
```

### Defaults

You may notice that `docker.io/library/nginx` is shortened to only `nginx` in
your list of images. When using the Docker tooling, the following defaults are
used:

If you omit the registry, `docker.io` is used.  
If you omit the _org_, `library` is used (the official images).  
If you omit tag is `latest` is used (indicating the latest release / newest version).

The image name `nginx` is thus the same as `docker.io/library/nginx:latest`.
