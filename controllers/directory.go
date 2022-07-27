package controllers

import (
	"drive/common"
	"drive/controllers/protos"
	"drive/services"
	"github.com/gin-gonic/gin"
)

func DirectoryGet(c *gin.Context) {
	var req protos.GetDirectoryReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		common.ResponseErrorWithMsg(c, common.CodeInvalidParams, err.Error())
		return
	}

	if req.Id != 0 {
		res, err := services.DirectoryRepository.Get(c, req.Id)
		if err != nil {
			return
		}

		common.ResponseSuccess(c, res)
	} else if req.ParentId != 0 {

	} else {

	}

	common.ResponseSuccess(c, nil)
}

func DirectoryCreate(c *gin.Context) {
	common.ResponseSuccess(c, nil)
}

func DirectoryUpdate(c *gin.Context) {
	common.ResponseSuccess(c, nil)
}

func DirectoryDelete(c *gin.Context) {
	common.ResponseSuccess(c, nil)
}
