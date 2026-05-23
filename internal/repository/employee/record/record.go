package record

import (
	"time"

	"gorm.io/gorm/schema"
)

var _ schema.Tabler = (*EmployeeRecord)(nil)

type EmployeeRecord struct {
	Id           uint `gorm:"primarykey"`
	DepartmentId uint
	FullName     string
	Position     string
	HiredAt      time.Time
	CreatedAt    time.Time
}

func (EmployeeRecord) TableName() string {
	return "employee"
}
