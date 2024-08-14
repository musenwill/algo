package main

import (
	"fmt"
	"os"

	"github.com/musenwill/algo/cmd"
	_ "github.com/musenwill/algo/src"
)

func main() {
	if err := cmd.New().Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
