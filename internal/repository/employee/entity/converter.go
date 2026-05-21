package entity

import "rest-api-hitalent/internal/model"

func ToEmployeeRecord(data model.Employee) EmployeeRecord {
	return EmployeeRecord{
		Id:           data.Id,
		DepartmentId: data.DepartmentId,
		FullName:     string(data.FullName),
		Position:     string(data.Position),
		HiredAt:      data.HiredAt,
		CreatedAt:    data.CreatedAt,
	}
}

func ToEmployee(employeeRecord EmployeeRecord) model.Employee {
	return model.Employee{
		Id:           employeeRecord.Id,
		DepartmentId: employeeRecord.DepartmentId,
		FullName:     model.Title(employeeRecord.FullName),
		Position:     model.Title(employeeRecord.Position),
		HiredAt:      employeeRecord.HiredAt,
		CreatedAt:    employeeRecord.CreatedAt,
	}
}
