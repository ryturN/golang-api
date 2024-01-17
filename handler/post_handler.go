package handler

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/golang-api/dto"
	"github.com/golang-api/errorhandler"
	"github.com/golang-api/helper"
	"github.com/golang-api/service"
	"github.com/google/uuid"
)

type postHandler struct {
	service service.PostService
}

func NewPostHandler(s service.PostService) *postHandler {
	return &postHandler{
		service: s,
	}
}

func (h *postHandler) Create(c *gin.Context) {
	var post dto.PostRequest

	if err := c.ShouldBind(&post); err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	if post.Picture != nil {
		if err := os.MkdirAll("public/picture", 0755); err != nil {
			errorhandler.HandleError(c, &errorhandler.InternalServerError{Message: err.Error()})
			return
		}

		// Rename picture
		ext := filepath.Ext(post.Picture.Filename)
		newFileName := uuid.New().String() + ext

		// Save image into directory
		dst := filepath.Join("public/picture", filepath.Base(newFileName))

		c.SaveUploadedFile(post.Picture, dst)
	}
	usersId := "user_T-hcRfeEKPVI_kdi9OeNi"
	post.UsersId = usersId
	if err := h.service.Create(&post); err != nil {
		errorhandler.HandleError(c, err)
		return
	}
	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Successfully created Post",
	})
	c.JSON(http.StatusCreated, res)
}