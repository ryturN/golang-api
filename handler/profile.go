package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-api/dto"
	"github.com/golang-api/entity"
	"github.com/golang-api/errorhandler"
	"github.com/golang-api/helper"
	"github.com/golang-api/models"
	"github.com/golang-api/service"
	"gorm.io/gorm"
)

type profileHandler struct {
	service service.ProfileService
}

func NewProfileHandler(s service.ProfileService) *profileHandler {
	return &profileHandler{
		service: s,
	}
}

func MyProfile(c *gin.Context) {
	var user entity.Users
	var foto entity.Foto
	usersId, _ := c.Get("users")

	// if err := models.DB.Preload("fotos").First(&user, "users_id=?", users.UsersId); err != nil {
	// 	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
	// 		return
	// 	}
	if err := models.DB.First(&user, "users_id=?", usersId); err != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			return
		}
		if err := models.DB.First(&foto, "users_id=?", usersId); err != nil {
			if errors.Is(err.Error, gorm.ErrRecordNotFound) {
				foto.Url = ""
				return
			}

			fmt.Println(foto.UsersId)
			res := helper.Response(dto.ResponseParams{
				StatusCode: http.StatusOK,
				Message:    "My profile",
				Data: &dto.ProfileResponse{
					FullName: user.FullName,
					Username: user.Username,
					Email:    user.Email,
					Url:      foto.Url,
				},
			})
			c.JSON(http.StatusOK, res)
		}
	}
}

func FindUserById(c *gin.Context) {
	var users entity.Users
	var foto entity.Foto

	// get the id from the url parameter
	username := c.Param("username")

	// query the database for the user with the given id
	if err := models.DB.First(&users, "username=?", username).Error; err != nil {
		// if the user does not exist, return a not found error
		c.JSON(http.StatusNotFound, gin.H{
			"status": "fail",
			"error":  "user not found!",
		})
		return
	}
	if err := models.DB.First(&foto, "users_id=?", users.UsersId); err != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			foto.Url = ""
			return
		}

		// return the user if found
		res := helper.Response(dto.ResponseParams{
			StatusCode: http.StatusOK,
			Message:    "find user success",
			Data: &dto.ProfileResponse{
				Username: users.Username,
				FullName: users.FullName,
				Email:    users.Email,
				Url:      foto.Url,
			},
		})
		c.JSON(http.StatusAccepted, res)
	}
}

func (h *profileHandler) UpdateProfile(c *gin.Context) {
	var users dto.UpdateProfileRequest
	id, _ := c.Get("users")
	if err := c.ShouldBind(&users); err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}
	fmt.Println(id)

	users.UsersId = id.(string)

	if err := h.service.UpdateProfile(&users); err != nil {
		errorhandler.HandleError(c, err)
		return
	}
	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusAccepted,
		Message:    "Successfully Update Profile",
	})
	c.JSON(http.StatusAccepted, res)

}
