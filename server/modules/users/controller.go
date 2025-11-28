package users

import "github.com/gofiber/fiber/v2"

type UserController struct {
    service *UserService
}

func NewUserController(s *UserService) *UserController {
    return &UserController{service: s}
}

func (ctrl *UserController) Register(router fiber.Router) {
    r := router.Group("/users")

    r.Get("/", func(c *fiber.Ctx) error {
        return c.JSON(ctrl.service.GetUsers())
    })
}
