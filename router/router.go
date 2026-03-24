package router

import (
	"go_demo/handler"
	"go_demo/middleware"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {

	r := gin.Default()

	v1 := r.Group("/api")
	{
		userRoute := v1.Group("/users")
		userRoute.POST("/register", handler.UserRegister)
		userRoute.POST("/login", handler.UserLogin)

		userRoute.Use(middleware.JWTAuth())
		{
			userRoute.GET("/info", handler.GetUserInfo)
		}
	}

	return r

}
