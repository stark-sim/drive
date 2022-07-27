package services

import (
	"drive/db"
	"drive/services/directory_service"
	directoryImpl "drive/services/directory_service/impl"
	"drive/services/user_service"
	userImpl "drive/services/user_service/impl"
)

var (
	UserRepository      user_service.Repository
	UserService         user_service.UserService
	DirectoryRepository directory_service.Repository
)

func Init() {
	UserRepository = userImpl.NewPgImpl(db.NewDBClient())
	UserService = userImpl.NewUserService(UserRepository)

	DirectoryRepository = directoryImpl.NewPgImpl(db.NewDBClient())
}
