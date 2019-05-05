package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/juankis/inventory/api/dao"
	"github.com/juankis/inventory/api/models"
	"github.com/juankis/inventory/api/utils"
)

func InsertProduct(c *gin.Context) {
	var product models.Product
	err := c.ShouldBindJSON(&product)
	if err != nil {
		utils.InvalidJSON(c, err)
		return
	}
	err = dao.InsertProduct(&product)
	if err != nil {
		utils.InvalidJSON(c, err)
		return
	}
	c.JSON(201, &product)
	return
}

//GetProductAll get all registries
func GetProductAll(c *gin.Context) {
	var products []models.Product
	products, err := dao.GetProductAll()
	if err != nil {
		utils.CustomResponse(c, "getting products", err, 404)
		return
	}
	c.JSON(200, products)
	return
}

// DeleteProduct delete some job in the db
func DeleteProduct(c *gin.Context) {
	var product models.Product
	productID, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		utils.InvalidURL(c, err)
		return
	}
	product.ID = productID
	err = dao.DeleteProduct(&product)
	if err != nil {
		utils.CustomResponse(c, "deleting product", err, 400)
		return
	}
	c.Status(200)
	return
}

//PutProduct actualizar registro
func PutProduct(c *gin.Context) {
	var product models.Product
	productID, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		utils.InvalidURL(c, err)
		return
	}
	err = c.ShouldBindJSON(&product)
	if err != nil {
		utils.InvalidJSON(c, err)
		return
	}
	product.ID = productID
	err = dao.UpdateProduct(&product)
	if err != nil {
		utils.InvalidJSON(c, err)
		return
	}
	c.JSON(201, &product)
	return
}

//GetProduct obtener registro
func GetProduct(c *gin.Context) {
	var product models.Product
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		utils.InvalidURL(c, err)
		return
	}
	product.ID = id
	err = dao.GetProduct(&product)
	if err != nil {
		utils.CustomResponse(c, "getting product", err, 404)
		return
	}
	c.JSON(200, product)
	return
}
