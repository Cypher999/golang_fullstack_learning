package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type Staff struct {
	Nama    string
	Jabatan string
}

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", nil, "nested_layout/admin", "nested_layout/app")
	})
	app.Get("/about", func(c *fiber.Ctx) error {
		data := []Staff{
			{Nama: "Sandi", Jabatan: "Programer"},
			{Nama: "Rifki", Jabatan: "Owner"},
			{Nama: "Deja", Jabatan: "Desainer Grafis"},
		}
		return c.Render("about", fiber.Map{
			"title": "About Page",
			"data":  data,
		}, "layout/app")
	})

	app.Listen(":8000")
}
