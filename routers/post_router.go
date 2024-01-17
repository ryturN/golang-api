package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-api/handler"
	"github.com/golang-api/models"
	"github.com/golang-api/repository"
	"github.com/golang-api/service"
)

func PostRouter(v3 *gin.RouterGroup) {
	postRepository := repository.NewPostRepository(models.DB)
	postService := service.NewPostService(postRepository)
	postHandler := handler.NewPostHandler(postService)
	r := v3.Group("/posts")
	r.POST("/", postHandler.Create)
	r.GET("test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Connected!",
		})
	})
}
