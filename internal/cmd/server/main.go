package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	// log.Init("tagger")

	dev, ok := os.LookupEnv("DEVELOPMENT")
	if !ok || dev != "true" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "3620"
	}

	router := gin.Default()
	api, err := injectApi()
	fmt.Println(api)

	if err != nil {
		log.Fatalf(fmt.Sprintf("cannot inject api: %v", err))
	}

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	err = router.Run(":" + port)

	if err != nil {
		log.Fatalf(fmt.Sprintf("Server Error: %v", err))
	}
}
