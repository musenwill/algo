package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli"
)

var (
	// Version should be a git tag
	Commit  = ""
	Version = ""
	Branch  = ""
	Name    = ""
)

var commands []cli.Command

func New() *cli.App {
	app := cli.NewApp()
	app.ErrWriter = os.Stderr
	app.EnableBashCompletion = true
	app.Name = Name
	app.Usage = "An algorithm suit."
	app.Version = Version + "(git:" + Branch + " " + Commit + ")"
	app.Author = "musenwill"
	app.Email = "musenwill@qq.com"
	app.Copyright = fmt.Sprintf("Copyright © 2024 - %v musenwill. All Rights Reserved.", time.Now().Year())
	app.Description = description
	app.Commands = commands
	return app
}

func RegisterCommands(cmds ...cli.Command) {
	commands = append(commands, cmds...)
}
