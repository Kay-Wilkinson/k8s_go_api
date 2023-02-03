package handlers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/Kay-Wilkinson/k8s_go_api/database"
    "github.com/Kay-Wilkinson/k8s_go_api/models"
)

func ListFacts(c *fiber.Ctx) error {
    facts := []models.Fact{}
    database.DB.Db.Find(&facts)

    return c.Status(200).JSON(facts)
}

func ListFactsAsTemplate(c *fiber.Ctx) error {
    facts := []models.Fact{}
    database.DB.Db.Find(&facts)

    return c.Render("index", fiber.Map{
        "Title": "Demo Go/Fiber Trivia App",
        "Subtitle": "This is just a test build",
        "Facts": facts, //send fact data from db to view
    })
}

func NewFactView(c *fiber.Ctx) error {
    return c.Render("new", fiber.Map{
        "Title":    "New Fact",
        "Subtitle": "Add a cool fact",
    })
}

func CreateFact(c *fiber.Ctx) error {
    fact := new(models.Fact)
    if err := c.BodyParser(fact); err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": err.Error(),
        })
    }

    database.DB.Db.Create(&fact)

//     return c.Status(200).JSON(fact)
    return ConfirmationView(c) // return redirect to confirmation view
}

func ConfirmationView(c *fiber.Ctx) error {
    return c.Render("confirmation", fiber.Map{
        "Title":    "Fact added successfully",
        "Subtitle": "Add more facts to the list or click 'Return to Homepage' to list those facts",
    })
}