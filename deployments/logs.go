package deployments

import (
	stdio "io"
	"os"

	"github.com/Scalingo/cli/config"
	"github.com/Scalingo/cli/io"
	"gopkg.in/errgo.v1"
)

func Logs(app, deployment string) error {
	client := config.ScalingoClient()
	deploy, err := client.Deployment(app, deployment)

	if err != nil {
		return errgo.Mask(err, errgo.Any)
	}

	res, err := client.DeploymentLogs(deploy.Links.Output)

	if err != nil {
		return errgo.Mask(err, errgo.Any)
	}

	defer res.Body.Close()

	if res.StatusCode == 404 {
		io.Error("There is no log for this deployment.")
	} else {
		stdio.Copy(os.Stdout, res.Body)
	}
	return nil
}
