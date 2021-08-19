# Bookstore Sample Service
This is a simple mock service for testing / evaluation purposes. It works out of the box without any external dependencies and comes with support for Prometheus.

The service exposes four endpoints:
|||
|-|-|
|/books|Bookstore|
|/metrics|Prometheus metrics|
|/version|Version info|
|/health|Simple health check|

A OpenAPI v3 doc for bookstore can be found here: [bookstore.yaml](bookstore.yaml)

## Usage
Use Docker to build an image:
```
docker build -t <image>:<tag> .
```
and run a container and verify it:
```
docker run -it -d --name <name> -p 8080:8080 <image>:<tag> \
&& curl http://localhost:8080/health
```