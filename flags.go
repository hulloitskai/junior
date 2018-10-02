package main

import (
	"gopkg.in/urfave/cli.v1"
)

// RegisterFlags confgiures app.flags.
func RegisterFlags(app *cli.App) {
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "port, p",
			Usage:  "port to serve on",
			Value:  "80",
			EnvVar: "PORT",
		},
		cli.StringFlag{
			Name:   "root, r",
			Usage:  "root directory to serve from",
			EnvVar: "ROOT_DIR",
		},
		cli.StringFlag{
			Name:   "not-found, n",
			Usage:  "location to serve during 404",
			Value:  "404.html",
			EnvVar: "NOT_FOUND",
		},
		cli.StringFlag{
			Name: "trailing-slash, t",
			Usage: "\"true\" enforces trailing slash, \"false\" enforces no " +
				"trailing slash; otherwise, trailing slash is left as-is",
			Value:  "off",
			EnvVar: "TRAILING_SLASH",
		},
	}
}
