package user

import (
	"employee/database"
	userModel "employee/model"
	utils "employee/utils"

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

	password := utils.HashPassword(user.Password)
	user.Password = password

	res := db.Create(&user)
	if res.Error != nil {
		c.Status(http.StatusBadRequest).JSON(res.Error)
		return
	}

	token, err := utils.GenerateToken(*user)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(err)
		return
	}
	c.Status(http.StatusOK).JSON(token)
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

	passwordIsValid, msg := utils.VerifyPassword(user.Password, orgUser.Password)

	if !passwordIsValid {
		c.Status(http.StatusBadRequest).JSON(msg)
		return
	}

	token, err := utils.GenerateToken(orgUser)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(err)
		return
	}
	c.Status(http.StatusOK).JSON(token)
}

func VerifyUser(c *fiber.Ctx) {
	type temp struct {
		Token string
	}

	obj := new(temp)
	if err := c.BodyParser(obj); err != nil {
		c.Status(503).Send(err)
		return
	}

	ubserObj, err := utils.ParseToken(obj.Token)

	if err != nil {
		c.Status(503).Send(err)
		return
	}
	c.Status(http.StatusOK).JSON(ubserObj)
}
