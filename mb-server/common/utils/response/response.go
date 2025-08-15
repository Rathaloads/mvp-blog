package response

import "github.com/gin-gonic/gin"

func ReturnJson(c *gin.Context, code int, message string, data any) {
	c.JSON(200, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	})
}

func ReturnSuccess(c *gin.Context) {
	ReturnJson(c, 0, "success", nil)
}

func ReturnError(c *gin.Context, code int, message string) {
	ReturnJson(c, code, message, nil)
}

func ReturnData(c *gin.Context, data any) {
	ReturnJson(c, 0, "success", data)
}
