# Scan Kubernetes Images from Manifests

[![Go Report Card](https://goreportcard.com/badge/github.com/yardenshoham/skim)](https://goreportcard.com/report/github.com/yardenshoham/skim)

skim is a CLI tool that extracts a list of container images from Kubernetes resources.

# Usage

```bash
skim list <path to k8s manifests>
```

# Build this project

```bash
go build
```

# Run tests

```bash
go test ./...
```

# Run this project

```bash
skim list ./testdata/deployment.yaml
```

## Docker

Docker images are available at
[DockerHub](https://hub.docker.com/r/yardenshoham/skim)
(docker.io/yardenshoham/skim).

Available docker tags

| Tag      | Description                       |
| -------- | --------------------------------- |
| `latest` | latest available release of skim. |
| `va.b.c` | skim version `a.b.c` .            |
| `a.b.c`  | skim version `a.b.c` .            |

### Docker run

```shell script
docker run \
    -v <your input path>:/input \
    yardenshoham/skim:latest list /input/*
```

### Docker build

You can build an own docker image by running

```shell
CGO_ENABLED=0 go build && docker build -t skim .
```
