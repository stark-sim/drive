package routers

import (
	"drive/controllers"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()

	r.POST("/register", controllers.UserAdd)
	//r.GET("/login")

	v1 := r.Group("/v1")

	//v1.Use(middleware.AuthMiddleware())

	v1.GET("/users", controllers.UserGet)

	return r
}
