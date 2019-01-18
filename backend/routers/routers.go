package routers

import (
	"go-vue/backend/controllers"

	"github.com/gin-gonic/gin"
)

//RegisterRouter RegisterRouter
func RegisterRouter(server *gin.Engine) {
	server.GET("/", controllers.Index)
}

//
func RegisterStaticFiles(server *gin.Engine) {
	server.Static("/static","")
}
