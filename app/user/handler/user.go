package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muskong/GoWechat/app/user/logic"
)

func UserInfo(c *gin.Context) {

	userId := c.GetInt64("userId")

	data, err := logic.UserInfo(userId)

	if err != nil {
		c.SecureJSON(http.StatusOK, err.Error())
		return
	}

	c.SecureJSON(http.StatusOK, data)
}
