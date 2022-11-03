package main

import (
	"goko/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getItems(c *gin.Context) {
	// c.JSON(http.StatusOK, gin.H{"message": "getItems Called"})
	var items []models.Item
	models.DB.Find(&items)
	c.JSON(http.StatusOK, gin.H{"data": items})
}

func getItemById(c *gin.Context) {
	// id := c.Param("id")
	// c.JSON(http.StatusOK, gin.H{"message": "getPersonById " + id + " Called"})
	var item models.Item
	if err := models.DB.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": item})

}

func addItem(c *gin.Context) {
	// c.JSON(http.StatusOK, gin.H{"message": "addItem Called"})
	var input models.CreateItem
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item := models.Item{Data: input.Data, Tags: input.Tags}
	models.DB.Create(&item)

	c.JSON(http.StatusOK, gin.H{"data": item})
}

func updateItem(c *gin.Context) {
	// id := c.Param("id")
	// c.JSON(http.StatusOK, gin.H{"message": "updateItem Called with " + id})

	// Get model if exist
	var item models.Item
	if err := models.DB.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.UpdateItem

	//log.Println(input)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Model(&item).Updates(models.Item{Tags: input.Tags, Data: input.Data}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": item})
}

func deleteItem(c *gin.Context) {
	// id := c.Param("id")
	// c.JSON(http.StatusOK, gin.H{"message": "deleteItem " + id + " Called"})

	// Get model if exist
	var item models.Item
	if err := models.DB.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&item)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
