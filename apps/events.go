package apps

import (
	"github.com/Scalingo/cli/config"
	"github.com/Scalingo/cli/events"
	"github.com/Scalingo/go-scalingo"
	"gopkg.in/errgo.v1"
)

func Events(app string, paginationOpts scalingo.PaginationOpts) error {
	c := config.ScalingoClient()
	appEvents, pagination, err := c.EventsList(app, paginationOpts)
	if err != nil {
		return errgo.Mask(err)
	}

	return events.DisplayTimeline(appEvents, pagination, events.DisplayTimelineOpts{})
}
