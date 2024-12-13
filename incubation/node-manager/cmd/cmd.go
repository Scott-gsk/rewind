package cmd

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/urfave/cli/v2"
	"node-manager/cmd/db"
	"node-manager/cmd/server"
	_ "node-manager/models"
	"os"
)

func Run() {
	app := &cli.App{
		Name: "app",
		Commands: []*cli.Command{
			server.Command(),
			db.Command(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
