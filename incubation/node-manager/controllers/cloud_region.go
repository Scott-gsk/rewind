package controllers

import (
	"github.com/WeOps-Lab/rewind/lib/web/server"
	"github.com/gofiber/fiber/v2"
	"node-manager/entity"
	"node-manager/services"
)

type CloudRegionController struct {
	CloudRegionService services.CloudRegionService
}

// List
// @Tags CloudRegion
// @Summary 获取所有云区域
// @Router /node-manager/internal/cloud-region/list [get]
// @Accept json
// @Produce json
// @Success 200 {object} []entity.CloudRegionItemResponse
func (ctrl CloudRegionController) List(c *fiber.Ctx) error {
	cloudRegions, err := ctrl.CloudRegionService.List()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch cloud regions")
	}
	return c.Status(fiber.StatusOK).JSON(cloudRegions)
}

// Create
// @Tags CloudRegion
// @Summary 创建云区域
// @Param req body entity.CloudRegionCreateRequest true "entity"
// @Router /node-manager/internal/cloud-region/create [post]
// @Accept json
// @Produce json
// @Success 200 {object} interface{}
func (ctrl CloudRegionController) Create(c *fiber.Ctx) error {
	req, err := server.ParseAndValidateRequest[entity.CloudRegionCreateRequest](c)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	err = ctrl.CloudRegionService.Create(*req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create cloud region")
	}
	return c.Status(fiber.StatusOK).JSON("Cloud region created successfully")
}

// Update
// @Tags CloudRegion
// @Summary 更新云区域
// @Param req body entity.CloudRegionUpdateRequest true "entity"
// @Router /node-manager/internal/cloud-region/update [put]
// @Accept json
// @Produce json
// @Success 200 {object} interface{}
func (ctrl CloudRegionController) Update(c *fiber.Ctx) error {
	req, err := server.ParseAndValidateRequest[entity.CloudRegionUpdateRequest](c)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	err = ctrl.CloudRegionService.Update(*req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update cloud region")
	}
	return c.Status(fiber.StatusOK).JSON("Cloud region updated successfully")
}
