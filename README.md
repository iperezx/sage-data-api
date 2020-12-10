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
There are only two operations supported:
- Get all the nodes

Get all the nodes:
```bash
curl GET 'http://localhost:8080/api/v1/nodes'
```

Returns:
```bash
[
    {
        "id": "4cd98fc4d2a8",
        "name": "Sage-NEON-01"
    }
]
```