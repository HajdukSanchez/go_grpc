package server

import (
	"context"

	"github.com/hajduksanchez/go_grpc/models"
	"github.com/hajduksanchez/go_grpc/proto/studentpb"
	"github.com/hajduksanchez/go_grpc/repository"
)

type StudentServer struct {
	repository repository.Repository
	// Composition over inheritance to match with specification created on Proto files
	studentpb.UnimplementedStudentServiceServer
}

// Dependency injection for the repository
func NewStudentServer(repository repository.Repository) *StudentServer {
	return &StudentServer{repository: repository}
}

func (server *StudentServer) GetStudent(ctx context.Context, request *studentpb.GetStudentRequest) (*studentpb.Student, error) {
	student, err := server.repository.GetStudent(ctx, request.GetId())
	if err != nil {
		return nil, err
	}

	return &studentpb.Student{
		Id:   student.Id,
		Name: student.Name,
		Age:  student.Age,
	}, nil
}

func (server *StudentServer) SetStudent(ctx context.Context, request *studentpb.Student) (*studentpb.SetStudentResponse, error) {
	student := &models.Student{
		Id:   request.GetId(),
		Name: request.GetName(),
		Age:  request.GetAge(),
	}

	err := server.repository.SetStudent(ctx, student)
	if err != nil {
		return nil, err
	}

	return &studentpb.SetStudentResponse{Id: student.Id}, nil
}
