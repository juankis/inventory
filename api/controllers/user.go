package controllers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/juankis/inventory/api/dao"
	"github.com/juankis/inventory/api/models"
	"github.com/juankis/inventory/api/utils"
)

//InsertUser insert user
func InsertUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		utils.InvalidJSON(c, err)
		return
	}
	err = dao.InsertUser(&user)
	if err != nil {
		utils.InvalidJSON(c, err)
		return
	}
	c.JSON(201, &user)
	return
}

//GetUserAll get all registries
func GetUserAll(c *gin.Context) {
	var users []models.User
	users, err := dao.GetUserAll()
	if err != nil {
		utils.CustomResponse(c, "getting users", err, 404)
		return
	}
	c.JSON(200, users)
	return
}

// DeleteUser delete some job in the db
func DeleteUser(c *gin.Context) {
	var user models.User
	userID, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		utils.InvalidURL(c, err)
		return
	}
	user.ID = userID
	err = dao.DeleteUser(&user)
	if err != nil {
		utils.CustomResponse(c, "deleting user", err, 400)
		return
	}
	c.Status(200)
	return
}

//PutUser actualizar registro
func PutUser(c *gin.Context) {
	var user models.User
	userID, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		utils.InvalidURL(c, err)
		return
	}
	err = c.ShouldBindJSON(&user)
	if err != nil {
		utils.InvalidJSON(c, err)
		return
	}
	user.ID = userID
	err = dao.UpdateUser(&user)
	if err != nil {
		utils.InvalidJSON(c, err)
		return
	}
	c.JSON(201, &user)
	return
}

//GetUser obtener registro
func GetUser(c *gin.Context) {
	var user models.User
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		utils.InvalidURL(c, err)
		return
	}
	user.ID = id
	err = dao.GetUser(&user)
	if err != nil {
		utils.CustomResponse(c, "getting user", err, 404)
		return
	}
	c.JSON(200, user)
	return
}

//Login obtener registro
func Login(c *gin.Context) {
	var user models.LoginRequest

	err := c.ShouldBindJSON(&user)
	if err != nil {
		fmt.Printf("error %v\n", err.Error())
		utils.InvalidJSON(c, err)
		return
	}
	res := dao.Login(&user)
	if res {
		c.JSON(200, gin.H{"message": "todo bien", "respond_code": 200})
		return
	}
	c.JSON(400, gin.H{"message": "usuario invalido", "respond_code": 400})
	return
}
