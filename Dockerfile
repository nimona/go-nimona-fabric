FROM golang:1.14-buster AS builder

WORKDIR /src/nimona.io

ENV CGO_ENABLED=1

ADD . .

RUN make build

###

FROM debian:buster-slim

COPY --from=builder /src/nimona.io/bin/nimona /nimona

ENTRYPOINT ["/nimona"]
