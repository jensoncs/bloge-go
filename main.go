package main

import (
	"fmt"
	"github.com/jensoncs/bloge/config"
	"github.com/jensoncs/bloge/logger"
	"github.com/jensoncs/bloge/server"
	"github.com/urfave/cli"
	"os"
)

func main() {
	config.Load()
	logger.SetupLogger()
	app := cli.NewApp()
	app.Name = "bloge"
	app.Version = "0.0.1"
	app.EnableBashCompletion = true
	app.Action = func(c *cli.Context) error {
		server.Startserver()
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:        "app",
			Aliases:     []string{"a"},
			Description: "To get the application name",
			Action: func(c *cli.Context) error {
				fmt.Println(config.AppName())
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}

}
