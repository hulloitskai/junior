package main

import (
	"fmt"
	"log"
	"os"

	fhttp "github.com/valyala/fasthttp"
	"gopkg.in/urfave/cli.v1"
)

func init() {
	log.SetFlags(0)
}

// Version describes the version of command junior.
const Version = "0.1.0"

func main() {
	app := cli.NewApp()
	app.Name = "junior"
	app.Version = Version
	app.Usage = "a fast, tiny HTTP server for serving static content"
	app.UsageText = fmt.Sprintf("%s [global options]", app.Name)
	app.Action = run

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(ctx *cli.Context) error {
	cfg, err := ReadConfig()
	if err != nil {
		return err
	}

	fhttp.ListenAndServe(":"+cfg.Port, cfg.HandleFastHTTP)
	return nil
}
