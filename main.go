package main

import (
	"github.com/abusivefish/portfolio-gin/httpd/handler"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	// Set the r as the default one shipped with Gin
	r := gin.Default()

	// Serve frontend static files
	r.Use(static.Serve("/", static.LocalFile("./client/build", true)))

	//r.StaticFile("/favicon.ico", "./views/favicon.ico")

	// Setup route group for the API
	blog := r.Group("/blog")
	{
		blog.GET("/", handler.BlogGet)
	}

	r.Run(":3000")
}
