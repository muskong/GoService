package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muskong/GoWechat/app/user/logic"
)

func Login(c *gin.Context) {
	var userData logic.LoginData
	err := c.ShouldBind(&userData)
	if err != nil {
		c.SecureJSON(http.StatusOK, "传入参数错误")
		return
	}

	data, err := logic.LoginVerify(userData)

	if err != nil {
		c.SecureJSON(http.StatusOK, err.Error)
		return
	}

	c.SecureJSON(http.StatusOK, data)
}

func WeChatLogin(c *gin.Context) {
	var code struct {
		Code string
	}
	err := c.ShouldBindQuery(&code)
	if err != nil {
		c.SecureJSON(http.StatusOK, "传入参数错误")
		return
	}

	data, err := logic.LoginWeChat(code.Code)

	if err != nil {
		c.SecureJSON(http.StatusOK, err.Error)
		return
	}

	c.SecureJSON(http.StatusOK, data)
}
