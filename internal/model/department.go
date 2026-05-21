package model

import (
	"rest-api-hitalent/internal/errors"
	"strings"
	"time"
)

type Department struct {
	Id        uint
	Name      Title
	ParentId  uint
	CreatedAt time.Time
}

func NewDepartment(name string, parentId uint) (Department, error) {
	fullName, err := NewTitle(strings.TrimSpace(name), TitleConfig{
		MinLength: 1,
		MaxLength: 200,
		Error:     errors.ErrDepartmentInvalidName,
	})
	if err != nil {
		return Department{}, err
	}

	return Department{
		Name:     fullName,
		ParentId: parentId,
	}, nil
}

type DepartmentDetail struct {
	Department
	Employees []Employee
	Children  []Department
}
