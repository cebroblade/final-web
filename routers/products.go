package routers

import (
	"web-final/handlers/products"
	"web-final/middleware"

	"github.com/gin-gonic/gin"
)

func (r *Router) AddProductRouter(apiRouter *gin.RouterGroup) {
	productRouter := apiRouter.Group("products").Use(middleware.Authorized())
	handler := products.ProductHandler{Db: r.Db}

	productRouter.GET("", handler.GetAll())
	productRouter.GET("/:id", handler.GetById())
	productRouter.POST("", handler.Create())
	productRouter.PUT("/:id", handler.Update())
	productRouter.DELETE("/:id", handler.Delete())
}
