package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const (
	invalidURL  = "Invalid format URL"
	invalidJSON = "Invalid format JSON"
	internal    = "Internal error "
)

//InvalidURL function to retun error
func InvalidURL(c *gin.Context, err error) {
	fmt.Errorf(err.Error(), err)
	c.JSON(400, gin.H{"message": invalidURL, "specific error": err.Error()})
}

//InvalidJSON function to retun error
func InvalidJSON(c *gin.Context, err error) {
	fmt.Errorf(err.Error(), err)
	c.JSON(400, gin.H{"message": invalidJSON, "specific error": err.Error()})
}

//CustomResponse function to retun error
func CustomResponse(c *gin.Context, where string, err error, code int) {
	fmt.Errorf(err.Error(), err)
	c.JSON(code, gin.H{"message": internal + where, "error": err.Error()})
}
