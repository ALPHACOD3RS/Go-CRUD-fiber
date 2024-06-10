package main

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func main(){

	app := fiber.New()
	db := initDatabase()

	app.Post("/users", func(c *fiber.Ctx) error {
		user := new(User)

		if err := c.BodyParser(user); err != nil{
			return c.Status(400).SendString(err.Error())
		}
		db.Create(&user)
		return c.JSON(user)
	})

	app.Post("/register", regiserUserHandler)
	app.Post("/login", loginHandler)
	

	app.Get("/users", func(c *fiber.Ctx) error {
		var users []User
		db.Find(&users)
		return c.JSON(users)
	})


	app.Get("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		var user User

		if err := db.First(&user, id); err.Error != nil {
			return c.Status(400).SendString("user not found with this id")
		}

		return c.JSON(user)
	})

	app.Delete("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		var user User

		if err := db.First(&user, id); err.Error != nil{
			return c.Status(400).SendString("user with this id doesnt exist")
		}
		db.Delete(&user)
		return c.SendString("user deleted successfully")
	})

	

	app.Put("/users/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
		        log.Printf("Invalid user ID: %s", idStr)
		        return c.Status(400).SendString("Invalid user ID")
		    }
		
		var user User

		if err := db.First(&user, id); err.Error != nil{
			log.Printf("Error finding user: %v", err)
			return c.Status(404).SendString("user not found " )
		}

		if err := c.BodyParser(&user); err != nil{
			return c.Status(400).SendString(err.Error())
		}
		db.Save(&user)
		return c.JSON(user)

	})
	 

	
	
	 app.Listen(":3000")


	 

}