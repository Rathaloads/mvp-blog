package middleware

import (
	"mb-server/common/db"
	"mb-server/common/utils/response"
	"mb-server/common/utils/token"

	"github.com/gin-gonic/gin"
)

// 权限中间件
func AuthMiddleware(c *gin.Context) {
	tk := c.GetHeader("Authorization")
	if tk == "" {
		response.ReturnError(c, 101, "Illegal request")
		c.Abort()
		return
	}
	_, err := db.RedisGetEx("token", tk)
	if err != nil {
		response.ReturnError(c, 101, "Token expired")
		c.Abort()
	}
	claim, err := token.ValidateToken(tk)
	if err != nil {
		response.ReturnError(c, 100, err.Error())
		c.Abort()
		return
	}
	c.Set("Email", claim.Email)
	c.Next()
}
