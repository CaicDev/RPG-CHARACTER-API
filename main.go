package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()

	server.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Hello")
	})

	server.Run("127.0.0.1:8080")
}