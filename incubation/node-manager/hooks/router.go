package hooks

import (
	"github.com/gofiber/fiber/v2"
	"node-manager/controllers"
)

func (e ExampleAppHooks) InstallInternalRouter(router fiber.Router) {
	variableGroup := router.Group("/variable")
	variableController := controllers.VariableController{}

	variableGroup.Get("/list/:cloudRegionId", variableController.ListByCloudRegion)
	variableGroup.Get("/search/:cloudRegionId", variableController.SearchByCloudRegion)
	variableGroup.Post("/create", variableController.CreateVariable)
	variableGroup.Put("/update", variableController.UpdateVariable)
	variableGroup.Delete("/delete/single/:id", variableController.DeleteSingleVariable)
	variableGroup.Delete("/delete/multiple", variableController.DeleteMultipleVariables)

	cloudRegionGroup := router.Group("/cloud-region")
	cloudRegionController := controllers.CloudRegionController{}

	cloudRegionGroup.Get("/list", cloudRegionController.List)
	cloudRegionGroup.Post("/create", cloudRegionController.Create)
	cloudRegionGroup.Put("/update", cloudRegionController.Update)

	collectorGroup := router.Group("/collector")
	collectorController := controllers.CollectorController{}

	collectorGroup.Get("/list", collectorController.List)
	collectorGroup.Post("/create", collectorController.Create)
}

func (e ExampleAppHooks) InstallPublicRouter(router fiber.Router) {
}

func (e ExampleAppHooks) InstallMeshRouter(router fiber.Router) {}
