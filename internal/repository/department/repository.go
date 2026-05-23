package department

import (
	"context"
	"errors"
	internalerrors "rest-api-hitalent/internal/errors"
	"rest-api-hitalent/internal/model"
	"rest-api-hitalent/internal/repository/department/record"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, department model.Department) (model.Department, error) {
	rec := record.ToDepartmentRecord(department)

	err := gorm.G[record.DepartmentRecord](r.db).Create(ctx, &rec)
	if err != nil {
		return model.Department{}, err
	}

	return record.ToDepartment(rec), nil
}

func (r *repository) Get(ctx context.Context, id uint) (model.Department, error) {
	rec, err := gorm.G[record.DepartmentRecord](r.db).
		Where("id = ?", id).
		First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Department{}, internalerrors.ErrDepartmentNotFound
		}
		return model.Department{}, err
	}

	return record.ToDepartment(rec), nil
}

// GetAllChildrenWithDepth
// Note: depth = 1 means only root department, depth = 0 means all children
func (r *repository) GetAllChildrenWithDepth(ctx context.Context, id uint, depth uint) ([]model.Department, error) {
	// Recursively searching for all children until there are no more children
	// or the depth is reached.
	query := `
		WITH RECURSIVE department_tree AS (
			SELECT *, 1 AS depth
			FROM department
			WHERE id = ?
			
			UNION ALL

			SELECT d.*, dt.depth + 1
			FROM department d
			INNER JOIN department_tree dt ON d.parent_id = dt.id
			WHERE ? = 0 OR dt.depth < ?
		)
		SELECT * FROM department_tree
	`

	records, err := gorm.G[record.DepartmentRecord](r.db).
		Raw(query, id, depth, depth).
		Find(ctx)
	if err != nil {
		return nil, err
	}

	if len(records) == 0 {
		return nil, internalerrors.ErrDepartmentNotFound
	}

	return record.ToDepartments(records), nil
}

func (r *repository) Delete(ctx context.Context, ids []uint) error {
	_, err := gorm.G[record.DepartmentRecord](r.db).
		Where("id IN ?", ids).
		Delete(ctx)
	return err
}

func (r *repository) Update(ctx context.Context, department model.Department) (model.Department, error) {
	rec := record.ToDepartmentRecord(department)

	_, err := gorm.G[record.DepartmentRecord](r.db).
		Where("id = ?", rec.Id).
		Updates(ctx, rec)
	if err != nil {
		return model.Department{}, err
	}

	return record.ToDepartment(rec), nil
}
