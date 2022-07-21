package app

import (
	"github.com/gin-gonic/gin"
	"github.com/muskong/GoCore/middlewares"
	"github.com/muskong/GoPkg/jwt"
	story "github.com/muskong/GoWechat/app/story/handler"
	user "github.com/muskong/GoWechat/app/user/handler"
)

func GinRouter() *gin.Engine {
	tokenName := jwt.Jwt.GetTokenName()
	notAuth := map[string]bool{
		"/user/login":   true,
		"/user/wechat":  true,
		"/story/list":   true,
		"/story/detail": true,
	}

	router := gin.Default()
	router.Use(middlewares.GinCORS())
	router.Use(middlewares.GinUserMiddleware(tokenName, notAuth))

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
