import asyncio

import sys
sys.path.append("./pypkg/bss-workplace-api/dist/grpc_bss_workplace_api-1.0.0-py3.8.egg")

from grpclib.client import Channel

from ozonmp.bss_workplace_api.v1.bss_workplace_api_grpc import BssWorkplaceApiServiceStub
from ozonmp.bss_workplace_api.v1.bss_workplace_api_pb2 import DescribeWorkplaceV1Request

async def main():
    async with Channel('127.0.0.1', 8082) as channel:
        client = BssWorkplaceApiServiceStub(channel)

        req = DescribeWorkplaceV1Request(workplace_id=1)
        reply = await client.DescribeWorkplaceV1(req)
        print(reply.message)


if __name__ == '__main__':
    asyncio.run(main())
