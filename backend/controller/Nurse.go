package controller

import (
	"github.com/KirolZ/SaProject/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// GET /user/:id

func GetNurse(c *gin.Context) {

	var user entity.Nurse

	if err := entity.DB().Table("nurses").Where("nurse_email = ?", "nueng@gmail.com").First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}
