package logincontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-api/models"
)

type users struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var user *users

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if user.Username == "ryan" && user.Password == "ryan" {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "login success",
		})
	} else {

		c.JSON(http.StatusNotAcceptable, gin.H{
			"status":  "fail",
			"message": "login fail",
		})
	}
}

func Register(c *gin.Context) {
	var register models.Users

	if err := c.BindJSON(&register); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	models.DB.Create(&register)
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "register success",
	})

}

func FindUser(c *gin.Context) {
	var users []models.Users

	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "find user success",
		"data":    users,
	})
}

func FindUserById(c *gin.Context) {
	var users models.Users

	id := c.Param("id")
	if err := models.DB.First(&users, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "fail",
			"error":  "user not found!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    users,
		"status":  "success",
		"message": "find user success",
	})
}
func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "api connected"})
}
