run-dennis-gateway:
	echo "running the api server"
	chmod +x ./scripts/run-server-grpc.sh
	./scripts/run-server-grpc.sh
config-up:
	sudo docker-compose up -d
config-down:
	docker-compose down
generate-proto:
	protoc --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb -I proto --go-grpc_opt=paths=source_relative proto/*.proto