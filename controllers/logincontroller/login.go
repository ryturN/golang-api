package logincontroller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-api/dto"
	"github.com/golang-api/models"
)

type Users struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var login Users

	if err := c.BindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}

func Register(c *gin.Context) {
	var register models.Users
	// hash, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	if err := c.BindJSON(&register); err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("error on field %s , condition : %s ", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
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
	var users []dto.RegisterRequest

	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "find user success",
		"data":    users,
	})
}

// FindUserById finds a user by their ID
func FindUserById(c *gin.Context) {
	var users models.Users

	// get the id from the url parameter
	id := c.Param("id")

	// query the database for the user with the given id
	if err := models.DB.First(&users, id).Error; err != nil {
		// if the user does not exist, return a not found error
		c.JSON(http.StatusNotFound, gin.H{
			"status": "fail",
			"error":  "user not found!",
		})
		return
	}

	// return the user if found
	c.JSON(http.StatusOK, gin.H{
		"data":    users,
		"status":  "success",
		"message": "find user success",
	})
}
func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "api connected"})
}
