package main

import (
	"fmt"
	"usersys/endpoint"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/lpernett/godotenv"
)

func main() {
	godotenv.Load(".env")
	app := fiber.New()
	app.Use(cors.New()) // ss

	app.Get("/Sayhi", endpoint.Sayhi)
	app.Post("/createUser", endpoint.CreateUser)
	app.Post("/login", endpoint.Login)
	app.Post("/order", endpoint.Order)

	app.Listen(":7200")
	fmt.Println("User system on !")
}
