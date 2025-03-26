package controller

import (
	"final/dto"
	"final/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

func CustomerController(router *gin.Engine) {
	routes := router.Group("/customer")
	{
		routes.GET("", getCustomerData)
		routes.POST("", login)
		routes.POST("/update", updateCustomer)
		routes.POST("/forgetpassword", resetPassword)
	}
}

func resetPassword(c *gin.Context) {

	type ResetPasswordRequest struct {
		Email   string `json:"email" binding:"required"`
		OldPass string `json:"oldPass" binding:"required"`
		NewPass string `json:"newPass" binding:"required"`
	}

	user := ResetPasswordRequest{}
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
	check := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(user.OldPass))
	if check != nil {
		c.JSON(400, gin.H{
			"message": "Your old password is incorrect",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "Your password has been updated",
		})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.NewPass), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Hash Failed",
		})
	}

	customer.Password = string(hash)
	db.Save(&customer)
}

func getCustomerData(c *gin.Context) {
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
	copier.Copy(&user, &customer)
	c.JSON(200, gin.H{
		"message": user,
	})
}

func login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	customer := model.Customer{}
	result := db.Where("email = ?", email).Take(&customer)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Invalid Email",
		})
	}

	user := dto.User{}
	copier.Copy(&user, &customer)

	check := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(password))
	if check != nil {
		c.JSON(400, gin.H{
			"message": "Invalid password",
		})
	} else {
		c.JSON(200, gin.H{
			"message": user,
		})
	}
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
		c.JSON(404, gin.H{"error": "Customer not found"})
		return
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
	copier.Copy(&user, &customer)
	c.JSON(200, gin.H{
		"message": user,
	})
}
