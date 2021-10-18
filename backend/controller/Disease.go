package controller

import (
	"github.com/KirolZ/SaProject/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

func ListDisease(c *gin.Context) {
	var Disease []entity.Disease
	if err := entity.DB().Table("diseases").Find(&Disease).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Disease})
}
