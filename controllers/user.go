package controllers

import (
	"drive/common"
	"drive/config"
	"drive/controllers/protos"
	"drive/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/util/json"
	"net/http"
	"net/url"
)

type tempReq struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	Code      string `json:"code"`
}
type LoginRes struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int32  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenId       string `json:"openid"`
	Scope        string `json:"scope"`
	UnionId      string `json:"unionid"`
	ErrCode      int32  `json:"errcode"`
	ErrMsg       string `json:"errmsg"`
}

func Temp(c *gin.Context) {
	var tempReq tempReq
	c.ShouldBindJSON(&tempReq)

	response, err := http.Get(fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", tempReq.AppId, tempReq.AppSecret, tempReq.Code))
	if err != nil {
		logrus.Errorf("wechat login http get error: %s", err)

	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logrus.Errorf("wechat login read resp error: %s", err)

	}
	var res LoginRes
	json.Unmarshal(data, &res)

	logrus.Infof("wechat login res: %+v", res)

}

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

	if config.Conf.Code.Invite != req.InviteCode {
		common.ResponseErrorWithMsg(c, common.CodeInvalidParams, "创建用户需要邀请码")
		return
	}

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
		common.ResponseErrorWithMsg(c, common.CodeServerError, err.Error())
		return
	}

	// 设置cookie
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie(common.CookieName, url.PathEscape(token), common.CookieExpireTime, "/", "", false, true)

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

func UserDelete(c *gin.Context) {
	var req protos.UserDeleteReq
	err := c.ShouldBind(&req)
	if err != nil {
		common.ResponseErrorWithMsg(c, common.CodeInvalidParams, err.Error())
		return
	}
	user, err := services.UserRepository.DeleteOne(c, common.StringToInt64(req.Id))
	if err != nil {
		common.ResponseErrorWithMsg(c, common.CodeServerError, err.Error())
		return
	}
	serializer := common.ModelSerializer{Model: user, IsPlural: false}
	resp := serializer.Serialize()
	common.ResponseSuccess(c, resp)
}
