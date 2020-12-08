# sage-nodes-api

Docker container usage
-------------
The docker image is hosted on [DockerHub](https://hub.docker.com/repository/docker/sagecontinuum/sage-nodes-api)

To build image:
```bash
docker build -t iperezx/sage-nodes-api:latest .
```

To run container:
```bash
docker run -p 8080:8080 iperezx/sage-nodes-api:latest
```

To push a new tag to this repository:
```bash
docker push iperezx/sage-nodes-api:latest
```

Kubernetes Setup
-------------
The command deploys the node API on Nautilus.

User side
-------------