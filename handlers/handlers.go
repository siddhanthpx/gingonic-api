package handlers

import (
	"gingonic-api/data"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NewBook struct {
	Name   string `json:"name"`
	Author string `json:"author"`
}

func GetBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(err)
	}

	book := data.FindBook(id)
	if book == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Book with ID not found",
		})
		return
	}
	c.JSON(http.StatusOK, book)

}

func GetBooks(c *gin.Context) {
	books := data.Books

	if books == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to fetch books",
		})
	}

	c.JSON(http.StatusOK, books)
}

func AddBook(c *gin.Context) {
	var bookInput NewBook

	if err := c.ShouldBindJSON(&bookInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to add new book",
		})
		return
	}

	book := data.Book{Name: bookInput.Name, Author: bookInput.Author}
	data.AddBook(&book)

	c.JSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(err)
	}

	book := data.FindBook(id)
	if book == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Book with ID not found",
		})
		return
	}

	var bookInput *NewBook

	if err := c.ShouldBindJSON(&bookInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to update new book",
		})
		return
	}

	for i, b := range data.Books {
		if b.Id == id {
			data.Books[i].Name = bookInput.Name
			data.Books[i].Author = bookInput.Author
			c.JSON(http.StatusOK, data.Books)
			return
		}
	}

	c.JSON(404, gin.H{
		"message": "Could not edit book",
	})

}
