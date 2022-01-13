up:
	docker-compose up -d

down:
	docker-compose down --remove-orphans

start:
	go run .

request:
	curl --request POST --url http://localhost:8080/users --header 'Accept: application/json' --data '{"id":"test"}'

fmt:
	go fmt ./...