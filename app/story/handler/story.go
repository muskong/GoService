package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muskong/GoWechat/app/story/logic"
)

func Stories(c *gin.Context) {
	var page logic.Page
	err := c.ShouldBindJSON(&page)

	data, err := logic.Stories(&page)

	if err != nil {
		c.SecureJSON(http.StatusOK, err.Error())
		return
	}

	c.SecureJSON(http.StatusOK, data)
}

func Story(c *gin.Context) {
	var q struct{ StoryId int }
	err := c.ShouldBindQuery(&q)
	if err != nil || q.StoryId == 0 {
		c.SecureJSON(http.StatusOK, "传入参数错误")
		return
	}

	data, err := logic.Story(q.StoryId)

	if err != nil {
		c.SecureJSON(http.StatusOK, err.Error())
		return
	}

	c.SecureJSON(http.StatusOK, data)
}
