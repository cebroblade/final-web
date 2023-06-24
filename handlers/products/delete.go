package products

import (
	"errors"
	"net/http"
	"web-final/middleware"
	"web-final/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *ProductHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get id need to be deleted
		pathParamDetail := DetailPathParam{}
		if err := ctx.ShouldBindUri(&pathParamDetail); err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
			return
		}

		// Get the current UserID login now
		userId := middleware.GetCurrenUserId(ctx)

		// Get the product by Id
		product := &models.Product{}
		result := h.Db.Where("id = ? and user_id = ?", pathParamDetail.ID, userId).First(product)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"Error": "Product Not Found"})
			return
		}

		// Execute delete
		h.Db.Delete(&product)
		ctx.JSON(http.StatusOK, gin.H{"message": "Delete success", "productId": product.ID})
	}
}
