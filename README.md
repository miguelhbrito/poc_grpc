# poc_grpc

An API gRPC service example, with observability feature as custom logs, Jaeger and NATS(TODO).

## 🧰 Instalação
To install bloomRPC:
``` powershell
git clone https://github.com/uw-labs/bloomrpc.git
cd bloomrpc

yarn install && ./node_modules/.bin/electron-rebuild
npm run package
```
To start application:
``` powershell
docker-compose up -d
go run main.go
```

## 🛠 How to use
<img src="./resources/editor-preview.gif" />
Import notebook.proto and login.proto and make requests to localhost:50051 with metadata:
``` powershell
authorization":"Z2FuZGFsZjptaXRocmFuZGly
```
