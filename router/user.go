package router

import (
	db "edith/db"
	"edith/models"
	"edith/utils"
	"math/rand"
	"os"

	"golang.org/x/crypto/bcrypt"

	"github.com/gofiber/fiber/v2"
)

var jwtKey = []byte(os.Getenv("PRIV_KEY"))

// SetupUserRoutes sets up all the user routes
func SetupUserRoutes() {
	USER.Post("/signup", CreateUser) //creates a new user
}

func CreateUser(c *fiber.Ctx) error {
	u := new(models.User)

	if err := c.BodyParser(u); err != nil {
		return c.JSON(fiber.Map{
			"error": true,
			"input": "Incorrect data submitted!",
		})
	}

	// Validate if the email, Username and password are in correct format

	errors := utils.ValidateRegister(u)
	if errors.Err {
		return c.JSON(errors)
	}

	if count := db.DB.Where(&models.User{Email: u.Email}).First(new(models.User)).RowsAffected; count > 0 {
		errors.Err, errors.Email = true, "Email is already registered"
	}

	if count := db.DB.Where(&models.User{UserName: u.UserName}).First(new(models.User)).RowsAffected; count > 0 {
		errors.Err, errors.Email = true, "Username is already registered !"
	}
	if errors.Err {
		return c.JSON(errors)
	}

	// Hashing the password with a random salt
	password := []byte(u.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(
		password,
		rand.Intn(bcrypt.MaxCost-bcrypt.MinCost),
	)

	if err != nil {
		panic(err)
	}
	u.Password = string(hashedPassword)

	if err := db.DB.Create(&u).Error; err != nil {
		return c.JSON(fiber.Map{
			"error":   true,
			"general": "Something went wrong, please try again later.",
		})
	}

	// setting up the auth cookies
	accessToken, refreshToken := utils.GenerateTokens(u.UUID.String())
	accessCookie, refreshCookie := utils.GetAuthCookies(accessToken, refreshToken)
	c.Cookie(accessCookie)
	c.Cookie(refreshCookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})

}
