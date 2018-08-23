package cmd

import (
	"fmt"
	"github.com/verystar/logger"
	"github.com/urfave/cli"
)

var TestCmd = cli.Command{
	Name:  "test",
	Usage: "test command eg: ./app test --id=7",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "id",
			Usage: "user id",
		},
	},
	Action: func(ctx *cli.Context) error {
		if !ctx.IsSet("id") {
			ctx.Set("id", "7")
		}


		fmt.Println("111")

		logger.Error("test", "user")
		return nil
	},
}

func init() {
	Register(TestCmd)
}
