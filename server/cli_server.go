package server

import (
	"os"
	"context"
	"app/config"
	"app/cmd"
	"github.com/urfave/cli"
	"github.com/verystar/golib/osutil"
)

var (
	VERSION = "unknown"
)

type CliServer struct {
	serv *cli.App
}

func NewCliServer() *CliServer {
	return &CliServer{
		serv: cli.NewApp(),
	}
}

func (h *CliServer) Serv() *cli.App {
	return h.serv
}

func (h *CliServer) Run() {
	app := h.serv
	app.Name = "app"
	app.Version = VERSION
	app.Copyright = "(c) VeryStar"
	app.Writer = os.Stdout
	cli.ErrWriter = os.Stdout

	app.Commands = cmd.Commands
	app.Setup()

	ctx, cancel := context.WithCancel(context.Background())
	app.Metadata["ctx"] = ctx
	var args []string
	args = append(args, app.Name)
	args = append(args, config.Args...)

	osutil.RegisterShutDown(func(sig os.Signal) {
		cancel()
	})

	go osutil.WaitSignal()
	app.Run(args)
}
