package products

import (
	"net/http"
	"web-final/middleware"
	"web-final/models"

	"github.com/gin-gonic/gin"
)

func (h *ProductHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products := &[]models.Product{}

		// Get the current UserID login now
		userId := middleware.GetCurrenUserId(ctx)

		// Get all the list that is belong to current user
		result := h.Db.Find(products, "user_id = ?", userId)
		if result.Error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"Error": result.Error.Error()})
		}
		ctx.JSON(http.StatusOK, gin.H{"data": products})
	}
}
