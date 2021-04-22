package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/urfave/cli/v2"

	"github.com/mjm/boss/cfg"
	"github.com/mjm/boss/server"
)

var App = &cli.App{
	Name:  "boss",
	Usage: "Run several commands at once",
	Flags: []cli.Flag{
		&cli.PathFlag{
			Name:      "config-file",
			Aliases:   []string{"f"},
			Usage:     "check project config in `FILE`",
			Required:  false,
			TakesFile: true,
			Value:     cfg.DefaultProjectFile,
		},
	},
	Commands: []*cli.Command{
		CheckCommand,
		StartCommand,
	},
	Action: runApp,
}

func runApp(c *cli.Context) error {
	cfgFile := c.Path("config-file")
	projectCfg, err := cfg.ReadFile(cfgFile)
	if err != nil {
		return err
	}

	srv := &server.Server{
		Project:        projectCfg.Project,
		ConfigFilePath: cfgFile,
	}

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		<-c
		srv.Shutdown()
	}()

	return srv.Run()
}
