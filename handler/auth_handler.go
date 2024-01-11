package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-api/dto"
	"github.com/golang-api/errorhandler"
	"github.com/golang-api/helper"
	"github.com/golang-api/service"
)

type authHandler struct {
	service service.AuthService
}

func NewAuthHandler(s service.AuthService) *authHandler {
	return &authHandler{
		service: s}
}

func (h *authHandler) Register(c *gin.Context) {
	var register dto.RegisterRequest

	if err := c.ShouldBindJSON(&register); err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	if error := h.service.Register(&register); error != nil {
		errorhandler.HandleError(c, error)
		return
	}
	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "User registered successfully",
	})
	c.JSON(http.StatusCreated, res)
}
