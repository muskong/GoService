package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muskong/GoPkg/zaplog"
	"github.com/muskong/GoWechat/app/order/logic"
	"github.com/spf13/cast"
)

func StoryLike(c *gin.Context) {
	var likeOrder logic.LikeOrder
	err := c.ShouldBind(&likeOrder)
	if err != nil {
		zaplog.Sugar.Error(err)
		c.SecureJSON(http.StatusOK, "传入参数错误")
		return
	}

	userId, ok := c.Get("userId")
	if !ok {
		c.SecureJSON(http.StatusOK, "未登录")
	}

	likeOrder.UserId = cast.ToInt(userId)

	data, err := logic.StoryLike(likeOrder)

	if err != nil {
		c.SecureJSON(http.StatusOK, err.Error())
		return
	}

	c.SecureJSON(http.StatusOK, data)
}

func UserStories(c *gin.Context) {
	var page logic.Page
	err := c.ShouldBindJSON(&page)

	userId, ok := c.Get("userId")
	if !ok {
		c.SecureJSON(http.StatusOK, "未登录")
	}

	data, err := logic.UserStories(cast.ToInt(userId), &page)

	if err != nil {
		c.SecureJSON(http.StatusOK, err.Error())
		return
	}

	c.SecureJSON(http.StatusOK, data)
}
