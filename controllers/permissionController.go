package controllers

import (
  "github.com/volo-h/minors/database"
  "github.com/volo-h/minors/models"

  "github.com/gofiber/fiber/v2"
)

func AllPermissions(c *fiber.Ctx) error {
  var permissions []models.Permission

  database.DB.Find(&permissions)

  return c.JSON(permissions)
}
