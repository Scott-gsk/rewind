package controllers

import (
	"github.com/WeOps-Lab/rewind/lib/web/server"
	"github.com/gofiber/fiber/v2"
	"node-manager/entity"
	"node-manager/services"
)

type VariableController struct {
	VariableService services.VariableService
}

// @Tags Variable
// @Summary 获取该云区域下的所有变量
// @Router /node-manager/internal/variable/list/{cloudRegionId} [get]
// @Accept json
// @Produce json
// @Param cloudRegionId path uint true "Cloud Region ID"
// @Success 200 {object} []entity.VariableItemResponse
func (ctrl VariableController) ListByCloudRegion(c *fiber.Ctx) error {
	cloudRegionId, err := c.ParamsInt("cloudRegionId")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid Cloud Region ID")
	}
	variables, err := ctrl.VariableService.ListByCloudRegion(uint(cloudRegionId))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch variables")
	}
	return c.Status(fiber.StatusOK).JSON(variables)
}

// @Tags Variable
// @Summary 搜索该云区域下的变量
// @Param cloudRegionId path uint true "Cloud Region ID"
// @Param value query string true "Search Value"
// @Router /node-manager/internal/variable/search/{cloudRegionId} [get]
// @Accept json
// @Produce json
// @Success 200 {object} []entity.VariableItemResponse
func (ctrl VariableController) SearchByCloudRegion(c *fiber.Ctx) error {
	cloudRegionId, err := c.ParamsInt("cloudRegionId")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid Cloud Region ID")
	}
	value := c.Query("value")
	variables, err := ctrl.VariableService.SearchByCloudRegion(uint(cloudRegionId), value)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch variables")
	}
	return c.Status(fiber.StatusOK).JSON(variables)
}

// @Tags Variable
// @Summary 创建变量
// @Param req body entity.VariableCreateRequest true "entity"
// @Router /node-manager/internal/variable/create [post]
// @Accept json
// @Produce json
// @Success 200 {object} interface{}
func (ctrl VariableController) CreateVariable(c *fiber.Ctx) error {
	req, err := server.ParseAndValidateRequest[entity.VariableCreateRequest](c)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}
	err = ctrl.VariableService.CreateVariable(*req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create variable")
	}
	return c.Status(fiber.StatusOK).JSON("Variable created successfully")
}

// @Tags Variable
// @Summary 修改变量
// @Param req body entity.VariableUpdateRequest true "entity"
// @Router /node-manager/internal/variable/update [put]
// @Accept json
// @Produce json
// @Success 200 {object} interface{}
func (ctrl VariableController) UpdateVariable(c *fiber.Ctx) error {
	req, err := server.ParseAndValidateRequest[entity.VariableUpdateRequest](c)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}
	err = ctrl.VariableService.UpdateVariable(*req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update variable")
	}
	return c.Status(fiber.StatusOK).JSON("Variable updated successfully")
}

// @Tags Variable
// @Summary 删除单个变量
// @Param id path uint true "Variable ID"
// @Router /node-manager/internal/variable/delete/single/{id} [delete]
// @Accept json
// @Produce json
// @Success 200 {object} interface{}
func (ctrl VariableController) DeleteSingleVariable(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid ID")
	}
	err = ctrl.VariableService.DeleteSingleVariable(uint(id))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete variable")
	}
	return c.Status(fiber.StatusOK).JSON("Variable deleted successfully")
}

// @Tags Variable
// @Summary 删除多个变量
// @Param ids body []uint true "Variable IDs"
// @Router /node-manager/internal/variable/delete/multiple [delete]
// @Accept json
// @Produce json
// @Success 200 {object} interface{}
func (ctrl VariableController) DeleteMultipleVariables(c *fiber.Ctx) error {
	var ids []uint
	if err := c.BodyParser(&ids); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}
	err := ctrl.VariableService.DeleteMultipleVariables(ids)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete variables")
	}
	return c.Status(fiber.StatusOK).JSON("Variables deleted successfully")
}