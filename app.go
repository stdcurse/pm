package main

import (
	"github.com/stdcurse/pm/commands/install"
	"github.com/stdcurse/pm/commands/pull"
	"github.com/stdcurse/pm/output"
	"github.com/urfave/cli/v2"
	"os"
)

func NewApp() *cli.App {
	cli.HelpFlag = &cli.BoolFlag{}
	cli.AppHelpTemplate = `{{.Name}} - {{.Usage}}

Usage: {{.Name}} [ACTION] [OPTION...]

Subcommands:
{{range .Commands}}{{if not .HideHelp}}   {{join .Names ", "}}{{ "\t\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}
Options:
{{range .VisibleFlags}}   {{.}}{{end}}

Bug tracker: https://github.com/stdcurse/pm/issues
`

	app := &cli.App{
		Name:  "pm",
		Usage: "Package manager for stdcurse Linux distribution",
		Action: func(c *cli.Context) error {
			output.Info("Use \"help\" subcommand to get help")

			os.Exit(1)
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "config",
				Value: "/etc/pm/config.yaml",
				Usage: "Select custom path to config",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "version",
				Usage: "Get pm's version",
				Action: func(c *cli.Context) error {
					output.Info("9999")

					return nil
				},
			},
			{
				Name:    "pull",
				Aliases: []string{"p"},
				Usage:   "Pull packages build files from remote",
				Action:  pull.Command,
			},
			{
				Name:    "install",
				Aliases: []string{"i"},
				Usage:   "Install package(s)",
				Action:  install.Command,
			},
		},
	}

	return app
}
