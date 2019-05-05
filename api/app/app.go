package app

import (
	"github.com/gin-gonic/gin"
	"github.com/juankis/inventory/api/controllers"
)

//Start function to start up proyect
func Start() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/transaction", controllers.InsertTransaction)
	r.GET("/transaction", controllers.GetTransactionAll)
	r.GET("/transaction/:id/", controllers.GetTransaction)
	r.DELETE("/transaction/:id/", controllers.DeleteTransaction)
	r.PUT("/transaction/:id/", controllers.PutTransaction)

	r.POST("/product", controllers.InsertProduct)
	r.GET("/product", controllers.GetProductAll)
	r.GET("/product/:id/", controllers.GetProduct)
	r.DELETE("/product/:id/", controllers.DeleteProduct)
	r.PUT("/product/:id/", controllers.PutProduct)

	r.POST("/user", controllers.InsertUser)
	r.GET("/user", controllers.GetUserAll)
	r.GET("/user/:id/", controllers.GetUser)
	r.DELETE("/user/:id/", controllers.DeleteUser)
	r.PUT("/user/:id/", controllers.PutUser)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
