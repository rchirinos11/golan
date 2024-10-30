package wol

import (
	"bytes"
	"net"
)

type WolUtil struct{}

// Creates a MagicPacket in the form of a byte slice
func (w *WolUtil) MakeMagic(mac string) ([]byte, error) {
	macByte, err := net.ParseMAC(mac)
	if err != nil {
		return nil, err
	}

	header, err := net.ParseMAC("ff:ff:ff:ff:ff:ff")
	if err != nil {
		return nil, err
	}
	payload := bytes.Repeat(macByte, 16)
	packet := append(header, payload...)
	return packet, err
}

func (w *WolUtil) SendMagic(magic []byte) error {
	conn, err := net.Dial("udp", "255.255.255.255:9")
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write(magic)
	return err
}

func (w *WolUtil) GetMacAddr() (mac string, err error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return
	}
	for _, i := range interfaces {
		mac = i.HardwareAddr.String()
		if mac != "" {
			break
		}
	}
	return
}
