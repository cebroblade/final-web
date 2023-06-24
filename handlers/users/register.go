package users

import (
	"net/http"
	"web-final/models"

	"github.com/gin-gonic/gin"
)

type registerNewUser struct {
	Username string `json : "username" binding:"required"`
	Password string `json : "password" binding:"required"`
}

func (h *UserHandler) RegisterNewUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Check the body request
		body := &registerNewUser{}
		if err := ctx.ShouldBindJSON(body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		// Execute create User
		hashPassword, err := models.HashPassword(body.Password)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}
		user := &models.User{Username: body.Username, Password: hashPassword}
		result := h.Db.Create(user)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"Error": result.Error.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, gin.H{"Id": user.ID})
	}
}
