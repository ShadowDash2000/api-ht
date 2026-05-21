package dto

import "time"

type CreateEmployeeResponse struct {
	Id           uint      `json:"id"`
	DepartmentId uint      `json:"department_id"`
	FullName     string    `json:"full_name"`
	Position     string    `json:"position"`
	HiredAt      time.Time `json:"hired_at"`
	CreatedAt    time.Time `json:"created_at"`
}
