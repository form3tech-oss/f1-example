package main

import (
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var (
	sqsClient *sqs.SQS
	queueUrl  *string
)

func main() {
	goawsEndpoint := "http://goaws:4100"
	region := "eu-west-1"
	s, err := session.NewSession(&aws.Config{
		Endpoint:    &goawsEndpoint,
		Region:      &region,
		Credentials: credentials.NewStaticCredentials("foo", "bar", ""),
	})
	if err != nil {
		panic(err)
	}
	sqsClient = sqs.New(s)
	queueName := "test-queue"
	queueUrlResponse, err := sqsClient.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: &queueName,
	})
	if err != nil {
		panic(err)
	}
	queueUrl = queueUrlResponse.QueueUrl
	http.HandleFunc("/payments", paymentsHandler)
	http.ListenAndServe(":8080", nil)
}

func paymentsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("request received with invalid http method: %s", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	log.Println("http POST request received on /payments")
	message := "test message"
	_, err := sqsClient.SendMessage(&sqs.SendMessageInput{
		MessageBody: &message,
		QueueUrl:    queueUrl,
	})
	if err != nil {
		log.Printf("error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}
