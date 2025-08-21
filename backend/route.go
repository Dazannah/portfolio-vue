package main

import (
	"github.com/gin-gonic/gin"
)

func registerRoutes() *gin.Engine {

	router := gin.Default()
	router.Static("/static", "../frontend/dist")
	router.Static("/assets", "../frontend/dist/assets")
	router.StaticFile("/cv.pdf", "./static/cv.pdf")

	router.NoRoute(func(c *gin.Context) {
		c.File("../frontend/dist/index.html")
	})

	return router
}
