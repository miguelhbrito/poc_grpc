run-dennis-gateway:
	echo "running the api server"
	sudo ./scripts/run-server-grpc.sh
config-up:
	sudo docker-compose up -d
config-down:
	docker-compose down