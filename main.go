package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 책 자료
type book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

var books = []book{
	{ID: 1, Title: "Clean Code", Author: "James", ISBN: "971-80"},
	{ID: 2, Title: "개발자 글쓰기", Author: "bigdaditor", ISBN: "971-81"},
}

func getBooks(book *gin.Context) {
	book.IndentedJSON(http.StatusOK, books)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.Run(":8080")
}
