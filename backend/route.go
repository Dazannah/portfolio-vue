package main

import "github.com/gin-gonic/gin"

func registerRoutes() *gin.Engine {

	router := gin.Default()
	router.Static("/", "../frontend/dist")

	router.NoRoute(func(c *gin.Context) {
		c.File("./first-vue-app/dist/index.html")
	})

	return router
}
