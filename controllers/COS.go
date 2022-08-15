package controllers

import (
	"drive/common"
	"drive/pkg/minio"
	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		common.ResponseErrorWithMsg(c, common.CodeInvalidParams, "error during parse multipartForm")
		return
	}

	files := form.File["files"]
	minioClient := minio.NewMinioClient()
	results, err := minioClient.UploadFiles(c, files)
	if err != nil {
		common.ResponseErrorWithMsg(c, common.CodeServerError, err.Error())
		return
	}

	common.ResponseSuccess(c, results)
}


func ListFiles(c *gin.Context) {
	minioClient := minio.NewMinioClient()
	files := minioClient.ListFiles(c)

	common.ResponseSuccess(c, files)
}