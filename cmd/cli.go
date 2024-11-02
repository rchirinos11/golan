package cmd

import (
	"fmt"

	"github.com/rchirinos11/golan/wol"
)

func Execute(mac string) {
	wol := wol.WolUtil{}
	packet, err := wol.MakeMagic(mac)
	if err != nil {
		fmt.Println("Error creating magic packet", err)
	}
	if err = wolService.SendMagic(packet); err != nil {
		fmt.Println("Error sending magic packet", err)
	}
	fmt.Println("Sent magic packet to:", mac)
}
