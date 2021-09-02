FROM golang:1.16-alpine AS builder

LABEL vendor=thelitdeveloper
LABEL maintainer=SimplySteph

RUN set -ex; \
    apk update; \
    apk add --no-cache git

WORKDIR ${GOPATH}/src/thelitdeveloper.com/bookstore-api/

ARG tag

ENV CGO_ENABLED=0 TAG=${tag}

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN cd cmd/bookstore && go build -v -o /api-server .

FROM alpine:3.14

COPY --from=builder /api-server /api-server

EXPOSE 8080

CMD ["/api-server"]
