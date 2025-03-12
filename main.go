package main

import (
	"fmt"
	"github.com/jzero-io/jzero-swagger/action"
	"github.com/urfave/cli/v2"
	"os"
	"runtime"
)

var (
	version  = "v1.1.0"
	commands = []*cli.Command{
		{
			Name:   "swagger",
			Usage:  "generates swagger.json",
			Action: action.Generator,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "host",
					Usage: "api request address",
				},
				&cli.StringFlag{
					Name:  "basepath",
					Usage: "url request prefix",
				},
				&cli.StringFlag{
					Name:  "filename",
					Usage: "swagger save file name",
				},
				&cli.StringFlag{
					Name:  "schemes",
					Usage: "swagger support schemes: http, https, ws, wss",
				},
				&cli.BoolFlag{
					Name:  "route2code",
					Usage: "route2code flag from jzero to generate route code mapping",
				},
			},
		},
	}
)

func main() {
	app := cli.NewApp()
	app.Usage = "a plugin of jzero to generate swagger.json"
	app.Version = fmt.Sprintf("%s %s/%s", version, runtime.GOOS, runtime.GOARCH)
	app.Commands = commands
	if err := app.Run(os.Args); err != nil {
		fmt.Printf("jzero-swagger: %+v\n", err)
	}
}
