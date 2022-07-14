package middleware

import (
	"drive/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
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
		log.Printf("=======> cookie: %s", cookie)
		if err != nil {
			common.ResponseErrorWithMsg(c, common.CodeInvalidToken, "Cookie is not set")
			c.Abort()
			return
		}

		// 将 token 保存到上下文中
		c.Set(common.ContextTokenKey, fmt.Sprintf(CookieContext, cookie))

		c.Next()
	}
}
