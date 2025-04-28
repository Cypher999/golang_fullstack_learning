package main

import (
	"fmt"
	"golang_crud/database"
	"golang_crud/models"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	engine := html.New("./views", ".html")
	if err != nil {
		panic("error, no .env found")
	}
	database.Connection()
	database.DB.AutoMigrate(&models.Product{})
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		var products []models.Product
		database.DB.Find(&products)
		fmt.Println(products)
		return c.Render("index", fiber.Map{
			"Data":  products,
			"Title": "Product list",
		}, "layouts/app")
	})
	app.Listen(fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT")))
}
