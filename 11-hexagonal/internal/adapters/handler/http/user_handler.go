package http

import (
	"strconv"

	"github.com/Nextjingjing/go-god/11-hexagonal/internal/adapters/handler/http/dto"
	"github.com/Nextjingjing/go-god/11-hexagonal/internal/core/ports"
	"github.com/gofiber/fiber/v3"
)

type userHandler struct {
	userService ports.UserService
}

func NewUserHandler(userService ports.UserService) *userHandler {
	return &userHandler{userService: userService}
}

func (h *userHandler) UserRoute(routeGroup fiber.Router) {
	routeGroup.Post("/", h.CreateUser)
	routeGroup.Get("/:id", h.GetUserByID)
	routeGroup.Get("/", h.GetAllUsers)
	routeGroup.Put("/:id", h.UpdateUser)
}

func (h *userHandler) CreateUser(c fiber.Ctx) error {
	var req dto.UserRequestDTO
	c.Bind().Body(&req)
	user, err := h.userService.CreateUser(req.Name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *userHandler) GetUserByID(c fiber.Ctx) error {
	idStr := c.Params("id")
	temp, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}
	id := uint(temp)
	user, err := h.userService.GetUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(user)
}

func (h *userHandler) GetAllUsers(c fiber.Ctx) error {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}
func (h *userHandler) UpdateUser(c fiber.Ctx) error {
	idStr := c.Params("id")
	temp, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}
	id := uint(temp)
	var req dto.UserRequestDTO
	c.Bind().Body(&req)
	user, err := h.userService.UpdateUser(id, req.Name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(user)
}
