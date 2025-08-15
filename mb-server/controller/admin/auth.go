package admin

import (
	"mb-server/common/utils/response"

	"github.com/gin-gonic/gin"
)

func AdminLogin(c *gin.Context) {
	response.ReturnSuccess(c)
}
