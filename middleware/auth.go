package middleware

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-api/dto"
	"github.com/golang-api/helper"
)

// func Auth(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		accessToken := r.Header.Get("Authorization")
// 		if accessToken == "" {
// 			w.Header().Set("Content-Type", "application/json")
// 			w.WriteHeader(http.StatusUnauthorized)
// 			w.Write([]byte("unauthorized"))
// 			return
// 		}
// 		user, err := helper.ValidateToken(accessToken)
// 		if err != nil {
// 			w.Header().Set("Content-Type", "application/json")
// 			w.WriteHeader(http.StatusUnauthorized)
// 			w.Write([]byte("unauthorized"))
// 			return
// 		}
// 		ctx := context.WithValue(r.Context(), "userinfo", user)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie("token")
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

		users, err := helper.ValidateToken(accessToken)
		// fmt.Println(users)
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

		ctx := context.WithValue(c.Request.Context(), "userinfo", users)
		c.Set("ctx", ctx) // Store the gin.Context directly
		c.Next()
	}
}
