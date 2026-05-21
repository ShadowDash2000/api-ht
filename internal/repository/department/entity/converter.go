package entity

import "rest-api-hitalent/internal/model"

func ToDepartmentRecord(department model.Department) DepartmentRecord {
	var parentId *uint
	if department.ParentId != 0 {
		parentId = &department.ParentId
	}

	return DepartmentRecord{
		Id:        department.Id,
		Name:      string(department.Name),
		ParentId:  parentId,
		CreatedAt: department.CreatedAt,
	}
}

func ToDepartment(departmentRecord DepartmentRecord) model.Department {
	department := model.Department{
		Id:        departmentRecord.Id,
		Name:      model.Title(departmentRecord.Name),
		CreatedAt: departmentRecord.CreatedAt,
	}

	if departmentRecord.ParentId != nil {
		department.ParentId = *departmentRecord.ParentId
	}

	return department
}

func ToDepartments(departmentRecords []DepartmentRecord) []model.Department {
	if len(departmentRecords) == 0 {
		return nil
	}

	departments := make([]model.Department, 0, len(departmentRecords))
	for _, departmentRecord := range departmentRecords {
		departments = append(departments, ToDepartment(departmentRecord))
	}
	return departments
}
