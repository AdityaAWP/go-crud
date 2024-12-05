package controller

import (
	"github.com/AdityaAWP/go-crud/database"
	"github.com/AdityaAWP/go-crud/model/entity"
	"github.com/AdityaAWP/go-crud/model/request"
	"github.com/gofiber/fiber/v2"
)

func UserControllerGetAll(c *fiber.Ctx) error {
	var users []entity.User
	result := database.DB.Debug().Find(&users)
	if result.Error != nil {
		panic("Failed to get users!")
	}
	return c.JSON(users)
}

func UserControllerCreate(c *fiber.Ctx) error {
	user := new(request.UserCreateRequest)
	if err := c.BodyParser(user); err != nil {
		return err
	}
	newUser := entity.User{
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
	}

	err := database.DB.Create(&newUser).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Failed to create user!",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User created!",
		"data":    newUser,
	})
}
