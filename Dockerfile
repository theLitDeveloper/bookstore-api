FROM golang:1.16-alpine AS builder

LABEL vendor=SimplyDelivery

RUN set -ex; \
    apk update; \
    apk add --no-cache git

WORKDIR ${GOPATH}/src/gitlab.simplyadmin.de/sdGolang/bookstore-api/

ARG latest_git_tag

ENV CGO_ENABLED=0 LATEST_GIT_TAG=${latest_git_tag}

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN cd cmd/bookstore && go build -v -o /api-server .

FROM alpine:3.14

COPY --from=builder /api-server /api-server

EXPOSE 8080

CMD ["/api-server"]
