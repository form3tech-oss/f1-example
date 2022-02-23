test:
	docker-compose up -d
	go run ./cmd/f1/main.go run constant -r 1/s -d 10s -v testScenario
