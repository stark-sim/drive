package services

import (
	"drive/db"
	"drive/services/user"
	"drive/services/user/impl"
)

var (
	UserRepository user.Repository
)

func Init() {
	UserRepository = impl.NewPgImpl(db.NewDBClient())
}
