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
Get all the nodes data:
```bash
curl GET 'http://localhost:8080/api/v1/nodes-data'
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
    {
      "name": "Sage-TTU-01",
      "id": "2cea7f5a0a3d",
      "status": "Up",
      "provisionDate": "11/19/20",
      "OSVersion": "dell-1.0.0.local-6da62a8",
      "serviceTag": "J42P853",
      "SpecialDevices": "N/A",
      "BiosVersion": "2.8.2",
      "lat": "33.584306",
      "lon": "-101.871984"
    }
]
```

Get all the nodes metadata:
```bash
curl GET 'http://localhost:8080/api/v1/nodes-metadata'
```

Returns:
```bash
[
    {
      "id": "name",
      "type": "text",
      "label": "name",
      "description": "Name of node."
    },
    ...
    {
      "id": "lon",
      "type": "numeric",
      "label": "Longitude",
      "description": "Longitude of Node."
    }
]
```

Get all the nodes data and metadata:
```bash
curl GET 'http://localhost:8080/api/v1/nodes-all'
```

Returns:
```bash
[
{
  "data": [
    {
      "name": "Sage-NEON-01",
      "id": "4cd98fc4d2a8",
      "status": "Up",
      "provisionDate": "10/22/20",
      "OSVersion": "dell-1.0.0.local-6da62a8",
      "serviceTag": "79BBZ23",
      "SpecialDevices": "N/A",
      "BiosVersion": "2.8.2",
      "lat": "40.01631",
      "lon": "-105.24585"
    },
    ...
    {
      "name": "Sage-TTU-01",
      "id": "2cea7f5a0a3d",
      "status": "Up",
      "provisionDate": "11/19/20",
      "OSVersion": "dell-1.0.0.local-6da62a8",
      "serviceTag": "J42P853",
      "SpecialDevices": "N/A",
      "BiosVersion": "2.8.2",
      "lat": "33.584306",
      "lon": "-101.871984"
    }
  ],
  "metadata": [
    {
      "id": "name",
      "type": "text",
      "label": "name",
      "description": "Name of node."
    },
    ...
    {
      "id": "lon",
      "type": "numeric",
      "label": "Longitude",
      "description": "Longitude of Node."
    }
  ]
}
```