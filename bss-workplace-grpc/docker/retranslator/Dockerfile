# Builder
FROM golang:1.16-alpine AS builder

ARG GITHUB_PATH=github.com/ozonmp/bss-workplace-api

WORKDIR /home/${GITHUB_PATH}

RUN apk add --update make git protoc protobuf protobuf-dev curl

COPY Makefile Makefile
RUN make deps-go

COPY . .
#COPY go.mod /home/${GITHUB_PATH}/go.mod
#COPY go.sum /home/${GITHUB_PATH}/go.sum


RUN make build-go-retranslator

# Retranslator

FROM alpine:latest as server

ARG GITHUB_PATH=github.com/ozonmp/bss-workplace-api

LABEL org.opencontainers.image.source=https://${GITHUB_PATH}
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /home/${GITHUB_PATH}/bin/retranslator .
COPY --from=builder /home/${GITHUB_PATH}/retranslator-config.yml .
#COPY --from=builder /home/${GITHUB_PATH}/migrations/ ./migrations

RUN chown root:root retranslator

EXPOSE 50051
EXPOSE 8080
EXPOSE 9100

CMD ["./retranslator"]
