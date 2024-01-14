package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-api/dto"
	"github.com/golang-api/helper"
)

func MyProfile(c *gin.Context) {
	claims := c.Request.Context().Value("userinfo").(*helper.GetJWT)

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "My profile",
		Data:       claims,
	})
	c.JSON(http.StatusOK, res)

}
