package main

import "fmt"

var (
	secretSequence = []byte{0x00, 0x00, 0x00, 0x01, 0x00}
)

func encodeMODBUS(request []byte) ([]byte, error) {
	if len(request) > 65535 {
		return nil, fmt.Errorf("too much data sent in the MODBUS request")
	}
	request = append([]byte{byte(len(request) % 256), byte(len(request) / 256)}, request...)
	return append(secretSequence, request...), nil
}
