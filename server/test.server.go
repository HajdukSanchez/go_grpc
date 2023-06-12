package server

import (
	"context"

	"hajduksanchez.com/go/grpc/models"
	"hajduksanchez.com/go/grpc/proto/testpb"
	"hajduksanchez.com/go/grpc/repository"
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

	return &testpb.SetTestResponse{Id: test.Id}, nil
}
