package main

import (
	"log"

	"github.com/cetinboran/basicsec/database"
	"github.com/cetinboran/basicsec/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer database.DBConn.Close()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/static", "./views")
	routers.SetRouters(app)

	app.Listen(":3000")
}
