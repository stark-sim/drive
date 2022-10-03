package routers

import "C"
import (
	"drive/controllers"
	"drive/middleware"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()

	driveAPI := r.Group("/drive/api")
	driveAPI.Use(middleware.CorsMiddleware())

	cacheMiddleware := middleware.CacheMiddleware()

	driveAPI.POST("/temp", controllers.Temp)
	driveAPI.POST("/register", controllers.UserAdd)
	driveAPI.GET("/login", controllers.Login)

	v1 := driveAPI.Group("/v1")
	// v1 接口上认证
	v1.Use(middleware.AuthMiddleware())
	v1.POST("/cos", controllers.UploadFile)
	v1.GET("/cos", controllers.ListFiles)

	usersR := v1.Group("/users")

	// 部分接口上缓存
	usersR.GET("", controllers.UserGet, cacheMiddleware)
	usersR.DELETE("", controllers.UserDelete)

	directoriesR := v1.Group("/directories")

	directoriesR.GET("", controllers.DirectoryGet, cacheMiddleware)
	directoriesR.POST("", controllers.DirectoryCreate)
	directoriesR.PUT("", controllers.DirectoryUpdate)
	directoriesR.DELETE("", controllers.DirectoryDelete)

	objectsR := v1.Group("/objects")

	objectsR.POST("", controllers.ObjectCreate)
	objectsR.DELETE("", controllers.ObjectDelete)
	objectsR.PUT("", controllers.ObjectUpdate)

	return r
}
