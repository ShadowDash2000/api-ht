package errors

import "errors"

var (
	ErrDepartmentInvalidName = errors.New("department name is invalid")

	ErrEmployeeInvalidFullName     = errors.New("employee's full name is invalid")
	ErrEmployeeInvalidPosition     = errors.New("employee's position is invalid")
	ErrEmployeeInvalidDepartmentId = errors.New("employee's department id is invalid")

	ErrDepartmentInvalidDeleteMode = errors.New("department delete mode is invalid")

	ErrDepartmentSelfParent = errors.New("department cannot be its own parent")
	ErrDepartmentCycle      = errors.New("cannot move department into its own subtree")
	ErrDepartmentNotFound   = errors.New("department not found")
)
