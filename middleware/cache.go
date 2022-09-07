package middleware

import (
	"github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
	"time"
)


/*
	CacheMiddleware 缓存中间件（非官方）
 */
func CacheMiddleware() gin.HandlerFunc {
	memoryStore := persist.NewMemoryStore(10 * time.Minute)
	return cache.CacheByRequestURI(memoryStore, 3 * time.Second)
}
