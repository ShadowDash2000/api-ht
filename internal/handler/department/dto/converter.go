package dto

import "rest-api-hitalent/internal/model"

func ToCreateDepartmentResponse(department model.Department) CreateDepartmentResponse {
	return CreateDepartmentResponse{
		Id:        department.Id,
		Name:      string(department.Name),
		ParentId:  department.ParentId,
		CreatedAt: department.CreatedAt,
	}
}

func ToDepartmentResponse(department model.Department) Department {
	return Department{
		Id:        department.Id,
		Name:      string(department.Name),
		ParentId:  department.ParentId,
		CreatedAt: department.CreatedAt,
	}
}

func ToEmployeeResponse(employee model.Employee) Employee {
	return Employee{
		Id:           employee.Id,
		DepartmentId: employee.DepartmentId,
		FullName:     string(employee.FullName),
		Position:     string(employee.Position),
		HiredAt:      employee.HiredAt,
		CreatedAt:    employee.CreatedAt,
	}
}

func ToGetDepartmentResponse(detail model.DepartmentDetail) GetDepartmentResponse {
	employees := make([]Employee, 0, len(detail.Employees))
	for _, employee := range detail.Employees {
		employees = append(employees, ToEmployeeResponse(employee))
	}

	children := make([]Department, 0, len(detail.Children))
	for _, child := range detail.Children {
		children = append(children, ToDepartmentResponse(child))
	}

	return GetDepartmentResponse{
		Department: ToDepartmentResponse(detail.Department),
		Employees:  employees,
		Children:   children,
	}
}
