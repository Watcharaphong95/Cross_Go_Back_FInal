package controller

import (
	"final/dto"
	"final/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

var db = Conn()

func CustomerController(router *gin.Engine) {
	routes := router.Group("/customer")
	{
		routes.GET("", getCustomerData)
		routes.POST("", login)
		routes.POST("/update", updateCustomer)
	}
}

func getCustomerData(c *gin.Context) {
	user := dto.User{}
	err := c.ShouldBindJSON()
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}
}

func login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	customer := model.Customer{}
	result := db.Where("email = ?", email).Take(&customer)
	if result.Error != nil {
		panic(result.Error)
	}

	check := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(password))
	if check != nil {
		c.JSON(400, gin.H{
			"message": "Invalid password",
		})
	}

	user := dto.User{}
	copier.Copy(&user, &customer)
	c.JSON(200, gin.H{
		"message": user,
	})
}

func updateCustomer(c *gin.Context) {
	user := dto.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}
	customer := model.Customer{}
	result := db.Where("email = ?", user.Email).Take(&customer)
	if result.Error != nil {
		panic(result.Error)
	}
	if user.FirstName != "" {
		customer.FirstName = user.FirstName
	}
	if user.LastName != "" {
		customer.LastName = user.LastName
	}
	if user.PhoneNumber != "" {
		customer.PhoneNumber = user.PhoneNumber
	}
	if user.Address != "" {
		customer.Address = user.Address
	}

	db.Save(&customer)
	c.JSON(200, gin.H{
		"message": customer,
	})
}
