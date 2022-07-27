package middleware

import (
	"drive/common"
	"drive/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/url"
)

const (
	CookieContext = "bearer %s"
)

/*
	AuthMiddleware 认证中间件
*/
func AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 获取 cookie
		cookie, err := c.Cookie(common.CookieName)
		cookie, err = url.PathUnescape(cookie)
		logrus.Printf("=======> cookie: %s", cookie)
		if err != nil {
			common.ResponseErrorWithMsg(c, common.CodeInvalidToken, "Cookie is not set")
			c.Abort()
			return
		}

		// 验证 JWT， 顺便把 userId 存于上下文
		customClaims, err := tools.ParseToken(cookie)
		if err != nil {
			common.ResponseErrorWithMsg(c, common.CodeInvalidToken, err.Error())
			c.Abort()
			return
		}
		c.Set("userId", customClaims.UserId)

		// 将 token 保存到上下文中，便于发送 GRPC 请求，由于 GRPC 请求 METADATA key 全小写，所以 Authorization 换成 token
		c.Set(common.ContextTokenKey, fmt.Sprintf(CookieContext, cookie))

		c.Next()
	}
}
