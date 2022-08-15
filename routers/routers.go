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
	v1.POST("/cos", controllers.UploadFile)
	v1.GET("/cos", controllers.ListFiles)

	usersR := v1.Group("/users")

	usersR.GET("", controllers.UserGet)
	usersR.DELETE("", controllers.UserDelete)

	directoriesR := v1.Group("/directories")

	directoriesR.GET("", controllers.DirectoryGet)
	directoriesR.POST("", controllers.DirectoryCreate)
	directoriesR.PUT("", controllers.DirectoryUpdate)
	directoriesR.DELETE("", controllers.DirectoryDelete)

	objectsR := v1.Group("/objects")

	objectsR.POST("", controllers.ObjectCreate)
	objectsR.DELETE("", controllers.ObjectDelete)
	objectsR.PUT("", controllers.ObjectUpdate)

	return r
}
