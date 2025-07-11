package main

import (
	"fmt"
	"log"

	"trello/config"
	"trello/database"
	"trello/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.LoadEnv()
	database.ConnectDB()
	app := fiber.New()
	app.Use(cors.New())
	routes.SetupRoutes(app)
	port := ":3001"
	fmt.Println("http://localhost" + port)
	log.Fatal(app.Listen(port))
}
