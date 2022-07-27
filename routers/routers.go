package routers

import (
	"drive/controllers"
	"drive/middleware"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()

	r.POST("/temp", controllers.Temp)

	r.POST("/register", controllers.UserAdd)
	r.GET("/login", controllers.Login)

	v1 := r.Group("/v1")

	v1.Use(middleware.AuthMiddleware())

	v1.GET("/users", controllers.UserGet)
	v1.DELETE("/users", controllers.UserDelete)

	return r
}
