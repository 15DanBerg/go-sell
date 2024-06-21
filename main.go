package main

import (
	"fmt"
	"log"
	"os"

	"github.com/15DanBerg/go-sell/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	start := os.Getenv("INIT")
	fmt.Println(start)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
