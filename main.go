package main

import (
	"github.com/dorman99/go_gin_mysql/config"
	"github.com/dorman99/go_gin_mysql/route"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()
)

func main() {
	defer config.CloseDBConnection(db)
	r := gin.Default()
	v1Routes := route.V1Routes(r, db)
	r.Use(v1Routes.Handlers...)
	r.Run()
}
