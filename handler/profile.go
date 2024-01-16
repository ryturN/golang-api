package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-api/dto"
	"github.com/golang-api/entity"
	"github.com/golang-api/helper"
	"github.com/golang-api/models"
	"gorm.io/gorm"
)

func MyProfile(c *gin.Context) {
	var user entity.Users
	var foto entity.Foto
	ctx := c.Value("ctx").(context.Context)
	users, ok := ctx.Value("userinfo").(*helper.GetJWT)
	if !ok {
		res := helper.Response(dto.ResponseParams{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to retrieve user claims",
			Data:       nil,
		})
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	// if err := models.DB.Preload("fotos").First(&user, "users_id=?", users.UsersId); err != nil {
	// 	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
	// 		return
	// 	}
	if err := models.DB.First(&user, "users_id=?", users.UsersId); err != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			return
		}
		if err := models.DB.First(&foto, "users_id=?", users.UsersId); err != nil {
			if errors.Is(err.Error, gorm.ErrRecordNotFound) {
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
