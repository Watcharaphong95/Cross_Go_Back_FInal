package controller

import (
	"final/dto"
	"final/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

func ProductController(router *gin.Engine) {
	routes := router.Group("/product")
	{
		routes.GET("", searchProduct)
		routes.POST("/add", addItemToCart)
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
	fmt.Print(product.ProductName)
	products := []model.Product{}
	result := db.Where("(product_name LIKE ? AND description LIKE ?) AND (price BETWEEN ? AND ?)",
		"%"+product.ProductName+"%",
		"%"+product.Description+"%",
		product.Min,
		product.Max).
		Find(&products)
	if result.Error != nil {
		panic(result.Error)
	}
	c.JSON(200, gin.H{
		"message": products,
	})
}

func addItemToCart(c *gin.Context) {
	cartAdd := dto.CartItem{}
	err := c.ShouldBindJSON(&cartAdd)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	cartItems := []model.CartItem{}
	result := db.Joins("ProductData").Joins("CartData").Joins("CartData.CustomerData").Where("CartData.customer_id = ?", cartAdd.CustomerID).Find(&cartItems)
	if result.Error != nil {
		panic(result.Error)
	}
	c.JSON(200, gin.H{
		"message": cartItems,
	})
}
