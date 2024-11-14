package cmd

import (
	"fmt"

	"github.com/rchirinos11/golan/wol"
)

func Execute(mac string) error {
	wol := wol.WolUtil{}
	packet, err := wol.MakeMagic(mac)
	if err != nil {
		fmt.Println("Error creating magic packet", err)
		return err
	}
	if err = wolService.SendMagic(packet); err != nil {
		fmt.Println("Error sending magic packet", err)
		return err
	}
	fmt.Println("Sent magic packet to:", mac)
	return nil
}
