package controllers

import (
	"net/http"
	"strconv"

	"../structs"
	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetBooks(c *gin.Context) {
	var (
		books  []structs.Book
		result gin.H
	)

	query := c.Query("title")

	if query == "" {
		idb.DB.Find(&books)
		if len(books) <= 0 {
			result = gin.H{
				"result": books,
				"count":  len(books),
			}
		} else {
			result = gin.H{
				"result": books,
				"count":  len(books),
			}
		}
	} else {
		idb.DB.Where("title LIKE ?", query+"%").Find(&books)
		result = gin.H{
			"result": books,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetBook(c *gin.Context) {
	var (
		book   structs.Book
		result gin.H
	)

	id := c.Param("id")

	err := idb.DB.First(&book, id).Error

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": book,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) CreateBook(c *gin.Context) {
	var (
		book   structs.Book
		result gin.H
	)

	title := c.PostForm("title")
	author := c.PostForm("author")
	releaseDate := c.PostForm("release_date")
	release_date, _ := strconv.Atoi(releaseDate)
	book.Title = title
	book.Author = author
	book.Release_Date = release_date
	idb.DB.Create(&book)
	result = gin.H{
		"result": book,
	}

	c.JSON(http.StatusOK, result)
}
