package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Identity string `gorm:"column:identity; type:varchar(50); not null;unique;" json:"identity"`
	Name     string `gorm:"column:name; type:varchar(50); not null;unique;" json:"name"`
	Desc     string `gorm:"column:desc; type:varchar(50);" json:"desc"`
}

func (table Role) TableName() string {
	return "athana_role"
}
