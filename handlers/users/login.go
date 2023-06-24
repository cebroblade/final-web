package users

import (
	"errors"
	"net/http"
	"web-final/models"

	"web-final/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type loginBody struct {
	Username string `json : "username" binding:"required"`
	Password string `json : "password" binding:"required"`
}

func (h *UserHandler) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Check the body request
		body := &registerNewUser{}
		if err := ctx.ShouldBindJSON(body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		// Get the infomation of username
		user := &models.User{Username: body.Username, Password: body.Password}
		result := h.Db.First(user, "username = ?", body.Username)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Username"})
			return
		}

		// Check the information of password
		if err := models.CheckPassword(user.Password, body.Password); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Password"})
			return
		}

		// generate the token
		tokenString, err := middleware.GenerateJWT(user.ID, user.Username)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Fail to generate token"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"tokenString": tokenString})
	}
}
