package main

import (
	"github.com/urfave/cli/v2"
)

var App = &cli.App{
	Name: "boss",
	Usage: "Run several commands at once",
	Commands: []*cli.Command{
		CheckCommand,
	},
}
