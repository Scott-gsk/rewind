package hooks

import (
	"github.com/gofiber/fiber/v2"
	_ "node-manager/docs"
)

type ExampleAppHooks struct {
}

func (e ExampleAppHooks) PreAppSetup(app *fiber.App) {

}

func (e ExampleAppHooks) PostAppSetup(app *fiber.App) {

}
