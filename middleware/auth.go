package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		tokenValue, err := ctx.Cookie("session_token")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusSeeOther, model.NewErrorResponse("unauthorized"))
			return
		}

		token, err := jwt.Parse(tokenValue, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return model.JwtKey, nil
		})
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, model.NewErrorResponse("unauthorized"))
			return
		}

		// cast claims interface to mapClaims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.NewErrorResponse("unauthorized"))
			return
		}

		// convert map to json then turn it into struct model.Claims
		b, _ := json.Marshal(claims)
		var customClaims model.Claims
		json.Unmarshal(b, &customClaims)

		// set jwt payload to gin context that can be shared within a request
		ctx.Set("email", customClaims.Email)
		// ctx.Set("email", customClaims.Email)
		// ctx.Set("isAdmin", customClaims.Scope == "ADMIN")
		ctx.Next()
		// TODO: answer here
	})
}

func Auth1() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		tokenValue, err := ctx.Cookie("session_token")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.NewErrorResponse("unauthorized"))
			return
		}

		token, err := jwt.Parse(tokenValue, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return model.JwtKey, nil
		})
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, model.NewErrorResponse("unauthorized"))
			return
		}

		// cast claims interface to mapClaims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.NewErrorResponse("unauthorized"))
			return
		}

		// convert map to json then turn it into struct model.Claims
		b, _ := json.Marshal(claims)
		var customClaims model.Claims
		json.Unmarshal(b, &customClaims)

		// set jwt payload to gin context that can be shared within a request
		ctx.Set("email", customClaims.Email)
		// ctx.Set("email", customClaims.Email)
		// ctx.Set("isAdmin", customClaims.Scope == "ADMIN")
		ctx.Next()
		// TODO: answer here
	})
}
