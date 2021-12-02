# Builder
FROM golang:1.16-alpine AS builder

ARG GITHUB_PATH=github.com/ozonmp/bss-workplace-bot

WORKDIR /home/${GITHUB_PATH}

RUN apk add --update make git protoc protobuf protobuf-dev curl

COPY Makefile Makefile

RUN make deps-go

COPY . .

RUN make build-go

# TG Bot

FROM alpine:latest as server

ARG GITHUB_PATH=github.com/ozonmp/bss-workplace-bot

LABEL org.opencontainers.image.source=https://${GITHUB_PATH}

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /home/${GITHUB_PATH}/bin/bot .
COPY --from=builder /home/${GITHUB_PATH}/bot-config.yml .

RUN chown root:root bot

EXPOSE 50051
EXPOSE 8080
EXPOSE 9100

CMD ["./bot"]
