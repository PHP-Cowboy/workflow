package global

import (
	"gorm.io/gorm"
	"workflow/config"
)

var (
	DB           *gorm.DB
	ServerConfig = &config.ServerConfig{}
)
