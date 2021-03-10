FROM golang:1.16.0-alpine3.13 AS builder
ENV GO111MODULE=on CGO_ENABLED=0

WORKDIR /go/src/app
COPY . .

RUN go build \
  -a \
  -trimpath \
  -ldflags "-s -w -extldflags '-static'" \
  -o /bin/main .

FROM alpine:3.13.2

COPY --from=builder /bin/main /bin/main

ENTRYPOINT ["/bin/main"]