package services

import (
	"drive/db"
	"drive/services/directory_service"
	directoryImpl "drive/services/directory_service/impl"
	"drive/services/object_service"
	objectImpl "drive/services/object_service/impl"
	"drive/services/user_service"
	userImpl "drive/services/user_service/impl"
)

var (
	UserRepository      user_service.Repository
	UserService         user_service.UserService
	DirectoryRepository directory_service.Repository

	ObjectRepository object_service.Repository
)

func Init() {
	dbClient1 := db.NewDBClient()
	UserRepository = userImpl.NewPgImpl(dbClient1)
	UserService = userImpl.NewUserService(UserRepository)

	dbClient2 := db.NewDBClient()
	DirectoryRepository = directoryImpl.NewPgImpl(dbClient2)

	dbClient3 := db.NewDBClient()
	ObjectRepository = objectImpl.NewPgImpl(dbClient3)
}
