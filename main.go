package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Book Model
type Book struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Isbn   int     `json:"isbn"`
	Author *Author `json:"author"`
}

//Author Model
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []Book

func main() {
	//Create new server
	r := gin.Default()

	//Mock Data
	books = []Book{
		{ID: 1, Title: "Book One", Isbn: 234, Author: &Author{Firstname: "John", Lastname: "Doe"}},
		{ID: 2, Title: "Book Two", Isbn: 468, Author: &Author{Firstname: "Steve", Lastname: "Smith"}},
	}

	//Functions and Handlers
	r.GET("/books", GetBooks)
	r.GET("/books/:id", GetBook)
	r.POST("/books", AddBook)

	//Listen and serve on localhost:8080/
	r.Run()
}

//GetBooks function
func GetBooks(c *gin.Context) {
	for _, book := range books {
		c.JSON(http.StatusOK, book)
	}
}

//GetBook function returns values with the given id
func GetBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, book := range books {
		if book.ID == id {
			c.JSON(http.StatusOK, book)
			return
		}
	}
	c.String(http.StatusNotFound, "book with id %v not found.", id)
}

//AddBook function adds a book provided in json format
func AddBook(c *gin.Context) {
	requestbody := Book{}
	c.Bind(&requestbody)

	book := Book{
		ID:     requestbody.ID,
		Title:  requestbody.Title,
		Isbn:   requestbody.Isbn,
		Author: requestbody.Author,
	}

	books = append(books, book)
}
