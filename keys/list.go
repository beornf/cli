package keys

import (
	"os"

	"github.com/olekukonko/tablewriter"
	"gopkg.in/errgo.v1"
	"github.com/Scalingo/cli/config"
)

func List() error {
	c := config.ScalingoClient()
	keys, err := c.KeysList()
	if err != nil {
		return errgo.Mask(err)
	}

	t := tablewriter.NewWriter(os.Stdout)
	t.SetColWidth(60)
	t.SetHeader([]string{"Name", "Content"})

	for _, k := range keys {
		t.Append([]string{k.Name, k.Content[0:20] + "..." + k.Content[len(k.Content)-30:]})
	}

	t.Render()
	return nil
}
