package routers

import (
	"web-final/handlers/users"

	"github.com/gin-gonic/gin"
)

func (r *Router) AddUserRouter(apiRouter *gin.RouterGroup) {
	userRouter := apiRouter.Group("users")
	handler := users.UserHandler{Db: r.Db}

	userRouter.POST("sign-up", handler.RegisterNewUser())
	userRouter.GET("log-in", handler.Login())
}
