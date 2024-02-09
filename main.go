package main

import (
	"log"
	"time"

	"showunifier/pkg/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})

	app.Get("/", func(c *fiber.Ctx) error {
		serie := domain.Serie {
			Name: "serie básica",
			InclussionDate: time.Now(),
			InterestLevel: 2,
		}
		return c.Render("index", fiber.Map{
			"Serie": serie,
		})
	})
	log.Fatal(app.Listen(":3000"))
}
