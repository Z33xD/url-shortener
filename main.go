package main

import (
	"fmt"
	"go-url-shortener/handler"
	"go-url-shortener/store"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Main function running")

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the URL Shortener API",
		})
	})

	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	store.InitialiseStore()

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed the start the web server. Error: %v", err))
	}
}
