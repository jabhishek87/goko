package main

import (
	"goko/configs"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	//run database
	configs.ConnectDB()

	engine := gin.Default()

	// Load all HTMLs
	engine.LoadHTMLGlob("templates/**/*")

	//Assets Folder  Add Static Folder
	engine.Static("/static", "./static")

	// API v1
	v1 := engine.Group("/api/v1")
	{
		v1.GET("item", getItems)
		v1.GET("item/:id", getItemById)
		v1.POST("item", addItem)
		v1.PUT("item/:id", updateItem)
		v1.DELETE("item/:id", deleteItem)
		// v1.OPTIONS("item", options)
	}

	// frontend group
	fe := engine.Group("frontend")
	{
		fe.GET("/", FeHome)
	}

	engine.NoRoute(func(c *gin.Context) {
		// c.AbortWithStatus(http.StatusNotFound)
		c.JSON(404, gin.H{
			"code": http.StatusNotFound, "message": "Page not found",
		})
	})

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	// engine.Run("0.0.0.0:8080")
	engine.Run("127.0.0.1:8080")
}
