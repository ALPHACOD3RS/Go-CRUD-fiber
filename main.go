package main

import (
	"github.com/gofiber/fiber/v2"
)

func main(){

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, this is a GET API")
	 })

	 app.Post("/", func(c *fiber.Ctx) error{
		return c.SendString("hello this is a post request")
	 })


	 app.Listen(":3000")

	 

}