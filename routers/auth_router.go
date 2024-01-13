package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-api/handler"
	"github.com/golang-api/models"
	"github.com/golang-api/repository"
	"github.com/golang-api/service"
)

func AuthRouter(v3 *gin.RouterGroup) {
	authRepository := repository.NewAuthRepository(models.DB)
	authService := service.NewAuthService(authRepository)
	authHandler := handler.NewAuthHandler(authService)

	v3.POST("/register", authHandler.Register)
	v3.POST("/login", authHandler.Login)

}
