package routers

import (
	"go-vue/backend/controllers"

	"github.com/gin-gonic/gin"
	"github.com/h2san/h2sanGinJWT"
)

//RegisterRouter RegisterRouter
func RegisterRouter(server *gin.Engine) {
	route1 := server.Group("/admin")
	route1.Use(h2sanGinJWT.GinJWT(nil, "xxx"))
	route1.GET("/index", controllers.Index)
	route2 := server.Group("/index")
	route2.GET("/xx", controllers.Login)
}

//
func RegisterStaticFiles(server *gin.Engine) {
	server.Static("/static", "")
}
