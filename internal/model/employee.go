package model

import (
	"rest-api-hitalent/internal/errors"
	"time"
)

type Employee struct {
	Id           uint
	DepartmentId uint
	FullName     Title
	Position     Title
	HiredAt      time.Time
	CreatedAt    time.Time
}

type NewEmployeeData struct {
	DepartmentId uint
	FullName     string
	Position     string
	HiredAt      time.Time
}

func NewEmployee(data NewEmployeeData) (Employee, error) {
	fullName, err := NewTitle(data.FullName, TitleConfig{
		MinLength: 1,
		MaxLength: 200,
		Error:     errors.ErrEmployeeInvalidFullName,
	})
	if err != nil {
		return Employee{}, err
	}

	position, err := NewTitle(data.Position, TitleConfig{
		MinLength: 1,
		MaxLength: 200,
		Error:     errors.ErrEmployeeInvalidPosition,
	})
	if err != nil {
		return Employee{}, err
	}

	if data.DepartmentId == 0 {
		return Employee{}, errors.ErrEmployeeInvalidDepartmentId
	}

	return Employee{
		DepartmentId: data.DepartmentId,
		FullName:     fullName,
		Position:     position,
		HiredAt:      data.HiredAt,
	}, nil
}
