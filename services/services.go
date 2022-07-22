package services

import (
	"drive/db"
	"drive/services/user_service"
	"drive/services/user_service/impl"
)

var (
	UserRepository user_service.Repository
	UserService    user_service.UserService
)

func Init() {
	UserRepository = impl.NewPgImpl(db.NewDBClient())
	UserService = impl.NewUserService(UserRepository)
}
