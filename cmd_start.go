package main

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"

	"github.com/mjm/boss/api"
	"github.com/mjm/boss/cfg"
)

var StartCommand = &cli.Command{
	Name:   "start",
	Usage:  "start one or more tasks",
	Action: runStartCommand,
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
}

func runStartCommand(c *cli.Context) error {
	cfgFile := c.Path("config-file")
	projectCfg, err := cfg.ReadFile(cfgFile)
	if err != nil {
		return err
	}

	socketName := fmt.Sprintf("/tmp/boss.%s.sock", projectCfg.Project.Name)
	conn, err := grpc.Dial("unix:"+socketName, grpc.WithInsecure())
	if err != nil {
		return err
	}

	client := api.NewBossClient(conn)
	_, err = client.StartTasks(context.Background(), &api.StartTasksRequest{
		Names: c.Args().Slice(),
	})
	if err != nil {
		return err
	}

	return nil
}
