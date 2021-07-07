package middlewares

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/lenarbatdalov/learn/service"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := ctx.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]

		token, err := service.NewJWTService().ValidateToken(tokenString)

		if token.Valid {
			claim := token.Claims.(jwt.MapClaims)
			log.Println("Claims[Name]: ", claim["name"])
			log.Println("Claim[Admin]: ", claim["admin"])
			log.Println("Claim[Issuer]: ", claim["iss"])
			log.Println("Claim[IssuedAt]: ", claim["iat"])
			log.Println("Claim[ExpiresAt]: ", claim["exp"])
		} else {
			log.Println(err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
