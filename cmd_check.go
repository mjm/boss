package main

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/mjm/boss/cfg"
)

var CheckCommand = &cli.Command{
	Name:   "check",
	Usage:  "check project file for issues",
	Action: runCheckCommand,
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

func runCheckCommand(c *cli.Context) error {
	cfgFile := c.Path("config-file")
	projectCfg, err := cfg.ReadFile(cfgFile)
	if err != nil {
		return err
	}

	fmt.Printf("no issues found in project %q\n", projectCfg.Project.Name)
	return nil
}
