package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/Kay-Wilkinson/k8s_go_api/handlers"
)

func setupRoutes(app *fiber.App) {
    app.Get("/api/list", handlers.ListFacts)
    app.Get("/", handlers.ListFactsAsTemplate)
    app.Get("/create-fact", handlers.NewFactView)

    app.Post("/api/create-fact", handlers.CreateFact)
}