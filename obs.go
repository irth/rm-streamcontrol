package main

import (
	"fmt"

	"github.com/andreykaipov/goobs"
	"github.com/andreykaipov/goobs/api/requests/inputs"
	"github.com/andreykaipov/goobs/api/typedefs"
)

type OBS struct {
	Address  string
	Password string
	// TODO: WithPassword bool
	client *goobs.Client

	Scenes      []*typedefs.Scene
	AudioInputs []*typedefs.Input
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
	if err := o.RefreshScenes(); err != nil {
		return err
	}
	return o.RefreshAudioInputs()
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

func (o *OBS) RefreshAudioInputs() error {
	if o.client == nil {
		return NotConnectedErr
	}

	resp, err := o.client.Inputs.GetInputList()
	if err != nil {
		return fmt.Errorf("getting inputs: %w", err)
	}

	audio := []*typedefs.Input{}
	for _, input := range resp.Inputs {
		_, err := o.client.Inputs.GetInputVolume(&inputs.GetInputVolumeParams{InputName: input.InputName})
		if err != nil {
			continue
		}
		audio = append(audio, input)
	}
	o.AudioInputs = audio
	return nil
}
