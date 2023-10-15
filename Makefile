run-compose:
	docker-compose -f ./deployment/docker-compose.yaml up -d
run-app:
	go run cmd/main.go