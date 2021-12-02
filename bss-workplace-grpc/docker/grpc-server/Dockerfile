# Builder
FROM golang:1.16-alpine AS builder

ARG GITHUB_PATH=github.com/ozonmp/bss-workplace-api

WORKDIR /home/${GITHUB_PATH}

RUN apk add --update make git protoc protobuf protobuf-dev curl

COPY Makefile Makefile

RUN make deps-go

COPY . .
#COPY go.mod /home/${GITHUB_PATH}
#COPY go.sum /home/${GITHUB_PATH}

RUN make build-go-grpc-api

# gRPC Server

FROM alpine:latest as server

ARG GITHUB_PATH=github.com/ozonmp/bss-workplace-api

LABEL org.opencontainers.image.source=https://${GITHUB_PATH}

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /home/${GITHUB_PATH}/bin/grpc-server .
COPY --from=builder /home/${GITHUB_PATH}/grpc-server-config.yml .
COPY --from=builder /home/${GITHUB_PATH}/migrations/ ./migrations

RUN chown root:root grpc-server

EXPOSE 50051
EXPOSE 8080
EXPOSE 9100

CMD ["./grpc-server"]
