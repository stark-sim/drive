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
			common.ResponseErrorWithMsg(c, common.CodeServerError, err.Error())
			return
		}

		common.ResponseSuccess(c, res)
	} else if req.ParentId != 0 {
		res, err := services.DirectoryRepository.ListByParentId(c, req.ParentId, c.Value("userId").(int64))
		if err != nil {
			common.ResponseErrorWithMsg(c, common.CodeServerError, err.Error())
			return
		}
		data := (&common.ModelSerializer{Model: res, IsPlural: true}).Serialize()
		common.ResponseSuccess(c, data)
	} else {
		res, err := services.DirectoryRepository.ListByParentId(c, 0, c.Value("userId").(int64))
		if err != nil {
			common.ResponseErrorWithMsg(c, common.CodeServerError, err.Error())
			return
		}
		data := (&common.ModelSerializer{Model: res, IsPlural: true}).Serialize()
		common.ResponseSuccess(c, data)
	}

}

func DirectoryCreate(c *gin.Context) {
	var req protos.CreateDirectoryReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		common.ResponseErrorWithMsg(c, common.CodeInvalidParams, err.Error())
		return
	}
	directory, err := services.DirectoryRepository.Create(c, req.Name, common.StringToInt64(req.ParentId), c.Value("userId").(int64), req.IsPublic)
	if err != nil {
		common.ResponseErrorWithMsg(c, common.CodeServerError, err.Error())
		return
	}

	serializer := common.ModelSerializer{Model: directory, IsPlural: false}
	data := serializer.Serialize()

	common.ResponseSuccess(c, data)
}

func DirectoryUpdate(c *gin.Context) {
	common.ResponseSuccess(c, nil)
}

func DirectoryDelete(c *gin.Context) {
	common.ResponseSuccess(c, nil)
}
