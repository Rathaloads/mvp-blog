package admin

import (
	"mb-server/common/logger"
	"mb-server/common/utils/response"
	"mb-server/dto"
	"mb-server/services"

	"github.com/gin-gonic/gin"
)

func AdminLogin(c *gin.Context) {
	var loginParams dto.LoginRequest
	if err := c.ShouldBind(&loginParams); err != nil {
		logger.Errorf("get params error: %v", err)
		response.ReturnError(c, 100, "params error!")
		return
	}
	token, err := services.AdminUserLogin(loginParams.Email, loginParams.Password)
	if err != nil {
		response.ReturnError(c, 100, err.Error())
		return
	}
	data := &dto.LoginResponse{
		Token: token,
	}
	response.ReturnData(c, data)
}

func AdminRegister(c *gin.Context) {
	var params dto.LoginRequest
	if err := c.ShouldBind(&params); err != nil {
		response.ReturnError(c, 100, "params error!")
		return
	}
	err := services.AdminUserRegister(params.Email, params.Password)
	if err != nil {
		response.ReturnError(c, 100, err.Error())
		return
	}
	response.ReturnSuccess(c)
}
