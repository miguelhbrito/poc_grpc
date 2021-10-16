run-dennis-gateway:
	echo "running the api server"
	sudo docker-compose up -d
	sudo ./scripts/run-server-grpc.sh

config-down:
	docker-compose down