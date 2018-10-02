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
	// Version describes the current version of junior.
	Version = "1.0.0"
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
	RegisterFlags(app)

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func makeServer(handler fhttp.RequestHandler) *fhttp.Server {
	return &fhttp.Server{Name: Name, Handler: handler}
}

func run(ctx *cli.Context) error {
	cfg, err := ReadConfig(ctx)
	if err != nil {
		return err
	}

	server := makeServer(cfg.HandleFastHTTP)
	fmt.Printf("Listening on port %s...\n", cfg.Port)
	server.ListenAndServe(":" + cfg.Port)
	return nil
}
