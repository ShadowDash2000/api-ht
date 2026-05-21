package department

import (
	"context"
	internalerrors "rest-api-hitalent/internal/errors"
	"rest-api-hitalent/internal/model"
	"testing"

	"github.com/stretchr/testify/require"
)

type mockDepartmentRepository struct {
	store  map[uint]model.Department
	nextId uint
}

func newMockDepartmentRepository() *mockDepartmentRepository {
	return &mockDepartmentRepository{
		store:  make(map[uint]model.Department),
		nextId: 1,
	}
}

func (r *mockDepartmentRepository) Create(_ context.Context, department model.Department) (model.Department, error) {
	department.Id = r.nextId
	r.store[r.nextId] = department
	r.nextId++
	return department, nil
}

func (r *mockDepartmentRepository) GetAllChildrenWithDepth(_ context.Context, id uint, depth uint) ([]model.Department, error) {
	rootDepartment, ok := r.store[id]
	if !ok {
		return nil, internalerrors.ErrDepartmentNotFound
	}

	// If depth is 0, it means we must search all children without depth limit.
	searchAllChildren := depth == 0

	res := []model.Department{rootDepartment}
	idsForSearch := map[uint]struct{}{id: {}}
	for (depth > 1 && len(idsForSearch) > 0) || (searchAllChildren && len(idsForSearch) > 0) {
		nextIdsForSearch := make(map[uint]struct{})

		// Not efficient way to search children, but it's ok for testing.
		// Here we just iterate through the entire map every depth.
		for _, department := range r.store {
			_, ok = idsForSearch[department.ParentId]
			if ok {
				nextIdsForSearch[department.Id] = struct{}{}
				res = append(res, department)
			}
		}
		idsForSearch = nextIdsForSearch
		depth--
	}
	return res, nil
}

func (r *mockDepartmentRepository) Delete(_ context.Context, ids []uint) error {
	for _, id := range ids {
		delete(r.store, id)
	}
	return nil
}

func (r *mockDepartmentRepository) Update(_ context.Context, department model.Department) (model.Department, error) {
	storeDepartment, ok := r.store[department.Id]
	if !ok {
		return model.Department{}, internalerrors.ErrDepartmentNotFound
	}
	department.Id = storeDepartment.Id
	r.store[department.Id] = department
	return department, nil
}

func (r *mockDepartmentRepository) Get(_ context.Context, id uint) (model.Department, error) {
	storeDepartment, ok := r.store[id]
	if !ok {
		return model.Department{}, internalerrors.ErrDepartmentNotFound
	}
	return storeDepartment, nil
}

type mockEmployeeRepository struct {
	store  map[uint]model.Employee
	nextId uint
}

func newMockEmployeeRepository() *mockEmployeeRepository {
	return &mockEmployeeRepository{
		store:  make(map[uint]model.Employee),
		nextId: 1,
	}
}

func (r *mockEmployeeRepository) GetAllByDepartmentIds(_ context.Context, departmentIds []uint) ([]model.Employee, error) {
	var res []model.Employee
	for _, id := range departmentIds {
		for _, employee := range r.store {
			if employee.DepartmentId == id {
				res = append(res, employee)
			}
		}
	}
	return res, nil
}

func (r *mockEmployeeRepository) DeleteByDepartmentIds(_ context.Context, departmentIds []uint) error {
	ids := make(map[uint]struct{})
	for _, id := range departmentIds {
		ids[id] = struct{}{}
	}
	for id, employee := range r.store {
		if _, ok := ids[employee.DepartmentId]; ok {
			delete(r.store, id)
		}
	}
	return nil
}

func (r *mockEmployeeRepository) Reassign(_ context.Context, oldId, newId uint) error {
	for _, employee := range r.store {
		if employee.DepartmentId != oldId {
			continue
		}
		employee.DepartmentId = newId
	}
	return nil
}

func Test_CreateDepartment_Success(t *testing.T) {
	departmentRepository := newMockDepartmentRepository()
	employeeRepository := newMockEmployeeRepository()
	service := NewService(departmentRepository, employeeRepository)

	department, err := service.Create(context.Background(), "Department", 1)

	require.NoError(t, err)
	require.Greater(t, department.Id, uint(0))
}

func Test_GetDepartment_Success(t *testing.T) {
	departmentRepository := newMockDepartmentRepository()
	employeeRepository := newMockEmployeeRepository()
	service := NewService(departmentRepository, employeeRepository)

	rootDepartment, _ := service.Create(context.Background(), "Department root", 0)
	subRootDepartment, _ := service.Create(context.Background(), "Department sub root", rootDepartment.Id)

	_, _ = service.Create(context.Background(), "Main Root Department 1", rootDepartment.Id)
	_, _ = service.Create(context.Background(), "Main Root Department 2", rootDepartment.Id)

	_, _ = service.Create(context.Background(), "Sub Root Department 1", subRootDepartment.Id)
	_, _ = service.Create(context.Background(), "Sub Root Department 2", subRootDepartment.Id)

	detail, err := service.Get(context.Background(), rootDepartment.Id, 3, true)
	require.NoError(t, err)
	require.Len(t, detail.Children, 5)
}
