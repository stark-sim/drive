package routers

import (
	"drive/controllers"
	"drive/middleware"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()

	r.POST("/register", controllers.UserAdd)
	//r.GET("/login")

	v1 := r.Group("/v1")

	v1.Use(middleware.AuthMiddleware())

	return r
}
