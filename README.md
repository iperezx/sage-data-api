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
The API reads a csv file (`manifest.csv`) and outputs the content in a json format to the user.
Get all the nodes:
```bash
curl GET 'http://localhost:8080/api/v1/nodes'
```

Returns:
```bash
[
    {
        "id": "4cd98fc4d2a8",
        "name": "Sage-NEON-01",
        "status": "Up",
        "OSVersion": "dell-1.0.0.local-6da62a8",
        "serviceTag": "79BBZ23",
        "SpecialDevices": "N/A",
        "BiosVersion": "2.8.2",
        "lat": "40.01631",
        "lon": "-105.24585"
    },
    ...
]
```