package userController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rezaakbar35/task-5-pbi-btpns-reza-akbar-attariq/model"
	"gorm.io/gorm"
)

func GetAllUser(c *gin.Context) {
	var users []model.User

	if err := model.DB.Preload("Photos").Find(&users).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func GetUserById(c *gin.Context) {
	var user model.User
	id := c.Param("id")

	if err := model.DB.Preload("Photos").First(&user, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found!"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func CreateUser(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	model.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"message": "Create Success!", "user": user})
}

func UpdateUser(c *gin.Context) {
	var user model.User
	id := c.Param("id")

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if model.DB.Model(&user).Where("id = ?", id).Updates(&user).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Can't Update User!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Update Success!", "user": user})
}

func DeleteUser(c *gin.Context) {
	var user model.User
	id := c.Param("id")

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if model.DB.Delete(&user, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Can't Delete User!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Delete Success!", "user": user})

}
