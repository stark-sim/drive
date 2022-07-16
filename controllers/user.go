package controllers

import (
	"drive/common"
	"drive/protos"
	"drive/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strconv"
)

/*
	UserAdd 添加用户
*/
func UserAdd(c *gin.Context) {
	// 获取参数
	var req protos.UserAddReq
	if err := c.ShouldBindJSON(&req); err != nil {
		common.ResponseError(c, common.CodeInvalidParams)
		return
	}
	logrus.Debugf("req: %+v", req)

	// 获取 token
	//token, err := auth.GetToken(c)
	//if err != nil {
	//	common.ResponseErrorWithMsg(c, common.CodeInvalidToken, err.Error())
	//	return
	//}

	// 判断是否已经注册
	//user, err := service.ExplorerUserRepository.Get(registerResponse.UserId)
	//if err != nil {
	//	logrus.Errorf("query user failed, err: %s", err.Error())
	//	common.ResponseError(c, common.CodeServerDBError)
	//	return
	//}

	//if user.UserId == 0 {
	//	// 未注册，添加用户
	//	if err = service.ExplorerUserRepository.Create(nil, &models.ExplorerUser{
	//		UserId: registerResponse.UserId,
	//		Status: req.Status,
	//		Role:   req.Role,
	//	}); err != nil {
	//		logrus.Errorf("create user failed, err: %s", err.Error())
	//		common.ResponseError(c, common.CodeServerDBError)
	//		return
	//	}
	//}

	user, err := services.UserRepository.Create(c, req.Name, req.Password, req.Phone)
	if err != nil {
		logrus.Errorf("create user failed, err: %s", err.Error())
		common.ResponseError(c, common.CodeServerDBError)
		return
	}

	common.ResponseSuccess(c, &user)
}

func UserGet(c *gin.Context) {

	id := c.Query("id")
	idStr, _ := strconv.ParseInt(id, 10, 64)

	user, err := services.UserRepository.Get(c, idStr)
	if err != nil {
		logrus.Errorf("get user failed, err: %v", err.Error())
		common.ResponseError(c, common.CodeServerDBError)
		return
	}

	common.ResponseSuccess(c, user)
}
