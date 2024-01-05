package middleware

import (
	"errors"
	"log"
	"net/http"

	"github.com/agusbasari29/xjx-code/helper"
	"github.com/agusbasari29/xjx-code/service"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(jwt service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := helper.ResponseFormatter(http.StatusUnauthorized, "error", "Authorization needed.", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		token, err := jwt.ValidateToken(authHeader)
		if token.Valid {
			// claim := token.Claims(jwt.MapClaims)
		} else {
			log.Println(err)
			response := helper.ResponseFormatter(http.StatusUnauthorized, "error", errors.New("token is invalid."), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
