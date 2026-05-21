package employee

import (
	"context"
	"rest-api-hitalent/internal/model"
	"rest-api-hitalent/internal/service/department"
)

type Repository interface {
	Create(ctx context.Context, employee model.Employee) (model.Employee, error)
}

type Service struct {
	repository           Repository
	departmentRepository department.Repository
}

func NewService(repository Repository, departmentRepository department.Repository) *Service {
	return &Service{repository: repository, departmentRepository: departmentRepository}
}

func (s *Service) Create(ctx context.Context, data model.NewEmployeeData) (model.Employee, error) {
	employee, err := model.NewEmployee(data)
	if err != nil {
		return model.Employee{}, err
	}

	_, err = s.departmentRepository.Get(ctx, employee.DepartmentId)
	if err != nil {
		return model.Employee{}, err
	}

	return s.repository.Create(ctx, employee)
}
