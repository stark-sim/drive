package main

import (
	"drive/config"
	"drive/db"
	"drive/logger"
	"drive/tools"
)

func main() {
	err := config.Init()
	if err != nil {
		return
	}
	err = tools.Init()
	if err != nil {
		return
	}
	logger.InitLog()
	db.Init()
}
