package middleware

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	UserID   uint   `json:"userID"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID uint, username string) (string, error) {
	// Declare the expiration time of the token
	// here, we have kept it as 60 minutes
	expirationTime := time.Now().Add(60 * time.Minute)

	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: username,
		UserID:   userID,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		err = errors.New("Can't parse the claim")
		return
	}

	if claims.ExpiresAt.Compare(time.Now().Local()) == -1 {
		err = errors.New("Expired token")
		return
	}
	return
}

func GetCurrenUserId(ctx *gin.Context) uint {
	signedToken := ctx.GetHeader("Authorization")
	token, _ := jwt.ParseWithClaims(
		signedToken,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})

	claims, _ := token.Claims.(*Claims)
	return claims.UserID
}

func Authorized() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		err := ValidateToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"Error": err.Error()})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
