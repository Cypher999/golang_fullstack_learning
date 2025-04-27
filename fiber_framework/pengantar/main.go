package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type IncomingData struct {
	Nama   string `form:"nama" validate:"required,min=3,max=20"`
	Alamat string `form:"alamat" validate:"required,min=3,max=20"`
}

var validate = validator.New()

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: 20 * 1024 * 1024,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("this is home function")
	})

	app.Post("/", func(c *fiber.Ctx) error {
		var fileName string
		var fileSize int64
		var fileExtension string
		var fileMime string
		req := IncomingData{
			Nama:   c.FormValue("nama"),
			Alamat: c.FormValue("alamat"),
		}
		err := validate.Struct(req)
		if err != nil {
			errorMessage := ""
			for _, err := range err.(validator.ValidationErrors) {
				errorMessage += fmt.Sprintf("%s: %s %s", err.Namespace(), err.Tag(), err.Param())
			}
			return errors.New(errorMessage)
		}

		fileHeader, err := c.FormFile("picture")

		if fileHeader != nil {
			file, err := fileHeader.Open()
			mime, err := mimetype.DetectReader(file)
			fileMime = mime.String()
			fileName = fileHeader.Filename
			fileSize = fileHeader.Size
			fileExtension = strings.Split(fileName, ".")[1]
			if err != nil {
				return c.Status(400).SendString("upload gagal")
			}

			// (Optional) Save file if needed
			err = c.SaveFile(fileHeader, fmt.Sprintf("./uploads/%s", fileHeader.Filename))
			if err != nil {
				return c.Status(500).SendString("upload gagal")
			}
		}

		hasil := fmt.Sprintf("Nama :%s\n Alamat:%s\nNama File:%s\n Ukuran:%d\nTipe:%s\n Ekstensi:%s", req.Nama, req.Alamat, fileName, fileSize, fileMime, fileExtension)
		return c.SendString(hasil)
	})

	app.Listen(":8080")
}
