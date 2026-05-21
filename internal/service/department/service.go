package department

import (
	"context"
	"rest-api-hitalent/internal/errors"
	"rest-api-hitalent/internal/model"
)

type Repository interface {
	Create(ctx context.Context, department model.Department) (model.Department, error)
	GetAllChildrenWithDepth(ctx context.Context, id uint, depth uint) ([]model.Department, error)
	Delete(ctx context.Context, ids []uint) error
	Update(ctx context.Context, department model.Department) (model.Department, error)
	Get(ctx context.Context, id uint) (model.Department, error)
}

type EmployeeRepository interface {
	GetAllByDepartmentIds(ctx context.Context, departmentIds []uint) ([]model.Employee, error)
	DeleteByDepartmentIds(ctx context.Context, departmentIds []uint) error
	Reassign(ctx context.Context, oldId, newId uint) error
}

type Service struct {
	repository         Repository
	employeeRepository EmployeeRepository
}

func NewService(repository Repository, employeeRepository EmployeeRepository) *Service {
	return &Service{repository: repository, employeeRepository: employeeRepository}
}

func (s *Service) Create(ctx context.Context, name string, parentId uint) (model.Department, error) {
	department, err := model.NewDepartment(name, parentId)
	if err != nil {
		return model.Department{}, err
	}

	return s.repository.Create(ctx, department)
}

func (s *Service) Get(ctx context.Context, id uint, depth uint, includeEmployees bool) (*model.DepartmentDetail, error) {
	if depth == 0 {
		depth = 1
	} else if depth > 5 {
		depth = 5
	}

	departments, err := s.repository.GetAllChildrenWithDepth(ctx, id, depth)
	if err != nil {
		return nil, err
	}

	res := &model.DepartmentDetail{}

	ids := make([]uint, 0, len(departments))
	for _, department := range departments {
		if department.ParentId == 0 {
			res.Department = department
		} else {
			res.Children = append(res.Children, department)
		}
		ids = append(ids, department.Id)
	}

	if includeEmployees {
		employees, err := s.employeeRepository.GetAllByDepartmentIds(ctx, ids)
		if err != nil {
			return nil, err
		}

		res.Employees = employees
	}

	return res, nil
}

func (s *Service) Delete(ctx context.Context, id uint, mode model.DepartmentDeleteMode, reassignToId uint) error {
	if mode == model.DepartmentDeleteModeCascade {
		departments, err := s.repository.GetAllChildrenWithDepth(ctx, id, 0)
		if err != nil {
			return err
		}

		ids := make([]uint, 0, len(departments))
		for _, d := range departments {
			ids = append(ids, d.Id)
		}

		err = s.employeeRepository.DeleteByDepartmentIds(ctx, ids)
		if err != nil {
			return err
		}

		return s.repository.Delete(ctx, ids)
	}

	if mode == model.DepartmentDeleteModeReassign {
		err := s.employeeRepository.Reassign(ctx, id, reassignToId)
		if err != nil {
			return err
		}

		return s.repository.Delete(ctx, []uint{id})
	}

	return nil
}

func (s *Service) Update(ctx context.Context, id uint, name *string, parentId *uint) (model.Department, error) {
	department, err := s.repository.Get(ctx, id)
	if err != nil {
		return model.Department{}, err
	}

	if name != nil {
		department.Name = model.Title(*name)
	}

	if parentId != nil {
		newParentId := *parentId

		if newParentId == id {
			return model.Department{}, errors.ErrDepartmentSelfParent
		}

		if newParentId != department.ParentId {
			children, err := s.repository.GetAllChildrenWithDepth(ctx, id, 0)
			if err != nil {
				return model.Department{}, err
			}

			for _, child := range children {
				if child.Id == newParentId {
					return model.Department{}, errors.ErrDepartmentCycle
				}
			}
		}

		department.ParentId = newParentId
	}

	return s.repository.Update(ctx, department)
}
