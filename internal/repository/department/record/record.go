package record

import (
	"time"

	"gorm.io/gorm/schema"
)

var _ schema.Tabler = (*DepartmentRecord)(nil)

type DepartmentRecord struct {
	Id        uint `gorm:"primarykey"`
	Name      string
	ParentId  *uint
	Children  []DepartmentRecord `gorm:"foreignKey:ParentId"`
	CreatedAt time.Time
}

func (DepartmentRecord) TableName() string {
	return "department"
}
