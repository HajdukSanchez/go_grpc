package server

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/hajduksanchez/go_grpc/models"
	"github.com/hajduksanchez/go_grpc/proto/studentpb"
	"github.com/hajduksanchez/go_grpc/proto/testpb"
	"github.com/hajduksanchez/go_grpc/repository"
)

type TestServer struct {
	repository repository.Repository
	// Composition over inheritance to match with specification created on Proto files
	testpb.UnimplementedTestServiceServer
}

// Dependency injection for the repository
func NewTestServer(repository repository.Repository) *TestServer {
	return &TestServer{repository: repository}
}

func (server *TestServer) GetTest(ctx context.Context, request *testpb.GetTestRequest) (*testpb.Test, error) {
	test, err := server.repository.GetTest(ctx, request.GetId())
	if err != nil {
		return nil, err
	}

	return &testpb.Test{
		Id:   test.Id,
		Name: test.Name,
	}, nil
}

func (server *TestServer) SetTest(ctx context.Context, request *testpb.Test) (*testpb.SetTestResponse, error) {
	test := &models.Test{
		Id:   request.GetId(),
		Name: request.GetName(),
	}

	err := server.repository.SetTest(ctx, test)
	if err != nil {
		return nil, err
	}

	return &testpb.SetTestResponse{Id: test.Id, Name: test.Name}, nil
}

func (server *TestServer) SetQuestions(stream testpb.TestService_SetQuestionsServer) error {
	// Loop because we are receiving a stream of questions
	for {
		msg, err := stream.Recv() // Program will be locked here until receive a new message
		// This error means user has closed the stream and want his return answer
		// This is becasuse here we are not goinf to handle an error
		if err == io.EOF {
			// return information to the client and close server connection
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: true,
			})
		}
		// Here we validate another errors
		if err != nil {
			return err
		}

		// Here we can save the question in the database
		question := &models.Question{
			Id:       msg.GetId(),
			Answers:  msg.GetAnswer(),
			Question: msg.GetQuestion(),
			TestId:   msg.GetTestId(),
		}

		err = server.repository.SetQuestions(context.Background(), question)
		// Validate if there is no error saving the question
		if err != nil {
			// There was an error
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: false,
			})
		}
	}
}

func (server *TestServer) EnrollStudent(stream testpb.TestService_EnrollStudentServer) error {
	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&testpb.SetEnrollStudentResponse{
				Ok: true,
			})
		}
		if err != nil {
			return err
		}

		enrollment := &models.Enrollment{
			StudentId: msg.GetStudentId(),
			TestId:    msg.GetTestId(),
		}
		err = server.repository.EnrollStudent(context.Background(), enrollment)

		if err != nil {
			return stream.SendAndClose(&testpb.SetEnrollStudentResponse{
				Ok: false,
			})
		}
	}
}

func (server *TestServer) GetStudentsPerTest(request *testpb.GetStudentsPerTestRequest, stream testpb.TestService_GetStudentsPerTestServer) error {
	students, err := server.repository.GetStudentsPerTest(context.Background(), request.GetTestId())

	if err != nil {
		return err
	}

	for _, student := range students {
		student := &studentpb.Student{
			Id:   student.Id,
			Name: student.Name,
			Age:  student.Age,
		}
		err := stream.Send(student)
		time.Sleep(2 * time.Second) // Delay of 2 seconds for each stream send

		if err != nil {
			return err
		}
	}
	return nil
}

func (server *TestServer) TakeTest(stream testpb.TestService_TakeTestServer) error {
	msg, err := stream.Recv() // TODO: Validate if this works or change as teacher
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}

	questions, err := server.repository.GetQuestionsPerTest(context.Background(), msg.GetTestId())
	if err != nil {
		return err
	}

	i := 0                                   // Indicates the amount of questions answered
	var currentQuestion = &models.Question{} // Current questions send

	for {
		if i < len(questions) {
			currentQuestion = questions[i] // Set the current question
		}

		if i <= len(questions) {
			questionToSend := &testpb.Question{
				Id:       currentQuestion.Id,
				Question: currentQuestion.Question,
			}
			err := stream.Send(questionToSend) // Send the question to the client

			if err != nil {
				return err
			}

			i++
		}

		answer, err := stream.Recv() // Receive the answer from the client
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		log.Println("Answer received: ", answer.GetAnswer())
	}
}
