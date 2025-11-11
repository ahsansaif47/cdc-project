package main

import (
	"fmt"
	"log"

	"github.com/ahsansaif47/cdc-app/config"
	"github.com/ahsansaif47/cdc-app/http/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// "gorm.io/gorm"

func init() {

}

func main() {

}

func startHTTP() {
	app := fiber.New()
	// Add logger middleware
	app.Use(logger.New())

	routes.InitRoutes(app)

	port := config.GetConfig().ServerPort
	log.Printf("Fiber server listening on port: %s", port)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
