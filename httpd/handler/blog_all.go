package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// BlogGet will grab the default or most recent post. Probably need some sort of  paginator at the model level.
func BlogGet(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"message": "pong",
	})	
}
