FROM golang:1.18-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/Swapica/order-indexer-svc
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/order-indexer-svc /go/src/github.com/Swapica/order-indexer-svc


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/order-indexer-svc /usr/local/bin/order-indexer-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["order-indexer-svc"]
