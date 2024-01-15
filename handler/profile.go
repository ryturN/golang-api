package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-api/dto"
	"github.com/golang-api/helper"
)

func MyProfile(c *gin.Context) {
	ctx := c.Value("ctx").(context.Context) // Access the updated context
	users, ok := ctx.Value("userinfo").(*helper.GetJWT)
	fmt.Println(users)
	if !ok {
		res := helper.Response(dto.ResponseParams{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to retrieve user claims",
			Data:       nil,
		})
		c.JSON(http.StatusInternalServerError, res)
		return
	} else {

		res := helper.Response(dto.ResponseParams{
			StatusCode: http.StatusOK,
			Message:    "My profile",
			Data:       users,
		})
		c.JSON(http.StatusOK, res)
	}

}
