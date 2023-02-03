package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/template/html"
    "github.com/Kay-Wilkinson/k8s_go_api/database"
)

func main() {
    database.ConnectDb()
    engine := html.New("./views", ".html")

    app := fiber.New(fiber.Config{
        Views: engine,
        ViewsLayout: "layouts/main", //this is to set global templates
    })

    setupRoutes(app)

    app.Listen(":3000")
}
