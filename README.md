# Purpose Backend

## Installation

```
brew install protobuf
brew install protoc-gen-go
brew install protoc-gen-go-grpc
brew install protoc-gen-grpc-web
```

## Code Generation Commands

```
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    gRPC/purpose.proto
```

## Docker

```
docker build \
    --build-arg PAT=${PAT} \
    --build-arg GOPRIVATE=${GOPRIVATE} \
    --build-arg HOST=host.docker.internal \
    --build-arg COINBASE_PASSPHRASE=${COINBASE_PASSPHRASE} \
    --build-arg COINBASE_SECRET=${COINBASE_SECRET} \
    --build-arg COINBASE_KEY=${COINBASE_KEY} \
    -t $(basename $PWD) .
```

```
docker run \
    -p 80:80 \
    -d --restart unless-stopped \
    $(basename $PWD)
```
