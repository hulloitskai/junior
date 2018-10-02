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

const (
	// Version describes the version of command junior.
	Version = "0.1.0"
	// Name is the name of this server.
	Name = "junior"
)

func main() {
	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Usage = "a fast, tiny HTTP server for serving static content"
	app.UsageText = fmt.Sprintf("%s [global options]", Name)
	app.Action = run

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func makeServer(handler fhttp.RequestHandler) *fhttp.Server {
	return &fhttp.Server{Name: Name, Handler: handler}
}

func run(ctx *cli.Context) error {
	cfg, err := ReadConfig()
	if err != nil {
		return err
	}

	server := makeServer(cfg.HandleFastHTTP)
	server.ListenAndServe(":" + cfg.Port)
	return nil
}
