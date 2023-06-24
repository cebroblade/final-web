package products

import (
	"net/http"
	"web-final/middleware"
	"web-final/models"

	"github.com/gin-gonic/gin"
)

type CreateProduct struct {
	Name  string `json : "name" binding:"required"`
	Price uint   `json : "price" binding:"required, number"`
}

func (h *ProductHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Get information from the body
		body := &CreateProduct{}
		if err := ctx.ShouldBindJSON(body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		// Get the current UserID login now
		userId := middleware.GetCurrenUserId(ctx)

		newProduct := &models.Product{Name: body.Name, Price: body.Price, UserID: userId}
		result := h.Db.Create(newProduct)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"Error": result.Error.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message":   "Create new product success",
			"productId": newProduct.ID})
	}
}
