package main

import (
	userController "employee/controllers"
	"employee/database"
	model "employee/model"
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("HI")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)
	api := app.Group("/api")
	api.Get("/users", userController.GetUsers)
	api.Post("/regiser", userController.CreateUsers)
	api.Post("/login", userController.LogIn)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")

	if err != nil {
		panic("Failed to connect DB")
	}
	fmt.Println("DB connected")

	database.DBConn.AutoMigrate(&model.User{})
}

func main() {
	app := fiber.New()

	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)
	app.Listen(3000)
}
