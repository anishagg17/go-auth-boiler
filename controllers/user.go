package user

import (
	"employee/database"
	userModel "employee/model"
	hashPassword "employee/utils"

	"net/http"

	"github.com/gofiber/fiber"
)

func GetUsers(c *fiber.Ctx) {
	db := database.DBConn

	var users []userModel.User
	db.Find(&users)
	c.Status(http.StatusOK).JSON(users)
}

func CreateUsers(c *fiber.Ctx) {
	db := database.DBConn
	user := new(userModel.User)

	if err := c.BodyParser(user); err != nil {
		c.Status(503).Send(err)
		return
	}

	password := hashPassword.HashPassword(user.Password)
	user.Password = password

	res := db.Create(&user)
	if res.Error != nil {
		c.Status(http.StatusBadRequest).JSON(res.Error)
		return
	}

	// TODO provide JWT token
	c.Status(http.StatusOK).JSON(user)
}

func LogIn(c *fiber.Ctx) {
	type auth struct {
		Email    string
		Password string
	}

	db := database.DBConn
	user := new(auth)

	if err := c.BodyParser(user); err != nil {
		c.Status(503).Send(err)
		return
	}

	var orgUser userModel.User
	db.Where(userModel.User{Email: user.Email}).Find(&orgUser)

	passwordIsValid, msg := hashPassword.VerifyPassword(user.Password, orgUser.Password)

	if !passwordIsValid {
		c.Status(http.StatusBadRequest).JSON(msg)
		return
	}

	// TODO provide JWT token
	c.Status(http.StatusOK).JSON(orgUser)
}
