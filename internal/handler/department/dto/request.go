package dto

type CreateDepartmentRequest struct {
	Name     string `json:"name"`
	ParentId uint   `json:"parent_id"`
}

type UpdateDepartmentRequest struct {
	Name     *string `json:"name"`
	ParentId *uint   `json:"parent_id"`
}
