package app

import (
	"github.com/gin-gonic/gin"
	"github.com/muskong/GoCore/middlewares"
	story "github.com/muskong/GoWechat/app/story/handler"
	user "github.com/muskong/GoWechat/app/user/handler"
)

func GinRouter() *gin.Engine {

	router := gin.Default()
	router.Use(middlewares.GinCORS())
	router.Use(middlewares.GinUserMiddleware())

	apiStory := router.Group("/story")
	{
		// story
		apiStory.GET("/list", story.Stories)
		apiStory.GET("/detail", story.Story)
		apiStory.POST("/like", story.StoryLike)

		apiStory.GET("/user", story.UserStories)
	}
	apiUser := router.Group("/user")
	{
		apiUser.POST("/login", user.Login)
		apiUser.GET("/wechat", user.WeChatLogin)
		apiUser.GET("/info", user.UserInfo)
	}

	return router
}
