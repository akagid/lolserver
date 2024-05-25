package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	return router
}

func runServer(addr string) error {
	router := setupRouter()
	if err := router.Run(addr); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func main() {
	if err := runServer(":8080"); err != nil {
		log.Fatalln(err)
	}
}
