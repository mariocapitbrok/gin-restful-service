package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Item struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

var items = []Item{
	{ID: "1", Name: "Item One"},
	{ID: "2", Name: "Item Two"},
}

func getItems(c *gin.Context) {
	c.JSON(http.StatusOK, items)
}

func getItem(c *gin.Context) {
	id := c.Param("id")

	for _, item := range items {
		if item.ID == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}
	
	c.JSON(http.StatusNotFound, gin.H{"message": "item not found"})
}

func addItem(c *gin.Context) {
	var newItem Item

	if err := c.BindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	items = append(items, newItem)
	c.JSON(http.StatusCreated, newItem)
}

func main() {
	r := gin.Default()

	r.GET("/items", getItems)
	r.GET("/items/:id", getItem)
	r.POST("/items", addItem)

	r.Run(":8080")
}