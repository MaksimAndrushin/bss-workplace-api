#!/bin/bash
GRPCURL_BIN=/home/maxima/soft/grpcurl/grpcurl

$GRPCURL_BIN -d '{"offset": "0", "limit": 10}' -plaintext 0.0.0.0:8083 ozonmp.bss_workplace_facade.v1.BssFacadeEventsApiService/ListEventsV1