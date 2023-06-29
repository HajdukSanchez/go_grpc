package repository

import (
	"context"

	"github.com/hajduksanchez/go_grpc/models"
)

// Abstract definition of the repository
type Repository interface {
	GetStudent(ctx context.Context, id string) (*models.Student, error)
	SetStudent(ctx context.Context, student *models.Student) error
	GetTest(ctx context.Context, id string) (*models.Test, error)
	SetTest(ctx context.Context, test *models.Test) error
	SetQuestions(ctx context.Context, question *models.Question) error
	SetEnrollmentStudent(ctx context.Context, enrollment *models.Enrollment) error
	GetStudentsPerTest(ctx context.Context, testId string) ([]*models.Student, error)
}

var implementation Repository

// Injection dependency
func SetRepository(repo Repository) {
	implementation = repo
}

func GetStudent(ctx context.Context, id string) (*models.Student, error) {
	return implementation.GetStudent(ctx, id)
}

func SetStudent(ctx context.Context, student *models.Student) error {
	return implementation.SetStudent(ctx, student)
}

func GetTest(ctx context.Context, id string) (*models.Test, error) {
	return implementation.GetTest(ctx, id)
}

func SetTest(ctx context.Context, test *models.Test) error {
	return implementation.SetTest(ctx, test)
}

func SetQuestions(ctx context.Context, question *models.Question) error {
	return implementation.SetQuestions(ctx, question)
}

func SetEnrollmentStudent(ctx context.Context, enrollment *models.Enrollment) error {
	return implementation.SetEnrollmentStudent(ctx, enrollment)
}

func GetStudentsPerTest(ctx context.Context, testId string) ([]*models.Student, error) {
	return implementation.GetStudentsPerTest(ctx, testId)
}
