FROM golang:1.24-alpine3.21 AS builder

RUN apk add --update  && \
    apk add --no-cache alpine-conf tzdata git

ADD ./src /go/src/delivery-service
ADD ./src/log /go/log
ADD ./src/migrations /go/migrations
ADD ./src/config /go/config

RUN cd /go/src/delivery-service && \
    go install delivery-service

FROM alpine:3.18.4 AS app

COPY --from=builder /go/bin/* /go/bin/delivery-service
COPY --from=builder /go/log /go/log
COPY --from=builder /go/migrations /go/migrations
COPY --from=builder /go/config /go/config

ENV CONTAINER_NAME=delivery-service
ENV HAS_WRITE_LOG_TO_FILE=false
ENV LOG_LEVEL=5

EXPOSE 8080

WORKDIR "/go"
ENTRYPOINT ["/go/bin/delivery-service"]