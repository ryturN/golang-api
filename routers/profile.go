package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-api/handler"
	"github.com/golang-api/middleware"
)

func UserRoutes(v3 *gin.RouterGroup) {
	v3.Use(middleware.Auth())
	v3.GET("/profile", handler.MyProfile)
}
