package products

import (
	"errors"
	"net/http"
	"web-final/middleware"
	"web-final/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DetailPathParam struct {
	ID uint `uri:"id" binding:"required"`
}

func (h *ProductHandler) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Get parameter form uri
		pathParamDetail := DetailPathParam{}
		if err := ctx.ShouldBindUri(&pathParamDetail); err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
			return
		}
		// Get current user logon
		userId := middleware.GetCurrenUserId(ctx)

		product := &models.Product{}
		result := h.Db.Where("id = ? and user_id = ?", pathParamDetail.ID, userId).First(product)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"Error": "Product Not Found"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": product})
	}
}
