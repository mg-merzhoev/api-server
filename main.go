package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", pong)
	r.Run()

}

func pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"messsage": "pong",
	})
}
