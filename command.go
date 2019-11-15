package main

import (
	"context"
	"fmt"
	ocf "github.com/go-ocf/sdk/local"
	"github.com/go-ocf/sdk/schema"
	"time"
)

func NewDeviceClient(cfg DeviceConfig) (DeviceClient, error) {

	ocf := ocf.NewClient()
	if ocf == nil {
		panic("unable to create local ocf client")
	}
	c := client{
		client: ocf,
		device: nil,
		config: cfg,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	c.findDevice(ctx)

	return &c, nil
}

type DeviceConfig struct {
	SwitchType string
	DeviceType string
}

type DeviceClient interface {
	SetSwitch(context.Context, bool) (SwitchState, error)
	GetSwitch(context.Context) (SwitchState, error)
}

type client struct {
	client      *ocf.Client
	device      *ocf.Device
	deviceLinks schema.ResourceLinks
	config      DeviceConfig
}

func (c *client) findDevice(ctx context.Context) error {
	err := c.client.GetDevices(ctx, c)
	if err != nil {
		return err
	}
	return nil
}

type SwitchState struct {
	State bool `codec:"state"`
}

func (c *client) SetSwitch(ctx context.Context, state bool) (SwitchState, error) {

	links := c.deviceLinks.GetResourceLinks(c.config.SwitchType)
	if len(links) == 0 {
		return SwitchState{}, fmt.Errorf("unable to get resource link")
	}
	link := links[0]
	sws := SwitchState{State: state}
	err := c.device.UpdateResource(ctx, link, &sws, nil)
	if err != nil {
		return SwitchState{}, err
	}
	return c.GetSwitch(ctx)
}

func (c *client) GetSwitch(ctx context.Context) (SwitchState, error) {

	links := c.deviceLinks.GetResourceLinks(c.config.SwitchType)
	if len(links) == 0 {
		return SwitchState{}, fmt.Errorf("unable to get resource link")
	}
	link := links[0]
	sws := SwitchState{}
	err := c.device.GetResource(ctx, link, &sws)
	if err != nil {
		return SwitchState{}, err
	}
	return sws, nil
}

func (c *client) Handle(ctx context.Context, device *ocf.Device, deviceLinks schema.ResourceLinks) {

	if c.device == nil && SliceContains(device.DeviceTypes(), c.config.DeviceType) {
		c.device = device
		c.deviceLinks = deviceLinks
		return
	}
}


func (c *client) Error(err error) {
	fmt.Errorf("unable to get devices: %s", err)
}

func SliceContains(slices []string, value string) bool {
	for _, v := range slices {
		if v == value {
			return true
		}
	}
	return false
}
