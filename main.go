package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name  string `json:"name" binding:"required"`
	Intro string `json:"intro" binding:"required"`
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
	db, _ := sql.Open("mysql", "root:password@(localhost:3306)/local?parseTime=true")
	defer db.Close()

	var u User
	validate := validator.New()

	if err := c.BindJSON(&u); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "error"})
		return
	}
	validate.Struct(&u)
	c.JSON(http.StatusOK, gin.H{"result": u})

	insert, _ := db.Prepare("INSERT INTO user (name, intro) VALUES (?,?)")
	insert.Exec(u.Name, u.Intro)
}
