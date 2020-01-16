package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	// Set the r as the default one shipped with Gin
	r := gin.Default()

	// Serve frontend static files
	r.Use(static.Serve("/", static.LocalFile("./views", true)))

	//r.StaticFile("/favicon.ico", "./views/favicon.ico")

	// Setup route group for the API
	api := r.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	/*
		blog := r.Group("/abusivefish")
		{
			blog.GET("/", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{

				})
			})
		}
	*/

	// Start and run the server
	r.Run(":3000")
}
