package errorhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-api/dto"
	"github.com/golang-api/helper"
)

func HandleError(c *gin.Context, e error) {
	var statusCode int

	switch e.(type) {
	case *NotFoundError:
		statusCode = http.StatusNotFound
	case *BadRequestError:
		statusCode = http.StatusBadRequest
	case *InternalServerError:
		statusCode = http.StatusInternalServerError
	case *UnauthorizedError:
		statusCode = http.StatusUnauthorized
	}

	response := helper.Response(dto.ResponseParams{
		StatusCode: statusCode,
		Message:    e.Error(),
	})

	c.JSON(statusCode, response)
}
