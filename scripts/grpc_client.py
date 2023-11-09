import asyncio

from grpclib.client import Channel

from ozonmp.ozon_keyword_api.v1.ozon_keyword_api_grpc import OzonKeywordApiServiceStub
from ozonmp.ozon_keyword_api.v1.ozon_keyword_api_pb2 import DescribeKeywordV1Request

async def main():
    async with Channel('127.0.0.1', 8082) as channel:
        client = OzonKeywordApiServiceStub(channel)

        req = DescribeKeywordV1Request(keyword_id=1)
        reply = await client.DescribeKeywordV1(req)
        print(reply.message)


if __name__ == '__main__':
    asyncio.run(main())
