package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/juankis/inventory/api/dao"
	"github.com/juankis/inventory/api/models"
	"github.com/juankis/inventory/api/utils"
)

func InsertTransaction(c *gin.Context) {
	var reg models.Response
	err := c.ShouldBindJSON(&reg)
	if err != nil {
		utils.InvalidJSON(c, err)
		return
	}
	transaction := reg.Data
	err = dao.InsertTransaction(&transaction)
	if err != nil {
		utils.InvalidJSON(c, err)
		return
	}
	c.JSON(201, &transaction)
	return
}

//GetTransactionAll get all registries
func GetTransactionAll(c *gin.Context) {
	var transactions []models.Transaction
	transactions, err := dao.GetTransactionAll()
	if err != nil {
		utils.CustomResponse(c, "getting transactions", err, 404)
		return
	}
	c.JSON(200, transactions)
	return
}

// DeleteTransaction delete some job in the db
func DeleteTransaction(c *gin.Context) {
	var transaction models.Transaction
	transactionID, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		utils.InvalidURL(c, err)
		return
	}
	transaction.ID = transactionID
	err = dao.DeleteTransaction(&transaction)
	if err != nil {
		utils.CustomResponse(c, "deleting transaction", err, 400)
		return
	}
	c.Status(200)
	return
}

//PutTransaction actualizar registro
func PutTransaction(c *gin.Context) {
	var reg models.Response
	transactionID, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		utils.InvalidURL(c, err)
		return
	}
	err = c.ShouldBindJSON(&reg)
	if err != nil {
		utils.InvalidJSON(c, err)
		return
	}
	transaction := reg.Data
	transaction.ID = transactionID
	err = dao.UpdateTransaction(&transaction)
	if err != nil {
		utils.InvalidJSON(c, err)
		return
	}
	c.JSON(201, &transaction)
	return
}

//GetTransaction obtener registro
func GetTransaction(c *gin.Context) {
	var transaction models.Transaction
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		utils.InvalidURL(c, err)
		return
	}
	transaction.ID = id
	err = dao.GetTransaction(&transaction)
	if err != nil {
		utils.CustomResponse(c, "getting transaction", err, 404)
		return
	}
	c.JSON(200, models.Response{Data: transaction})
	return
}

func ConfirmTransaction(c *gin.Context) {
	var confirm models.ConfirmTransaction
	err := c.ShouldBindJSON(&confirm)
	if err != nil {
		utils.InvalidJSON(c, err)
		return
	}
	err = dao.ConfirmTransaction(&confirm)
	if err != nil {
		utils.InvalidJSON(c, err)
		return
	}
	c.JSON(201, nil)
	return
}