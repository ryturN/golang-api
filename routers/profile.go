package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-api/handler"
	"github.com/golang-api/middleware"
	"github.com/golang-api/models"
	"github.com/golang-api/repository"
	"github.com/golang-api/service"
)

func UserRoutes(v3 *gin.RouterGroup) {
	profileRepository := repository.NewProfileRepository(models.DB)
	profileService := service.NewProfileService(profileRepository)
	profileHandler := handler.NewProfileHandler(profileService)
	v3.Use(middleware.Auth())

	v3.GET("/profile", handler.MyProfile)
	v3.GET("/profile/:username", handler.FindUserById)
	v3.PUT("/update", profileHandler.UpdateProfile)
}
