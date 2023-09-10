package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Identity string `gorm:"column:identity; type:varchar(50); not null;unique;" json:"identity"`
	Username string `gorm:"column:username; type:varchar(50); not null;unique;" json:"username"`
	Password string `gorm:"column:password; type:varchar(50); not null;" json:"password"`
	Email    string `gorm:"column:email; type:varchar(50);" json:"email"`
}

func (table User) TableName() string {
	return "athana_user"
}
