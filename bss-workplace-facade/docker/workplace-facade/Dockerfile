# Builder
FROM golang:1.16-alpine AS builder

ARG GITHUB_PATH=github.com/ozonmp/bss-workplace-facade

WORKDIR /home/${GITHUB_PATH}

RUN apk add --update make git protoc protobuf protobuf-dev curl

COPY Makefile Makefile

RUN make deps-go

COPY . .

RUN make build-go

# gRPC Server

FROM alpine:latest as server

ARG GITHUB_PATH=github.com/ozonmp/bss-workplace-facade

LABEL org.opencontainers.image.source=https://${GITHUB_PATH}

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /home/${GITHUB_PATH}/bin/bss-workplace-facade .
COPY --from=builder /home/${GITHUB_PATH}/facade-config.yml .
COPY --from=builder /home/${GITHUB_PATH}/migrations/ ./migrations

RUN chown root:root bss-workplace-facade

EXPOSE 50051
EXPOSE 8080
EXPOSE 9100

CMD ["./bss-workplace-facade"]
