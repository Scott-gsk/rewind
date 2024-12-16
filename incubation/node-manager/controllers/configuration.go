package controllers

import (
	"github.com/WeOps-Lab/rewind/lib/web/server"
	"github.com/gofiber/fiber/v2"
	"node-manager/entity"
	"node-manager/services"
)

type ConfigurationController struct {
	services.ConfigurationService
}

// Create
// @Tags Configuration
// @Summary Create a new configuration
// @Accept json
// @Produce json
// @Param req body entity.ConfigurationCreateRequest true "entity"
// @Success 200 {object} interface{}
// @Router /node-manager/internal/configuration/create [post]
func (ctrl ConfigurationController) Create(c *fiber.Ctx) error {
	req, err := server.ParseAndValidateRequest[entity.ConfigurationCreateRequest](c)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := ctrl.ConfigurationService.Create(*req); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create configuration")
	}

	return c.Status(fiber.StatusOK).JSON("Configuration created successfully")
}

// List
// @Tags Configuration
// @Summary Get a list of configurations
// @Accept json
// @Produce json
// @Success 200 {object} []entity.ConfigurationItemResponse
// @Router /node-manager/internal/configuration/list [get]
func (ctrl ConfigurationController) List(c *fiber.Ctx) error {
	configurationItemResponses, err := ctrl.ConfigurationService.List()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to list configurations")
	}

	return c.Status(fiber.StatusOK).JSON(configurationItemResponses)
}

// Update
// @Tags Configuration
// @Summary Update an existing configuration
// @Accept json
// @Produce json
// @Param req body entity.ConfigurationUpdateRequest true "entity"
// @Success 200 {object} interface{}
// @Router /node-manager/internal/configuration/update [put]
func (ctrl ConfigurationController) Update(c *fiber.Ctx) error {
	req, err := server.ParseAndValidateRequest[entity.ConfigurationUpdateRequest](c)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := ctrl.ConfigurationService.Update(*req); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update configuration")
	}

	return c.Status(fiber.StatusOK).JSON("Configuration updated successfully")
}

// Delete
// @Tags Configuration
// @Summary Delete a configuration
// @Accept json
// @Produce json
// @Param id path uint true "Configuration ID"
// @Success 200 {object} interface{}
// @Router /node-manager/internal/configuration/delete/{id} [delete]
func (ctrl ConfigurationController) Delete(c *fiber.Ctx) error {
	id, err := server.ParseAndValidateRequest[uint](c)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := ctrl.ConfigurationService.Delete(*id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete configuration")
	}

	return c.Status(fiber.StatusOK).JSON("Configuration deleted successfully")
}

// DeleteMultiple
// @Tags Configuration
// @Summary Delete multiple configurations
// @Accept json
// @Produce json
// @Param ids body []uint true "Configuration IDs"
// @Success 200 {object} interface{}
// @Router /node-manager/internal/configuration/delete-multiple [delete]
func (ctrl ConfigurationController) DeleteMultiple(c *fiber.Ctx) error {
	var ids []uint
	if err := c.BodyParser(&ids); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := ctrl.ConfigurationService.DeleteMultiple(ids); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete configurations")
	}

	return c.Status(fiber.StatusOK).JSON("Configurations deleted successfully")
}
