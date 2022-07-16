package services

import (
	"drive/db"
	"drive/services/user_service"
	"drive/services/user_service/impl"
)

var (
	UserRepository user_service.Repository
)

func Init() {
	UserRepository = impl.NewPgImpl(db.NewDBClient())
}
