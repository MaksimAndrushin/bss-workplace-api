#!/bin/sh

goose -dir migrations \
  postgres "user=postgres password=postgres host=0.0.0.0 port=5432 database=bss_workplace_api sslmode=disable" \
  down