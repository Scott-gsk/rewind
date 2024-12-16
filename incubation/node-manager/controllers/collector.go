package controllers

import (
	"github.com/WeOps-Lab/rewind/lib/web/server"
	"github.com/gofiber/fiber/v2"
	"node-manager/entity"
	"node-manager/services"
)

type CollectorController struct {
	services.CollectorService
}

// Create
// @Tags Collector
// @Summary 创建采集器
// @Accept json
// @Produce json
// @Param req body entity.CollectorCreateRequest true "entity"
// @Success 200 {object} interface{}
// @Router /node-manager/internal/collector/create [post]
func (ctrl CollectorController) Create(c *fiber.Ctx) error {
	req, err := server.ParseAndValidateRequest[entity.CollectorCreateRequest](c)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := ctrl.CollectorService.Create(*req); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create collector")
	}

	return c.Status(fiber.StatusOK).JSON("Collector created successfully")
}

// List
// @Tags Collector
// @Summary 获取采集器列表
// @Accept json
// @Produce json
// @Success 200 {object} []entity.CollectorItemResponse
// @Router /node-manager/internal/collector/list [get]
func (ctrl CollectorController) List(c *fiber.Ctx) error {
	collectorItemResponses, err := ctrl.CollectorService.List()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to list collectors")
	}

	return c.Status(fiber.StatusOK).JSON(collectorItemResponses)
}
