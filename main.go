package main

import (
	"os"

	"github.com/Pilfer/ultimate-guitar-scraper/cmd"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Before = cmd.BeforeCommand
	app.After = cmd.AfterCommand
	app.Name = "ug"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		cmd.FetchTab,
		cmd.SearchTab,
	}
	_ = app.Run(os.Args)

}
