package main

import (
	"customer-service/db"

	"github.com/gin-gonic/gin"
)

func main() {

	secret := db.GetSecretValue()
	db.GetDB(secret)

	r := gin.Default()

	r.POST("/customers", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "customer created",
		})
	})

	r.GET("/customers/:customerId", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "customer found",
		})
	})

	r.PUT("/customers/:customerId", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "customer updated",
		})
	})

	r.DELETE("/customers/:customerId", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "customer deleted",
		})
	})

	r.Run("localhost:8080") // listen and serve on

}
