package model

import "rest-api-hitalent/internal/errors"

type DepartmentDeleteMode string

const (
	DepartmentDeleteModeCascade  DepartmentDeleteMode = "cascade"
	DepartmentDeleteModeReassign DepartmentDeleteMode = "reassign"
)

func NewDepartmentDeleteMode(value string) (DepartmentDeleteMode, error) {
	mode := DepartmentDeleteMode(value)
	switch mode {
	case DepartmentDeleteModeCascade, DepartmentDeleteModeReassign:
		return mode, nil
	default:
		return "", errors.ErrDepartmentInvalidDeleteMode
	}
}
