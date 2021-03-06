# Bookstore API
A fairly simple RESTful service for testing / evaluation purposes, e.g. on a K8s cluster. It works without any external dependencies and provides Prometheus out of the box.

The service exposes four endpoints:

- **/books**    Bookstore API
- **/metrics**  Prometheus metrics
- **/version**  Version info
- **/health**   Health check

**A OpenAPI v3 doc for bookstore can be found here: [bookstore.yaml](api/bookstore.yaml)**

## Usage
Set nessessary env var for version info first:
```
export TAG=$(git describe --tags --abbrev=0)
```

## Docker
Build an image:
```
docker build --build-arg tag=${TAG} -t bookstore-api:${TAG} .
```
run a container and verify it:
```
docker run -it --rm -d --name bookstore -p 8080:8080 -e TAG=${TAG} bookstore-api:${TAG}

curl http://localhost:8080/version
```
You should see a JSON object like this:
```
{"version":"v0.2.5"}
```

## Deployments
Edit deployments/k8s/bookstore.yml, find image and env in Deployment section and replace placeholders with your values:
```
---
apiVersion: apps/v1
kind: Deployment
...
spec:
...
  spec:
    containers:
      - name: bookstore-api
        image: <image>:<tag>
        env:
          - name: LATEST_GIT_TAG
            value: <tag>
```