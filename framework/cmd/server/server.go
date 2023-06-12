package main

import (
	"log"
	"os"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-gonic/gin"
	"github.com/hack-caixa/framework/config/database"
	"github.com/hack-caixa/framework/routes"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

}

func main() {
	https := os.Getenv("HTTPS_ENABLED")

	conn, err := database.DB.Connect()

	if err != nil {
		log.Fatalf("error connecting to DB")
	}

	r := gin.Default()
	r.Use(helmet.Default())

	router := routes.SetupRouter(r, conn)

	if https == "enabled" {

		err = router.RunTLS(":8000", "framework/config/certs/server.crt", "framework/config/certs/server.key")

		if err != nil {
			panic(err)
		}

	} else {

		router.Run(":8000")
	}

}
