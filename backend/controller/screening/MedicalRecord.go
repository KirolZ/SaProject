package controller

import (
	"github.com/KirolZ/SaProject/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

func ListMedicalRecord(c *gin.Context) {
	var MedicalRecord []entity.MedicalRecord
	if err := entity.DB().Table("medical_records").Find(&MedicalRecord).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": MedicalRecord})
}
