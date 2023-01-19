package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json:"name" binding:"required"`
}

func router() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", get)
	r.POST("/ps", create)
	return r
}
func main() {
	router().Run()
}

func get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "pong"})
}

func create(c *gin.Context) {
	db, err := sql.Open("mysql", "root:password@(localhost:3306)/local?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var u User
	if err := c.BindJSON(&u); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": u})

	insert, _ := db.Prepare("INSERT INTO User (name) VALUES (?)")

	insert.Exec(u.Name)
}
