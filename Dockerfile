FROM golang:1.10 as builder
MAINTAINER Justin C. Miller <justin@echtra.net>

WORKDIR /go/src/github.com/echtra/netlogger
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build

FROM alpine:3.8

WORKDIR /app
CMD ["/app/netlogger"]

COPY --from=builder /go/src/github.com/echtra/netlogger/netlogger .

