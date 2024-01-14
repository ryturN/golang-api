package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-api/dto"
	"github.com/golang-api/helper"
)

//	func Auth(next http.Handler) http.Handler {
//		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//			tokenString := r.Header.Get("Authorization")
//			if tokenString == "" {
//				w.Header().Set("Content-Type", "application/json")
//				w.WriteHeader(http.StatusUnauthorized)
//				w.Write([]byte("unauthorized"))
//				return
//			}
//			_, err := helper.ValidateToken(tokenString)
//			if err != nil {
//				w.Header().Set("Content-Type", "application/json")
//				w.WriteHeader(http.StatusUnauthorized)
//				w.Write([]byte("unauthorized"))
//				return
//			}
//			next.ServeHTTP(w, r)
//		})
//	}
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie("token")
		fmt.Println("Access Token:", accessToken)
		if err != nil {
			res := helper.Response(dto.ResponseParams{
				StatusCode: http.StatusUnauthorized,
				Message:    "unauthorized",
				Data:       nil,
			})
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}

		user, err := helper.ValidateToken(accessToken)
		if err != nil {
			res := helper.Response(dto.ResponseParams{
				StatusCode: http.StatusUnauthorized,
				Message:    "unauthorized",
				Data:       nil,
			})
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}
		ctx := context.WithValue(c.Request.Context(), "userinfo", user)
		c.Set("ctx", c.Request.WithContext(ctx))
		c.Next()
	}
}
