package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Index Index
func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":   0,
		"errMsg": "ok",
		"data":   "",
	})
}
