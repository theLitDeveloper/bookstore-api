# Bookstore API
This is a fairly simple RESTful service for testing / evaluation purposes. It works out of the box without any external dependencies and comes with support for Prometheus.

The service exposes four endpoints:

- **/books**    Bookstore API
- **/metrics**  Prometheus metrics
- **/version**  Version info
- **/health**   Health check

**A OpenAPI v3 doc for bookstore can be found here: [bookstore.yaml](api/bookstore.yaml)**

## Usage
Set nessessary env var for version info first:
```
export LATEST_GIT_TAG=$(git describe --tags --abbrev=0)
```

## Docker
Build an image:
```
docker build --build-arg latest_git_tag=${LATEST_GIT_TAG} -t bookstore-api:${LATEST_GIT_TAG} .
```
run a container and verify it:
```
docker run -it --rm -d --name bookstore -p 8080:8080 -e LATEST_GIT_TAG=${LATEST_GIT_TAG} bookstore-api:${LATEST_GIT_TAG}

curl http://localhost:8080/version
```
You should see a JSON object like this:
```
{"version":"v0.2.3"}
```
