package controller

import (
	"final/dto"
	"final/model"

	"github.com/gin-gonic/gin"
)

func ProductController(router *gin.Engine) {
	routes := router.Group("/customer")
	{
		routes.GET("", searchProduct)
	}
}

func searchProduct(c *gin.Context) {
	product := dto.Product{}
	err := c.ShouldBindJSON(&product)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}
	products := []model.Product{}
	result := db.Where("product_name = ?", product.ProductName).Find(&products)
	if result.Error != nil {
		panic(result.Error)
	}
	c.JSON(200, gin.H{
		"message": products,
	})
}
