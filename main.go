package main

import (
	"os"

	"github.com/Bralimus/save_inspector/app"
	"github.com/Bralimus/save_inspector/cli"
)

func main() {
	app := app.NewApp()
	cli.Execute(app, os.Args)
}
