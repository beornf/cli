package apps

import (
	"fmt"

	"github.com/Scalingo/cli/api"
	"gopkg.in/errgo.v1"
)

func Restart(app string, args []string) error {
	params := api.AppsRestartParams{args}

	res, err := api.AppsRestart(app, &params)
	if err != nil {
		return errgo.Mask(err, errgo.Any)
	}
	res.Body.Close()

	fmt.Printf("You application is being restarted\n")

	return nil
}
