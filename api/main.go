package main

import (
	"github.com/demo-go-api/models"
	"github.com/demo-go-api/services"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/items", func(c *gin.Context) {
		is, err := services.NewItemService()
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})

			var items = is.GetAll()
			c.JSON(200, gin.H{
				"items": items,
			})
		}
	})

	r.GET("/items-static", func(c *gin.Context) {
		items := models.Items{
			models.Item{ID: 1, Name: "Item 1", Description: "Description 1"},
			models.Item{ID: 2, Name: "Item 2", Description: "Description 2"},
			models.Item{ID: 3, Name: "Item 3", Description: "Description 3"},
			models.Item{ID: 4, Name: "Item 4", Description: "Description 4"},
		}
		c.JSON(200, gin.H{
			"items": items,
		})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	r.Run()
}
