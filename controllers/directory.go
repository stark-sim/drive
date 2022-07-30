package controllers

import (
	"drive/common"
	"drive/controllers/protos"
	"drive/services"
	"github.com/gin-gonic/gin"
	"sort"
	"strconv"
)

func DirectoryGet(c *gin.Context) {
	var req protos.GetDirectoryReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		common.ResponseErrorWithMsg(c, common.CodeInvalidParams, err.Error())
		return
	}

	if req.ID != 0 {
		// 单目录查询，带上目录下的子目录和文件，按创建时间排序
		res, err := services.DirectoryRepository.Get(c, req.ID)
		if err != nil {
			common.ResponseErrorWithMsg(c, common.CodeServerError, err.Error())
			return
		}

		resp := protos.DirectoryWithContentResp{
			ID:        strconv.FormatInt(res.ID, 10),
			Name:      res.Name,
			ParentID:  strconv.FormatInt(res.ParentID, 10),
			IsPublic:  res.IsPublic,
			Contents:  nil,
			CreatedAt: res.CreatedAt,
			UpdatedAt: res.UpdatedAt,
		}
		for _, object := range res.Edges.Objects {
			resp.Contents = append(resp.Contents, &protos.DirectoryContent{
				ID:        strconv.FormatInt(object.ID, 10),
				Name:      object.Name,
				URL:       object.URL,
				UserID:    strconv.FormatInt(object.UserID, 10),
				IsPublic:  object.IsPublic,
				CreatedAt: object.CreatedAt,
				UpdatedAt: object.UpdatedAt,
			})
		}
		for _, childDirectory := range res.Edges.Children {
			resp.Contents = append(resp.Contents, &protos.DirectoryContent{
				ID:        strconv.FormatInt(childDirectory.ID, 10),
				Name:      childDirectory.Name,
				URL:       "",
				UserID:    strconv.FormatInt(childDirectory.CreatedBy, 10),
				IsPublic:  childDirectory.IsPublic,
				CreatedAt: childDirectory.CreatedAt,
				UpdatedAt: childDirectory.UpdatedAt,
			})
		}

		sort.Slice(resp.Contents, func(i, j int) bool {
			return resp.Contents[i].CreatedAt.Before(resp.Contents[j].CreatedAt)
		})

		common.ResponseSuccess(c, resp)
	} else if req.ParentID != 0 {
		res, err := services.DirectoryRepository.ListByParentID(c, req.ParentID, c.Value("userID").(int64))
		if err != nil {
			common.ResponseErrorWithMsg(c, common.CodeServerError, err.Error())
			return
		}
		data := (&common.ModelSerializer{Model: res, IsPlural: true}).Serialize()
		common.ResponseSuccess(c, data)
	} else {
		res, err := services.DirectoryRepository.ListByParentID(c, 0, c.Value("userID").(int64))
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
	directory, err := services.DirectoryRepository.Create(c, req.Name, common.StringToInt64(req.ParentID), c.Value("userID").(int64), req.IsPublic)
	if err != nil {
		common.ResponseErrorWithMsg(c, common.CodeServerError, err.Error())
		return
	}

	serializer := common.ModelSerializer{Model: directory, IsPlural: false}
	data := serializer.Serialize()

	common.ResponseSuccess(c, data)
}

func DirectoryUpdate(c *gin.Context) {
	var req protos.UpdateDirectoryReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		common.ResponseErrorWithMsg(c, common.CodeInvalidParams, err.Error())
		return
	}

	res, err := services.DirectoryRepository.UpdateOne(c, common.StringToInt64(req.ID), map[string]interface{}{
		"parentID": common.StringToInt64(req.ParentID),
		"name": req.Name,
		"isPublic": req.IsPublic,
	}, c.Value("userID").(int64))
	if err != nil {
		common.ResponseErrorWithMsg(c, common.CodeServerError, err.Error())
		return
	}

	data := (&common.ModelSerializer{
		Model:    res,
		IsPlural: false,
	}).Serialize()

	common.ResponseSuccess(c, data)
}

func DirectoryDelete(c *gin.Context) {
	var req protos.DeleteDirectoryReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		common.ResponseErrorWithMsg(c, common.CodeInvalidParams, err.Error())
		return
	}

	res, err := services.DirectoryRepository.DeleteOne(c, common.StringToInt64(req.ID), c.Value("userID").(int64))
	if err != nil {
		common.ResponseErrorWithMsg(c, common.CodeServerError, err.Error())
		return
	}

	data := (&common.ModelSerializer{Model: res}).Serialize()

	common.ResponseSuccess(c, data)
}
