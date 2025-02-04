package main

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)


func HashPassword(password string) (string, error){
	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil{
		return "", err
	}

	return string(HashedPassword), nil
}




func regiserUserHandler(c *fiber.Ctx) error{
	user := new(User)

	db := initDatabase()

	if err := c.BodyParser(user); err != nil{
		return c.Status(400).SendString("plase fill all of the neccery fields")
	}

	pass, err := HashPassword(user.Password)
	if err != nil {
		return c.Status(301).SendString("something went wrong")
	}

	user.Password = pass
	if err := db.Create(&user).Error; err != nil{
		return c.Status(400).SendString("error while registering")
	}

	return c.JSON(user)
}


func loginHandler(c *fiber.Ctx) error{

	db := initDatabase()

	 inputUser := new(User)


	 if err := c.BodyParser(inputUser); err != nil{
		return c.Status(400).SendString(err.Error())
	 }

	 var user User

	 if err := db.Where("email = ?", inputUser.Email).First(&user).Error; err != nil{
		return c.Status(400).SendString("invalid credentials")
	 }

	 if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(inputUser.Password)); err != nil{
		return c.Status(400).SendString("password or email is incorrect");
	 }

	 token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
		"email": user.Email,
		"role": user.Role,
	})

	jwtKey := []byte(os.Getenv("JWT_SECRET"))

	signedToken, err := token.SignedString(jwtKey)
	if err != nil{
		return err 
	}


	 c.Locals("email", user.Email)
	 return c.JSON(fiber.Map{
        "token": signedToken,
    })

}


func GenerateToken(user User) (string, error){
	


	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
		"email": user.Email,
		"role": user.Role,
	})

	jwtKey := []byte(os.Getenv("JWT_SECRET"))

	signedToken, err := token.SignedString(jwtKey)
	if err != nil{
		return "", err
	}

	return signedToken, nil


	
}

