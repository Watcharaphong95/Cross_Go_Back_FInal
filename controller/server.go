package controller

import "github.com/gin-gonic/gin"

func StartServer() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "it work!",
		})
	})
	CustomerController(router)
	ProductController(router)

	router.Run()
}
