package main

import (
	//"fmt"
	"github.com/gin-gonic/gin"
)

var count int = 0

func rootHandler(c *gin.Context){
	count++
	c.String(200, "Hello %d", count)
}

func helloHandler(c *gin.Context){
	c.String(200, "Hello!")
}

func main() {
	r := gin.Default()
	r.GET("/", rootHandler)
	r.GET("/hello", helloHandler)
	r.Run("localhost:8080")
}
