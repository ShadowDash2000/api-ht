package employee

import (
	"context"
	"rest-api-hitalent/internal/model"
	"rest-api-hitalent/internal/repository/employee/record"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, employee model.Employee) (model.Employee, error) {
	rec := record.ToEmployeeRecord(employee)

	err := gorm.G[record.EmployeeRecord](r.db).Create(ctx, &rec)
	if err != nil {
		return model.Employee{}, err
	}

	return record.ToEmployee(rec), nil
}

func (r *repository) GetAllByDepartmentIds(ctx context.Context, departmentIds []uint) ([]model.Employee, error) {
	records, err := gorm.G[record.EmployeeRecord](r.db).
		Where("department_id IN ?", departmentIds).
		Order("created_at DESC").
		Find(ctx)
	if err != nil {
		return nil, err
	}

	employees := make([]model.Employee, 0, len(records))
	for _, rec := range records {
		employees = append(employees, record.ToEmployee(rec))
	}

	return employees, nil
}

func (r *repository) DeleteByDepartmentIds(ctx context.Context, departmentIds []uint) error {
	_, err := gorm.G[record.EmployeeRecord](r.db).
		Where("department_id IN ?", departmentIds).
		Delete(ctx)
	return err
}

func (r *repository) Reassign(ctx context.Context, oldId, newId uint) error {
	_, err := gorm.G[record.EmployeeRecord](r.db).
		Where("department_id = ?", oldId).
		Update(ctx, "department_id", newId)
	return err
}
