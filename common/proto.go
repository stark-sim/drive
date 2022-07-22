package common

import "time"

const (
	CookieExpireTime = 3600 * 24 * 7 // cookie过期时间 7天
	TokenExpireTime  = 3600 * 24 * 7 // token过期时间 7天
)

const (
	ContextTokenKey = "token"

	CookieName = "Authorization"
)

const (
	TimeFormatNormal     = "2006-01-02 15:04:05"
	TimeFormatContinuity = "20060102150405"
)

var (
	ZeroTime = time.Time{}
)
