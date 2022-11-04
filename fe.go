package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FeHome(c *gin.Context) {
	// id := c.Param("id")
	// c.JSON(http.StatusOK, gin.H{"message": "Home Called"})
	c.HTML(http.StatusOK, "pages/home.html", gin.H{
		"title": "Home Page",
		//"routes": routeSlice,
	})

}
