package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var Serve = &cli.Command{
	Name:   "serve",
	Usage:  "Start server.",
	Action: serve,
}

func serve(_ *cli.Context) error {
	fmt.Println("asdf")
	return nil
}
