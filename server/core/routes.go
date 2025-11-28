package core

import (
    "github.com/gofiber/fiber/v2"
    "plintes/modules/users"
)

func RegisterRoutes(server *FiberServer, userCtrl *users.UserController) {
    api := server.App.Group("/api")

    api.Get("/", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{"status": "Plintes API running"})
    })

    userCtrl.Register(api)
}
