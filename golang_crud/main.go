package main

import (
	"fmt"
	"golang_crud/database"
	"golang_crud/models"
	"os"
	"strconv"

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
		result := database.DB.Find(&products)
		if result.Error != nil {
			fmt.Println(result.Error)
			return c.Status(500).SendString("encounter an error")
		}
		fmt.Println(products)
		return c.Render("index", fiber.Map{
			"Data":  products,
			"Title": "Product list",
		}, "layouts/app")
	})
	app.Get("/add", func(c *fiber.Ctx) error {
		return c.Render("add", fiber.Map{
			"Title": "Add Product",
		}, "layouts/app")
	})
	app.Post("/edit/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var product models.Product
		if err := database.DB.First(&product, id).Error; err != nil {
			return c.SendString("product not found")
		}
		price, _ := strconv.ParseFloat(c.FormValue("price"), 64)
		stock, _ := strconv.ParseInt(c.FormValue("price"), 0, 64)
		product.Name = c.FormValue("name")
		product.Price = price
		product.Stock = stock
		if err := database.DB.Save(product).Error; err != nil {
			return c.SendString("error when updating data")
		}
		return c.Redirect("/")
	})
	app.Post("/add", func(c *fiber.Ctx) error {
		product := new(models.Product)
		price, _ := strconv.ParseFloat(c.FormValue("price"), 64)
		stock, _ := strconv.ParseInt(c.FormValue("stock"), 0, 64)
		product.Name = c.FormValue("name")
		product.Price = price
		product.Stock = stock
		if err := database.DB.Save(product).Error; err != nil {
			return c.SendString("error when saving data")
		}
		return c.Redirect("/")
	})
	app.Get("/edit/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var products models.Product
		if err := database.DB.First(&products, id).Error; err != nil {
			return c.SendString("product not found")
		}

		return c.Render("edit", fiber.Map{
			"Data":  products,
			"Title": products.Name,
		}, "layouts/app")
	})
	app.Get("/del/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var products models.Product
		if err := database.DB.First(&products, id).Error; err != nil {
			return c.SendString("product not found")
		}
		database.DB.Delete(&products)
		return c.Redirect("/")
	})
	app.Listen(fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT")))
}
