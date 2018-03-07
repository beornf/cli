package scalingo

import (
	"gopkg.in/errgo.v1"
)

type AutoscalersService interface {
	AutoscalersList(app string) ([]Autoscaler, error)
	AutoscalerAdd(app string, email string) (Autoscaler, error)
	AutoscalerRemove(app string, id string) error
}

type AutoscalersClient struct {
	subresourceClient
}

type Autoscaler struct {
	ID            string  `json:"id"`
	ContainerType string  `json:"container_type"`
	Metric        string  `json:"metric"`
	Target        float64 `json:"target"`
	MinContainers int     `json:"min_containers"`
	MaxContainers int     `json:"max_containers"`
	Disabled      bool    `json:"disabled"`
}

type AutoscalersRes struct {
	Autoscalers []Autoscaler `json:"autoscalers"`
}

type AutoscalerRes struct {
	Autoscaler Autoscaler `json:"autoscaler"`
}

func (c *AutoscalersClient) AutoscalersList(app string) ([]Autoscaler, error) {
	var autoscalersRes AutoscalersRes
	err := c.subresourceList(app, "autoscalers", nil, &autoscalersRes)
	if err != nil {
		return nil, errgo.Mask(err)
	}
	return autoscalersRes.Autoscalers, nil
}

type AutoscalerAddParams struct {
	ContainerType string  `json:"container_type"`
	Metric        string  `json:"metric"`
	Target        float64 `json:"target"`
	MinContainers int     `json:"min_containers"`
	MaxContainers int     `json:"max_containers"`
}

func (c *AutoscalersClient) AutoscalerAdd(app string, params AutoscalerAddParams) (*Autoscaler, error) {
	var autoscalerRes AutoscalerRes
	err := c.subresourceAdd(app, "autoscalers", AutoscalerRes{
		Autoscaler: Autoscaler{
			ContainerType: params.ContainerType,
			Metric:        params.Metric,
			Target:        params.Target,
			MinContainers: params.MinContainers,
			MaxContainers: params.MaxContainers,
		},
	}, &autoscalerRes)
	if err != nil {
		return nil, errgo.Mask(err)
	}
	return &autoscalerRes.Autoscaler, nil
}

func (c *AutoscalersClient) AutoscalerShow(app, id string) (*Autoscaler, error) {
	var autoscalerRes AutoscalerRes
	err := c.subresourceGet(app, "autoscalers", id, nil, &autoscalerRes)
	if err != nil {
		return nil, errgo.Mask(err)
	}
	return &autoscalerRes.Autoscaler, nil
}

type AutoscalerUpdateParams struct {
	Metric        string  `json:"metric"`
	Target        float64 `json:"target"`
	MinContainers int     `json:"min_containers"`
	MaxContainers int     `json:"max_containers"`
	Disabled      bool    `json:"disabled"`
}

func (c *AutoscalersClient) AutoscalerUpdate(app, id string, params AutoscalerUpdateParams) (*Autoscaler, error) {
	var autoscalerRes AutoscalerRes
	err := c.subresourceUpdate(app, "autoscalers", id, params, &autoscalerRes)
	if err != nil {
		return nil, errgo.Mask(err)
	}
	return &autoscalerRes.Autoscaler, nil
}

func (c *AutoscalersClient) AutoscalerRemove(app, id string) error {
	err := c.subresourceDelete(app, "autoscalers", id)
	if err != nil {
		return errgo.Mask(err)
	}
	return nil
}
