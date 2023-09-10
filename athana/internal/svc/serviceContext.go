package svc

import (
	"github.com/pjimming/zeus/athana/internal/config"
	"github.com/pjimming/zeus/athana/internal/middleware"
	"github.com/pjimming/zeus/athana/models"
	"github.com/pjimming/zeus/utils/errorx"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ServiceContext struct {
	// config
	Config config.Config
	// Database
	DB *gorm.DB
	// middleware
	RBACAuth rest.Middleware
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
		// config
		Config: c,
		// database
		DB: db,
		// middleware
		RBACAuth: middleware.NewRBACAuthMiddleware().Handle,
	}
}
