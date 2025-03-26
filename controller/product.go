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

	userCart := model.CartItem{}
	nameCartCheck := db.Joins("ProductData").Joins("CartData").Joins("CartData.CustomerData").Where("CartData.cart_name = ?", cartAdd.CartName).First(&userCart)
	if nameCartCheck.Error == nil {
		cartAdd := model.CartItem{
			CartID:    userCart.CartID,
			ProductID: userCart.ProductID,
			Quantity:  userCart.Quantity,
		}
		existingCartItem := model.CartItem{}
		err := db.Where("cart_id = ? AND product_id = ?", userCart.CartID, userCart.ProductID).First(&existingCartItem)
		if err.Error == nil {
			existingCartItem.Quantity = cartAdd.Quantity
			db.Save(existingCartItem)
			if result.Error != nil {
				panic(result.Error)
			}
			c.JSON(200, gin.H{
				"message": "Quanity item added",
			})
		} else {
			result := db.Create(&cartAdd)
			if result.Error != nil {
				panic(result.Error)
			}
			c.JSON(200, gin.H{
				"message": "New item added",
			})
		}
	} else {
		cart := model.Cart{
			CustomerID: cartAdd.CustomerID,
			CartName:   cartAdd.CartName,
		}
		result := db.Create(&cart)
		if result.Error != nil {
			panic(result.Error)
		}
		cartAdd_2 := model.CartItem{
			CartID:    cart.CartID,
			ProductID: cartAdd.ProductID,
			Quantity:  cartAdd.Quantity,
		}
		result1 := db.Create(&cartAdd_2)
		if result1.Error != nil {
			panic(result1.Error)
		}
		c.JSON(200, gin.H{
			"message": "New cart create and has been item",
		})
	}

	// c.JSON(200, gin.H{
	// 	"message": cartItems,
	// })
}
