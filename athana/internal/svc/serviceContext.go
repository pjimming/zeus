package svc

import (
	"github.com/pjimming/zeus/athana/internal/config"
	"github.com/pjimming/zeus/athana/models"
	"github.com/pjimming/zeus/utils/errorx"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.MySQL.DSN), &gorm.Config{
		Logger: logger.Default,
	})
	if err != nil {
		err = errorx.ErrorDB(err)
		logx.Error(err)
	}

	if err = db.AutoMigrate(
		&models.User{},
		&models.Role{},
	); err != nil {
		err = errorx.ErrorDB(err)
		logx.Error(err)
	}

	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
