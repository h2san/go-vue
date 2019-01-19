package controllers

import (
	"net/http"

	"github.com/h2san/h2sanGinJWT"

	"github.com/gin-gonic/gin"
)

//Index Index
func Index(c *gin.Context) {
	s := c.Value(h2sanGinJWT.DefaultGinJWTKey)
	c.JSON(http.StatusOK, gin.H{
		"code":   0,
		"errMsg": "ok",
		"data":   s,
	})
}
