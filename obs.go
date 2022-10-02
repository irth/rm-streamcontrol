package main

import (
	"fmt"

	"github.com/andreykaipov/goobs"
	"github.com/andreykaipov/goobs/api/typedefs"
)

type OBS struct {
	Address  string
	Password string
	// TODO: WithPassword bool
	client *goobs.Client

	Scenes []*typedefs.Scene
}

func (o *OBS) Connect() error {
	var err error
	o.client, err = goobs.New(o.Address, goobs.WithPassword(o.Password))
	if err != nil {
		return fmt.Errorf("connecting to OBS: %w", err)
	}

	err = o.Refresh()
	if err != nil {
		return fmt.Errorf("refreshing data: %w", err)
	}

	return nil
}

var NotConnectedErr error = fmt.Errorf("OBS not connected")

func (o *OBS) Refresh() error {
	return o.RefreshScenes()
}

func (o *OBS) RefreshScenes() error {
	if o.client == nil {
		return NotConnectedErr
	}

	resp, err := o.client.Scenes.GetSceneList()
	if err != nil {
		return fmt.Errorf("getting scenes: %w", err)
	}
	o.Scenes = resp.Scenes
	return nil
}
