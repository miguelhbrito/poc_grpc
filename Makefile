run-dennis-gateway:
	echo "running the api server"
	chmod +x ./scripts/run-server-grpc.sh
	./scripts/run-server-grpc.sh
config-up:
	sudo docker-compose up -d
config-down:
	docker-compose down