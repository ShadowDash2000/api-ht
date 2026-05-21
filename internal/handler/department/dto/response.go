package dto

import "time"

type CreateDepartmentResponse struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	ParentId  uint      `json:"parent_id"`
	CreatedAt time.Time `json:"created_at"`
}

type Department struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	ParentId  uint      `json:"parent_id"`
	CreatedAt time.Time `json:"created_at"`
}

type Employee struct {
	Id           uint      `json:"id"`
	DepartmentId uint      `json:"department_id"`
	FullName     string    `json:"full_name"`
	Position     string    `json:"position"`
	HiredAt      time.Time `json:"hired_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type GetDepartmentResponse struct {
	Department Department   `json:"department"`
	Employees  []Employee   `json:"employees"`
	Children   []Department `json:"children"`
}
