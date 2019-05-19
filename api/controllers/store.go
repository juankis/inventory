package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/juankis/inventory/api/dao"
	"github.com/juankis/inventory/api/models"
	"github.com/juankis/inventory/api/utils"
)

func InsertStore(c *gin.Context) {
	var store models.Store
	err := c.ShouldBindJSON(&store)
	if err != nil {
		utils.InvalidJSON(c, err)
		return
	}
	err = dao.InsertStore(&store)
	if err != nil {
		utils.InvalidJSON(c, err)
		return
	}
	c.JSON(201, &store)
	return
}

//GetStoreAll get all registries
func GetStoreAll(c *gin.Context) {
	var stores []models.Store
	stores, err := dao.GetStoreAll()
	if err != nil {
		utils.CustomResponse(c, "getting stores", err, 404)
		return
	}
	c.JSON(200, stores)
	return
}

// DeleteStore delete some job in the db
func DeleteStore(c *gin.Context) {
	var store models.Store
	storeID, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		utils.InvalidURL(c, err)
		return
	}
	store.ID = storeID
	err = dao.DeleteStore(&store)
	if err != nil {
		utils.CustomResponse(c, "deleting store", err, 400)
		return
	}
	c.Status(200)
	return
}

//PutStore actualizar registro
func PutStore(c *gin.Context) {
	var store models.Store
	storeID, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		utils.InvalidURL(c, err)
		return
	}
	err = c.ShouldBindJSON(&store)
	if err != nil {
		utils.InvalidJSON(c, err)
		return
	}
	store.ID = storeID
	err = dao.UpdateStore(&store)
	if err != nil {
		utils.InvalidJSON(c, err)
		return
	}
	c.JSON(201, &store)
	return
}

//GetStore obtener registro
func GetStore(c *gin.Context) {
	var store models.Store
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		utils.InvalidURL(c, err)
		return
	}
	store.ID = id
	err = dao.GetStore(&store)
	if err != nil {
		utils.CustomResponse(c, "getting store", err, 404)
		return
	}
	c.JSON(200, store)
	return
}

//GetStoreStock get all registries
func GetStoreStock(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		utils.InvalidURL(c, err)
		return
	}
	var products []models.StockProduct
	products, err = dao.GetProductsStock(id)
	if err != nil {
		utils.CustomResponse(c, "getting products stock", err, 404)
		return
	}
	c.JSON(200, products)
	return
}
