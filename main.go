package main

import (
	"context"
	"flag"
	"log"
	"time"
)

const(
	On = "on"
	Off = "off"
)


func main()  {

	var on bool
	var off bool
	flag.BoolVar(&on, On, false, "turn on switch")
	flag.BoolVar(&off, Off, false, "turn off switch")
	flag.Parse()

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

	if on || off {
		var action bool
		if on {
			action = true
		} else if off {
			action = false
		}
		log.Println("Set value on switch", action)
		state, err := client.SetSwitch(ctx, action)
		if err != nil {
			log.Fatal("unable to get switch")
		}
		log.Printf("Switch enable: %v\n", state.State)
	} else {
		log.Println("Get value from switch")
		state, err := client.GetSwitch(ctx)
		if err != nil {
			log.Fatal("unable to get switch")
		}
		log.Printf("Switch enable: %v\n", state.State)
	}
}


