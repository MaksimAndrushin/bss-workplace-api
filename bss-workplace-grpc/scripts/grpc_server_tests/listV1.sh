#!/bin/bash
GRPCURL_BIN=/home/maxima/soft/grpcurl/grpcurl

$GRPCURL_BIN -d '{"offset": "0", "limit": 10}' -plaintext 0.0.0.0:18082 ozonmp.bss_workplace_api.v1.BssWorkplaceApiService/ListWorkplacesV1