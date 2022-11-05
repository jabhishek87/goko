package main

import (
	"context"
	"goko/configs"
	"goko/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var itemCollection *mongo.Collection = configs.GetCollection(configs.DB, "items")

func getItems(c *gin.Context) {
	// c.JSON(http.StatusOK, gin.H{"message": "getItems Called"})
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var items []models.Item
	defer cancel()

	results, _ := itemCollection.Find(ctx, bson.M{})
	// models.DB.Find(&items)
	// c.JSON(http.StatusOK, gin.H{"data": items})

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleItem models.Item
		if err := results.Decode(&singleItem); err != nil {
			c.JSON(http.StatusInternalServerError, models.ItemResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		}
		log.Println(singleItem)
		items = append(items, singleItem)
	}

	c.JSON(http.StatusOK,
		models.ItemResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"items": items}},
	)
}

func getItemById(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ID := c.Param("id")
	var item models.Item
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(ID)

	err := itemCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&item)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ItemResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	c.JSON(http.StatusOK, models.ItemResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"item": item}})
}

func addItem(c *gin.Context) {
	// c.JSON(http.StatusOK, gin.H{"message": "addItem Called"})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var item models.CreateItem
	defer cancel()

	//validate the request body
	if err := c.BindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, models.ItemResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	//use the validator library to validate required fields
	// if validationErr := validate.Struct(&item); validationErr != nil {
	// 	c.JSON(http.StatusBadRequest, models.ItemResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
	// 	return
	// }

	newItem := models.CreateItem{
		Data: item.Data,
		Tags: item.Tags,
	}

	result, err := itemCollection.InsertOne(ctx, newItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ItemResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	c.JSON(http.StatusCreated, models.ItemResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})

}

func updateItem(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ID := c.Param("id")
	var item models.Item
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(ID)
	log.Println(objId)
	//validate the request body
	if err := c.BindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, models.ItemResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	// //use the validator library to validate required fields
	// if validationErr := validate.Struct(&item); validationErr != nil {
	// 	c.JSON(http.StatusBadRequest, models.ItemResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
	// 	return
	// }

	update := bson.M{"data": item.Data, "tags": item.Tags}

	result, err := itemCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ItemResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	//get updated item details
	var updatedItem models.Item
	if result.MatchedCount == 1 {
		err := itemCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedItem)

		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ItemResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
	}

	c.JSON(http.StatusOK, models.ItemResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedItem}})
}

func deleteItem(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ID := c.Param("id")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(ID)

	result, err := itemCollection.DeleteOne(ctx, bson.M{"_id": objId})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ItemResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	if result.DeletedCount < 1 {
		c.JSON(http.StatusNotFound,
			models.ItemResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "User with specified ID not found!"}},
		)
		return
	}

	c.JSON(http.StatusOK,
		models.ItemResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "User successfully deleted!"}},
	)
}
