package services

import (
	"context"
	"practical_exercise/internal/databases"
	"practical_exercise/internal/httpbody/request"
	"practical_exercise/internal/httpbody/response"
	"practical_exercise/internal/models"

	"github.com/google/uuid"
)

type IStudent interface {
	Create(ctx context.Context, student request.CreateStudent) (response.User, error)
}

type Student struct {
	dbStore databases.DBStore
}

var userInstance IStudent

func NewUser(dbStore databases.DBStore) IStudent {
	if userInstance == nil {
		userInstance = &Student{
			dbStore: dbStore,
		}
	}
	return userInstance
}

func (student *Student) Create(ctx context.Context, requestStudent request.CreateStudent) (respStudent response.User, err error) {
	modelStudent := models.Student{
		Id:        uuid.New().String(),
		Name:      requestStudent.Name,
		Email:     requestStudent.Email,
		Pass:      requestStudent.Pass,
		StudentId: requestStudent.StudentId,
	}

	modelStudent, err = student.dbStore.Student.Create(ctx, modelStudent)

	if err != nil {
		return respStudent, err
	}

	respStudent = response.User{
		Id:        modelStudent.Id,
		StudentId: modelStudent.StudentId,
		Name:      modelStudent.Name,
		Email:     modelStudent.Email,
		CreatedAt: modelStudent.CreatedAt.String(),
		UpdatedAt: modelStudent.UpdatedAt.String(),
	}
	return respStudent, nil
}
