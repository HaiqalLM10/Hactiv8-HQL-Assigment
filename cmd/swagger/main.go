package main

import (
	_ "sesi-8/docs"
	"sesi-8/domain/todo"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Todo API
// @version v1.0
// @description Example API Swagger
// @contact.name Reyhan Jovie
// @contact.email reyhan@gmail.com
// @host localhost:4444
// @Basepath /
func main() {
	router := gin.New()

	v1 := router.Group("v1")
	todo.InitRouter(v1)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":4444")
}
