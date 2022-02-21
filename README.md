# F1 example
This repo demonstrates the use of [`f1`](https://github.com/form3tech-oss/f1) to load test services which provide feedback to the caller asynchonously.

This is demonstrated in this repo by a service (`./cmd/service/main.go`) which accepts HTTP POST request on `/payments` and responds with a `202 Accepted` to indicate that background work will be taking place.

When this "payment" is processed, an SQS message is sent to a local mock of SQS that uses `goaws` (see `docker-compose.yml` file for details.

A load test has been written using `f1` (`./cmd/f1/main.go`) which exercises this flow by making an HTTP POST request, and then waiting for a corresponding message to be delivered on the SQS queue.

## Usage
1. Run the docker compose file: `docker-compose up -d`.
2. Run a load test `go run ./cmd/f1/main.go run constant -r 1/s -d 10s -v testScenario`.
