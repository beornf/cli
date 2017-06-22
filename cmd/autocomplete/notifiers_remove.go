package autocomplete

import (
	"fmt"

	"github.com/Scalingo/cli/config"
	"github.com/Scalingo/codegangsta-cli"
)

func NotifiersRemoveAutoComplete(c *cli.Context) error {
	appName := CurrentAppCompletion(c)
	if appName == "" {
		return nil
	}

	client := config.ScalingoClient()
	resources, err := client.NotifiersList(appName)
	if err == nil {
		for _, resource := range resources {
			fmt.Println(resource.GetID())
		}
	}

	return nil
}