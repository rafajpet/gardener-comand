package main

import (
	"context"
	"log"
	"time"
)

func main()  {

	cfg := DeviceConfig{
		SwitchType: "core.switch.1",
		DeviceType: "oic.d.switch.device",
	}
	client, err := NewDeviceClient(cfg)
	if err != nil {
		log.Fatal("unable to create client", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	state, err := client.GetSwitch(ctx)
	if err != nil {
		log.Fatal("unable to get switch")
	}
	log.Println("State", state)

	state, err = client.SetSwitch(ctx, true)
	if err != nil {
		log.Fatal("unable to get switch")
	}
	log.Println("State", state)

	state, err = client.SetSwitch(ctx, false)
	if err != nil {
		log.Fatal("unable to get switch")
	}
	log.Println("State", state)
}


