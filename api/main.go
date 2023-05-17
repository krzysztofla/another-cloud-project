package main

import (
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
	r.Run()
}
