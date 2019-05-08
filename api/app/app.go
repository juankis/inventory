package app

import (
	"net/http"

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
	r.POST("/confirm", controllers.ConfirmTransaction)

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

	//front
	//r.LoadHTMLFiles("templates/js/*")
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/transactions", func(c *gin.Context) {
		c.HTML(http.StatusOK, "listTransactions.tmpl", nil)
	})
	r.GET("/transactions/create", func(c *gin.Context) {
		c.HTML(http.StatusOK, "createTransaction.tmpl", nil)
	})
	r.GET("/users", func(c *gin.Context) {
		c.HTML(http.StatusOK, "listUsers.tmpl", nil)
	})
	r.GET("/users/create", func(c *gin.Context) {
		c.HTML(http.StatusOK, "createUser.tmpl", nil)
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.tmpl", nil)
	})
	r.GET("/products", func(c *gin.Context) {
		c.HTML(http.StatusOK, "listProducts.tmpl", nil)
	})
	r.GET("/products/create", func(c *gin.Context) {
		c.HTML(http.StatusOK, "createProduct.tmpl", nil)
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
