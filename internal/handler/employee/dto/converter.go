package dto

import "rest-api-hitalent/internal/model"

func ToCreateEmployeeResponse(employee model.Employee) CreateEmployeeResponse {
	return CreateEmployeeResponse{
		Id:           employee.Id,
		DepartmentId: employee.DepartmentId,
		FullName:     string(employee.FullName),
		Position:     string(employee.Position),
		HiredAt:      employee.HiredAt,
		CreatedAt:    employee.CreatedAt,
	}
}
