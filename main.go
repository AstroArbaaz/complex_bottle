package main

import (
	"log"

	"github.com/AstroArbaaz/DevOps/complex_bottle/database"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Servers is Up and Running...")
}

func main () {
	app := fiber.New()
	database.ConnectDb()

	app.Get("/test", welcome)

	log.Fatal(app.Listen(":3000"))
}