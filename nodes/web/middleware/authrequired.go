package middleware

import (
	"github.com/gin-gonic/gin"
	superGin "github.com/po2656233/superplace/components/gin"
	"net/http"
	"strings"
	"superman/internal/token"
)

var skipRoutes = map[string]bool{
	"/":            true,
	"/login":       true,
	"/register":    true,
	"/list/pid":    true,
	"/list/server": true,
}

// AuthRequired gin jwt 认证中间件
func AuthRequired() superGin.GinHandlerFunc {
	return func(ctx *superGin.Context) {
		uri := ctx.Request.RequestURI
		_, ok := skipRoutes[uri]
		if ok || strings.Contains(uri, "/list/server") {
			return
		}
		tokenString := strings.TrimPrefix(ctx.GetHeader("Authorization"), "Bearer ")
		retCode, msg := token.VerifyToken(ctx, tokenString)
		if retCode != http.StatusOK {
			ctx.AbortWithStatusJSON(retCode, gin.H{"code": -1, "message": msg})
			return
		}
		ctx.Next()
	}
}
