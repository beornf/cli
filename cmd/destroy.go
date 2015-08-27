package cmd

import (
	"github.com/Scalingo/cli/Godeps/_workspace/src/github.com/Scalingo/codegangsta-cli"
	"github.com/Scalingo/cli/appdetect"
	"github.com/Scalingo/cli/apps"
	"github.com/Scalingo/cli/cmd/autocomplete"
)

var (
	DestroyCommand = cli.Command{
		Name:        "destroy",
		Category:    "Global",
		Flags:       []cli.Flag{appFlag},
		Usage:       "Destroy an app /!\\",
		Description: "Destroy an app /!\\ It is not reversible\n  Example:\n    'scalingo destroy my-app'",
		Action: func(c *cli.Context) {
			currentApp := appdetect.CurrentApp(c)
			if len(c.Args()) != 0 {
				cli.ShowCommandHelp(c, "destroy")
			} else {
				err := apps.Destroy(currentApp)
				if err != nil {
					errorQuit(err)
				}
			}
		},
		BashComplete: func(c *cli.Context) {
			autocomplete.CmdFlagsAutoComplete(c, "destroy")
		},
	}
)
