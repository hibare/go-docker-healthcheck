# go-docker-healthcheck


[![Go Report Card](https://goreportcard.com/badge/github.com/hibare/go-docker-healthcheck)](https://goreportcard.com/report/github.com/hibare/go-docker-healthcheck)
[![GitHub issues](https://img.shields.io/github/issues/hibare/go-docker-healthcheck)](https://github.com/hibare/go-docker-healthcheck/issues)
[![GitHub pull requests](https://img.shields.io/github/issues-pr/hibare/go-docker-healthcheck)](https://github.com/hibare/go-docker-healthcheck/pulls)
[![GitHub](https://img.shields.io/github/license/hibare/go-docker-healthcheck)](https://github.com/hibare/go-docker-healthcheck/blob/main/LICENSE)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/hibare/go-docker-healthcheck)](https://github.com/hibare/go-docker-healthcheck/releases)


Healtch check functionality for Golang Docker images built from scratch.

Original source: https://medium.com/google-cloud/dockerfile-go-healthchecks-k8s-9a87d5c5b4cb


This is useful to add healthcheck in Golang docker images created from scratch.

### Usage Example

Fetch source code asset from Github release. Compile into a binary.

```
FROM golang:1.17.0-alpine AS base

# Build golang healthcheck binary
FROM base AS healthcheck

ARG VERSION=0.2.0

RUN wget -O - https://github.com/hibare/go-docker-healthcheck/archive/refs/tags/v${VERSION}.tar.gz |  tar zxf -

WORKDIR /go/go-docker-healthcheck-${VERSION}
```

Copy binary into final stage.

```
COPY --from=healthcheck /bin/healthcheck /bin/healthcheck
```

Add heathcheck command to docker. Replace the URL in healthcheck with actual application healthcheck URL.

```
HEALTHCHECK \
    --interval=30s \
    --timeout=3s \
    CMD ["healthcheck","http://localhost:5000/ping/"]
```

### Example
A working example can be found [here](https://github.com/hibare/go-container-status/blob/main/Dockerfile).