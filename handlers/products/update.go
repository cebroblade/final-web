package products

import (
	"errors"
	"net/http"
	"web-final/middleware"
	"web-final/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	updateProduct struct {
		Name  string `json : "name" binding:"required"`
		Price uint   `json : "price" binding:"required, number"`
	}
)

func (h *ProductHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get Id from param name.
		pathParamDetail := DetailPathParam{}
		if err := ctx.ShouldBindUri(&pathParamDetail); err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
			return
		}

		// Get the current UserID login now
		userId := middleware.GetCurrenUserId(ctx)

		// Check exist Product
		product := &models.Product{}
		detailPro := h.Db.Where("id = ? and user_id = ?", pathParamDetail.ID, userId).First(product)
		if errors.Is(detailPro.Error, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"Error": "Product Not Found"})
			return
		}

		// Execute update product.
		body := &updateProduct{}
		if err := ctx.ShouldBindJSON(body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}
		product.Name = body.Name
		product.Price = body.Price
		updateProduct := h.Db.Save(&product)
		if updateProduct.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"Error": updateProduct.Error.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Update success",
			"ProductId": product.ID})
	}
}
