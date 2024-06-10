package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMidleware(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		
		// if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		// 	return nil, fiber.NewError(fiber.StatusUnauthorized, "unexpected signature")

		// }

		return []byte([]byte(os.Getenv("JWT_SECRET"))), nil

	})

	if err != nil{
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid something")
	}

	if !token.Valid{
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid Token")
	}

	c.Locals("user", getUserFromToken(token))

	return c.Next()
	
}


func getUserFromToken(token *jwt.Token) interface{} {
    // Here, you can extract custom claims from the token
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        return nil
    }

    // For example, if your token has a "username" claim
    return claims["email"]
}
