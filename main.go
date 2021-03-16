package main

import (
	"net/http"

	"./config"

	"./controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.DBInit()
	InDB := &controllers.InDB{DB: db}
	router := gin.Default()

	router.GET("/persons", InDB.GetPersons)
	router.GET("/person/:id", InDB.GetPerson)
	router.POST("/person", InDB.CreatePerson)
	router.PUT("/person/:id", InDB.UpdatePerson)
	router.DELETE("/person/:id", InDB.DeletePerson)

	router.GET("/books", InDB.GetBooks)
	router.GET("/book/:id", InDB.GetBook)
	router.POST("/book", InDB.CreateBook)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "po",
		})
	})
	router.Run()
}
