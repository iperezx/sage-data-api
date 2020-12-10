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
- Get all the current nodes
- Post a new node with specific metadata

Get all the nodes:
```bash
curl GET 'http://localhost:8080/api/v1/nodes'
```

Returns something like the following (without adding any data):
```bash
[
    {
        "id": "4cd98fc4d2a8",
        "name": "Sage-NEON-01"
    }
]
```

After adding more data the output looks something like the following:
```bash
[
    {
        "id": "4cd98fc4d2a8",
        "name": "Sage-NEON-01",
        "status": "up"
    },
    {
        "id": "4cd98fc67b75",
        "name": "Sage-NEON-02",
        "status": "up"
    }
]
```
Currently, the only `metadata_name` we are supporting are the following:
- name
- status
- lat
- long

Add a new node or new metadata:
```bash
curl --location --request POST 'http://localhost:8080/api/v1/node?nodeid=4cd98fc67b75&metadata_name=name&metadata_value=Sage-NEON-02'
```
Output:
```bash
{
  "nodeID": "4cd98fc67b75",
  "metadataName": "name",
  "metadataValue": "Sage-NEON-02"
}
```
