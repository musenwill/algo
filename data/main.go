package main

import (
	"os"

	"github.com/musenwill/algo/data/numbers"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.ErrWriter = os.Stderr
	app.EnableBashCompletion = true
	app.Name = "gen"
	app.Usage = "Generate test data randomly"
	app.Description = ""
	app.Commands = []cli.Command{numbers.GenNumberCmd}

	app.RunAndExitOnError()
}
