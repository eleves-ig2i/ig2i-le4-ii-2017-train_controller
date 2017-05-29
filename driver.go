package main

import (
	"fmt"
	"net"

	"github.com/kindermoumoute/schneider/xway"
)

type requestDriver struct {
	sender, receiver chan []byte
}

func runDriver(reqUnite chan requestDriver) error {
	myXWAY := xway.NewXWAYRequest(3, 8, 0)
	err := myXWAY.Encode()
	if err != nil {
		return err
	}

	fmt.Printf("\nDialing machine on port %s", machineAddress)
	conn, err := net.Dial("tcp", machineAddress)
	if err != nil {
		return err
	}
	for req := range reqUnite {
		for message := range req.sender {
			// encapsulate into XWAY and MODBUS
			request := append(myXWAY.Header, message...)
			request, err = encodeMODBUS(request)
			if err != nil {
				return err
			}

			// send request
			n, err := conn.Write(request)
			if err != nil {
				return err
			}
			fmt.Printf("\nMessage sent")
			printHex(request, n)
			buffer := make([]byte, 7)
			n, err = conn.Read(buffer)
			if err != nil || n != 7 {
				return err
			}
			lg := int(buffer[6])*256 + int(buffer[5])
			response := make([]byte, lg)
			n, err = conn.Read(response)
			if err != nil {
				return err
			}
			fmt.Printf("\nMessage received")
			printHex(response, n)
			req.receiver <- response
		}
	}
	return nil
}
