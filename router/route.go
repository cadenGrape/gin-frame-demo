package router

import (
	"gin-frame/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Load(g *gin.Engine) *gin.Engine {
	g.Use(gin.Recovery())
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 not found")
	})
	g.GET("/", controller.Index)
	return g
}
