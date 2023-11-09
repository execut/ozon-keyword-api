import setuptools

setuptools.setup(
    name="grpc-ozon-keyword-api",
    version="1.0.0",
    author="rusdevop",
    author_email="rusdevops@gmail.com",
    description="GRPC python client for ozon-keyword-api",
    url="https://github.com/execut/ozon-keyword-api",
    packages=setuptools.find_packages(),
    package_data={"ozonmp.ozon_keyword_api.v1": ["ozon_keyword_api_pb2.pyi"]},
    python_requires='>=3.5',
)