# poc_grpc

An API gRPC service example with observability features as custom logs, Jaeger, JWT token auth and NATS(TODO).

## 🧰 Configuration

To install golang just follow the steps from website:
- https://golang.org/doc/install

To install docker and docker-compose just follow the steps from website:
- https://docs.docker.com/engine/install/
- https://docs.docker.com/compose/install/

To install bloomRPC:
``` powershell
git clone https://github.com/uw-labs/bloomrpc.git
cd bloomrpc

yarn install && ./node_modules/.bin/electron-rebuild
npm run package
```
Start database server postgresql:
``` powershell
make config-up
```
To stop database server postgresql:
``` powershell
make config-down
```
To start application:
``` powershell
make run-server-grpc
```

## 🛠 How to use
- Import notebook.proto and login.proto into your Bloom:
<img src="https://github.com/miguelhbrito/poc_grpc/blob/master/images/bloomImportProtos.png" width="646" height="298">

- Make a request to create a new login user:
<img src="https://github.com/miguelhbrito/poc_grpc/blob/master/images/bloomCreateLogin.png" width="451" height="232">

- Then make a request to get a valid token:
<img src="https://github.com/miguelhbrito/poc_grpc/blob/master/images/bloomTokenLogin.png" width="844" height="231">

- Now you have a valid token to make a request to any notebook endpoint, just add a metadata into your request:
<img src="https://github.com/miguelhbrito/poc_grpc/blob/master/images/bloomMetadata.png" width="690" height="316">

``` powershell
"authorization":"token"
```
