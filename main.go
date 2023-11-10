package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// Autres champs selon vos besoins
}

var items []Item

func main() {
	router := gin.Default()

	router.GET("/items", getItems)
	router.POST("/items", addItem)
	router.PUT("/items/:id", updateItem)
	router.DELETE("/items/:id", deleteItem)

	router.Run(":8080")
}

func getItems(c *gin.Context) {
	c.JSON(http.StatusOK, items)
}

func addItem(c *gin.Context) {
	var newItem Item
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	items = append(items, newItem)
	c.JSON(http.StatusCreated, newItem)
}

func updateItem(c *gin.Context) {
	id := c.Param("id")
	var updatedItem Item

	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Logic to find and update the item in the 'items' slice
	for index, item := range items {
		if item.ID == id {
			// Item found, update it
			items[index] = updatedItem

			// Send back the updated item or a success message
			c.JSON(http.StatusOK, updatedItem)
			return
		}
	}

	// If the item with the specified ID is not found
	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}

func deleteItem(c *gin.Context) {
	id := c.Param("id")

	// Logic to find and delete the item from the 'items' slice
	for i, item := range items {
		if item.ID == id {
			// Remove the item from the slice using append to exclude the item at the specific index
			items = append(items[:i], items[i+1:]...)

			// Sending a 204 No Content status to indicate a successful deletion
			c.Status(http.StatusNoContent)
			return
		}
	}

	// If the item with the specified ID is not found
	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}
