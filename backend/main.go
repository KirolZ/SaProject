package main

import (
	"github.com/KirolZ/SaProject/entity"

	controller "github.com/KirolZ/SaProject/controller/screening"

	"github.com/KirolZ/SaProject/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	api := r.Group("")
	{
		protected := api.Use(middlewares.Authorizes())
		{
			// User Routes
			protected.GET("/api/GetNurse/:id", controller.GetNurse)
			protected.GET("/api/Disease", controller.ListDisease)
			protected.GET("/api/MedicalRecord", controller.ListMedicalRecord)
			protected.GET("/api/PreloadScreenings", controller.PreloadScreenings)

		}
	}

	// User Routes
	r.POST("/api/CreateScreening", controller.CreateScreening)

	// Authentication Routes
	r.POST("/login", controller.Login)

	//Run the server
	r.Run()

}

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
