# Docker Compose

## What is Docker Compose?
Docker Compose is a tool for defining and running multi-container Docker applications. With compose you can run multiple cooperating contianers with only a single command that builds your containers and sorrunding environment based on the configurations described in a `docker-compose.yaml` file.

TODO: elaborate

## YAML-configuration

### `services`
A service is equivalent to a container or an application and is referred to by its image-name. 

- Service properties: (Trenger kanskje ikke liste alle?)

```yaml
image: <image_name>
name: <service_name>
volumes:
  - <volume_this_service_uses>
  - <another_volume_this_service_uses>
networks: 
  - <network_this_service_should_use>
  - <another_network_this_service_should_use>
ports:
  - <port_mapping_0>
  - <port_mapping_1>
  - <port_mapping_2>
profiles: (only in compose v2)
  - <profile_this_service_apply_to>
  - <another_profile_this_service_apply_to> 
secrets: 
  - <secret_this_service_uses>
  - <another_secret_this_service_uses>
environment:
  - <environment_variable_name>: <value_of_the_env_variable>
  - <another_environment_variable_name>: <value_of_the_env_variable>
```

### `volumes`
TODO: elaborate

### `networks`
TODO: elaborate

### `secrets`
TODO: elaborate

## Complete Example
TODO: add example yml

## Usage??
```bash
$ docker compose up [OPTIONS] [SERVICE...]
```
The most basic command to create and start containers in a compose-file is only by adding the `-d` flag which starts the containers in a detached mode running in the background. 

Run the below command in a directory where your docker-compose.yml file resides:
```bash
$ docker compose up -d
```

When stopping all the containers run the following command:
```bash
$ docker compose down
```

It is also possible to stop only specific containers:
```bash
$ docker stop [container-name or contianer ID]
```
_Tip: you only need to specify the number of characters or digits (from the beginning) in the container ID that makes it unique among your running continaers._
