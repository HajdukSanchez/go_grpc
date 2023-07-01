package main

import (
	"context"
	"log"
	"time"

	"github.com/hajduksanchez/go_grpc/proto/testpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Create a conection without credentials
	cc, err := grpc.Dial("localhost:5070", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close() // Close connection at the end

	client := testpb.NewTestServiceClient(cc) // Create a new client for the test service from GRPC
	// DoUnary(client)                           // Call the unary function
	DoClientStreamin(client) // Call streaming function
}

// Unary function call to the server
func DoUnary(c testpb.TestServiceClient) {
	// Craete a new request
	request := &testpb.GetTestRequest{
		Id: "t1",
	}

	response, err := c.GetTest(context.Background(), request)
	if err != nil {
		log.Fatalf("Error while calling GetTest RPC: %v", err)
	}

	log.Printf("Response from GetTest: %v", response)
}

// Streaming function call to the server
func DoClientStreamin(c testpb.TestServiceClient) {
	questions := []*testpb.Question{
		{
			Id:       "q7t1",
			Answer:   "This is the first answer",
			Question: "This is the first question",
			TestId:   "t1",
		},
		{
			Id:       "q8t1",
			Answer:   "This is the second answer",
			Question: "This is the second question",
			TestId:   "t1",
		},
		{
			Id:       "q9t1",
			Answer:   "This is the third answer",
			Question: "This is the third question",
			TestId:   "t1",
		},
	}

	// Set the stream
	stream, err := c.SetQuestions(context.Background())
	if err != nil {
		log.Fatalf("Error while calling SetQuestions: %v", err)
	}

	// Iterate over the questions and send them
	for _, question := range questions {
		log.Println("Sending question: ", question.Id)
		stream.Send(question)
		time.Sleep(2 * time.Second) // Sleep for 2 seconds
	}

	msg, err := stream.CloseAndRecv() // Close the stream and receive the response
	if err != nil {
		log.Fatalf("Error while receiving response from SetQuestions: %v", err)
	}
	// Finally get response from streaming service
	log.Printf("Response from SetQuestions: %v", msg)
}
