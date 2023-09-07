package main

//docker compose run --service-ports web sh
import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/marcostota/trivia/database"
	"github.com/marcostota/trivia/handlers"
)

func main() {
	database.ConnectDb()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	setupRoutes(app)

	app.Static("/", "./public")
	app.Use(handlers.NotFound)
	app.Listen(":3000")
}
