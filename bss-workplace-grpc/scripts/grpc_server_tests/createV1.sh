#!/bin/bash
GRPCURL_BIN=/home/maxima/soft/grpcurl/grpcurl

$GRPCURL_BIN -plaintext -d '{"name": "Name 123", "size": 10}'  0.0.0.0:8082 ozonmp.bss_workplace_api.v1.BssWorkplaceApiService/CreateWorkplaceV1