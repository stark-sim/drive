package controllers

import (
	"drive/common"
	"drive/protos"
	"drive/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

	userSerializer := common.ModelSerializer{Model: user}
	resp := userSerializer.Serialize()

	common.ResponseSuccess(c, resp)
}

/*
	Login 登录获取 JWT
*/
func Login(c *gin.Context) {
	var req protos.UserLoginReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		common.ResponseErrorWithMsg(c, common.CodeInvalidParams, err.Error())
		return
	}

	token, err := services.UserService.Login(c, req.Phone, req.Password)
	if err != nil {
		common.ResponseErrorWithMsg(c, common.CodeGrpcError, err.Error())
		return
	}

	common.ResponseSuccess(c, token)
}

func UserGet(c *gin.Context) {

	var req protos.UserGetReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		common.ResponseErrorWithMsg(c, common.CodeInvalidParams, err.Error())
		return
	}

	if req.Id != 0 {

		user, err := services.UserRepository.Get(c, req.Id)
		if err != nil {
			logrus.Errorf("get user failed, err: %v", err.Error())
			common.ResponseError(c, common.CodeServerDBError)
			return
		}

		serializer := common.ModelSerializer{Model: user, IsPlural: false}
		resp := serializer.Serialize()

		common.ResponseSuccess(c, resp)
	} else {
		users, _ := services.UserRepository.List(c, nil)

		serializer := common.ModelSerializer{Model: users, IsPlural: true}
		resp := serializer.Serialize()
		common.ResponseSuccess(c, resp)
	}

}
