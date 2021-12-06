/*
	Copyright (c) 2021 Nikita Nikiforov <vokestd@gmail.com>

	This software is provided 'as-is', without any express or implied
	warranty. In no event will the authors be held liable for any damages
	arising from the use of this software.

	Permission is granted to anyone to use this software for any purpose,
	including commercial applications, and to alter it and redistribute it
	freely, subject to the following restrictions:

	1. The origin of this software must not be misrepresented; you must not
		 claim that you wrote the original software. If you use this software
		 in a product, an acknowledgement in the product documentation would be
		 appreciated but is not required.
	2. Altered source versions must be plainly marked as such, and must not be
		 misrepresented as being the original software.
	3. This notice may not be removed or altered from any source distribution.
*/

package main

import (
	"github.com/stdcurse/pm/commands/emerge"
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
				Name:    "emerge",
				Aliases: []string{"e"},
				Usage:   "Emerge package(s)",
				Action:  emerge.Command,
			},
		},
	}

	return app
}
