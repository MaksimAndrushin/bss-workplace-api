#!/bin/bash
GRPCURL_BIN=/home/maxima/soft/grpcurl/grpcurl

$GRPCURL_BIN -plaintext -d '{"workplace_id": 10010}'  127.0.0.1:8082 ozonmp.bss_workplace_api.v1.BssWorkplaceApiService/DescribeWorkplaceV1
