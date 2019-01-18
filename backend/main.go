package main

import (
	"fmt"
	"go-vue/backend/config"
	"go-vue/backend/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	conf := config.GetConfig()
	server := gin.Default()
	routers.RegisterRouter(server)
	server.Static("/static", "../front/dist")
	server.Run(fmt.Sprintf("%s:%d", conf.HTTPServer.IP, conf.HTTPServer.Port))
}
