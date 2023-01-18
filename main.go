package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	var u User
	if err := c.BindJSON(&u); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": u})
}
