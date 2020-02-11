FROM golang:1.13.7-alpine3.11

RUN apk update && apk add dep git
RUN mkdir -p /go/src/api
WORKDIR /go/src/api
COPY ./ /go/src/api

ENTRYPOINT ["sh", "-c"]
CMD ["dep ensure && go run scale.go"]
