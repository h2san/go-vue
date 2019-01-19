package controllers

import (
	"log"
	"net/http"

	"github.com/h2san/h2sanGinJWT"

	"github.com/gin-gonic/gin"
)

//Index Index
func Login(c *gin.Context) {
	m := map[string]string{
		"username": "hillguo",
		"passwd":   "123456",
	}
	tokenStr, err := h2sanGinJWT.CreateToken(m)
	if err != nil {
		log.Println(err)
	}
	log.Println(tokenStr)
	c.JSON(http.StatusOK, gin.H{
		"code":   0,
		"errMsg": "ok",
		"data":   tokenStr,
	})
}
