package cmd

import (
	"github.com/urfave/cli"
	"context"
)

var Commands [] cli.Command

func Register(cmd cli.Command) {
	Commands = append(Commands, cmd)
}

func getContext(c *cli.Context) context.Context {
	return c.App.Metadata["ctx"].(context.Context)
}
