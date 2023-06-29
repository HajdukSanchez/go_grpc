package database

import (
	"context"
	"database/sql"
	"log"

	"github.com/hajduksanchez/go_grpc/models"
	_ "github.com/lib/pq"
)

// Concrete definition of the Repository interface
type PostgresRepository struct {
	db *sql.DB
}

// Init the PostgresRepository concrete implementation
func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	return &PostgresRepository{db}, nil
}

// Concrete implementation for Set new student from Repository interface
func (repo *PostgresRepository) SetStudent(ctx context.Context, student *models.Student) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO students (id, name, age) VALUES ($1, $2, $3)", student.Id, student.Name, student.Age)
	return err
}

// Concrete implementation for Get new student from Repository interface
func (repo *PostgresRepository) GetStudent(ctx context.Context, id string) (*models.Student, error) {
	var student = models.Student{}

	err := repo.db.QueryRowContext(ctx, "SELECT id, name, age FROM students WHERE id = $1", id).Scan(&student.Id, &student.Name, &student.Age)
	if err != nil {
		return nil, err
	}

	return &student, nil
}

// Concrete implementation for Set new test for students from Repository interface
func (repo *PostgresRepository) SetTest(ctx context.Context, test *models.Test) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO tests (id, name) VALUES ($1, $2)", test.Id, test.Name)
	return err
}

// Concrete implementation for Get new test for student from Repository interface
func (repo *PostgresRepository) GetTest(ctx context.Context, id string) (*models.Test, error) {
	var test = models.Test{}

	err := repo.db.QueryRowContext(ctx, "SELECT id, name FROM tests WHERE id = $1", id).Scan(&test.Id, &test.Name)
	if err != nil {
		return nil, err
	}

	return &test, nil
}

// Set a new question for a test
func (repo *PostgresRepository) SetQuestions(ctx context.Context, question *models.Question) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO questions (id, answer, question, test_id) VALUES ($1, $2, $3, $4)", question.Id, question.Answers, question.Question, question.TestId)
	return err
}

// Enroll a new student to a specific examn (test)
func (repo *PostgresRepository) SetEnrollmentStudent(ctx context.Context, enrollment *models.Enrollment) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO enrollments (student_id, test_id) VALUES ($1, $2)", enrollment.StudentId, enrollment.TestId)
	return err
}

// Get all students enrolled to a specific examn (test)
func (repo *PostgresRepository) GetStudentsPerTest(ctx context.Context, testId string) ([]*models.Student, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name, age FROM students WHERE id in (SELECT student_id FROM enrollments WHERE test_id = $1)", testId)

	if err != nil {
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var students []*models.Student // Set array of students
	for rows.Next() {
		var student = models.Student{}
		if err = rows.Scan(&student.Id, &student.Name, &student.Age); err == nil {
			students = append(students, &student) // Add new entry into the array
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}
