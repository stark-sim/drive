package controllers

import (
	"drive/common"
	"drive/controllers/protos"
	"drive/services"
	"github.com/gin-gonic/gin"
)

func ObjectCreate(c *gin.Context) {
	var req protos.ObjectCreateReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		common.ResponseErrorWithMsg(c, common.CodeInvalidParams, err.Error())
		return
	}
	userId := c.Value("userId")
	object, err := services.ObjectRepository.Create(c, req.Name, req.URL, common.StringToInt64(req.DirectoryID), userId.(int64), req.IsPublic)
	if err != nil {
		common.ResponseErrorWithMsg(c, common.CodeServerError, err.Error())
		return
	}
	data := (&common.ModelSerializer{
		Model:    object,
		IsPlural: false,
	}).Serialize()
	common.ResponseSuccess(c, data)
}

func ObjectDelete(c *gin.Context) {
	var req protos.ObjectDeleteReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		common.ResponseErrorWithMsg(c, common.CodeInvalidParams, err.Error())
		return
	}
	userId := c.Value("userId")
	object, err := services.ObjectRepository.DeleteOne(c, common.StringToInt64(req.ID), userId.(int64))
	if err != nil {
		common.ResponseErrorWithMsg(c, common.CodeServerError, err.Error())
		return
	}
	data := (&common.ModelSerializer{
		Model:    object,
		IsPlural: false,
	}).Serialize()
	common.ResponseSuccess(c, data)
}
