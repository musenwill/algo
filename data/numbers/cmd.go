package numbers

import "github.com/urfave/cli"

var GenNumberCmd = cli.Command{
	Name:   "number",
	Usage:  "generate numbers",
	Action: run,
	Flags:  []cli.Flag{countFlag, fileFlag},
}

func run(c *cli.Context) error {
	file := c.String(fileFlag.Name)
	count := c.Int(countFlag.Name)

	gen := NewGenerator(count, file)
	return gen.GenNumbers()
}

var countFlag = cli.IntFlag{
	Name:     "count",
	Usage:    "count",
	Required: true,
}

var fileFlag = cli.StringFlag{
	Name:     "file",
	Usage:    "file",
	Required: true,
}
