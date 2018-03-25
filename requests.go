package main

import (
	"fmt"

	u "github.com/eleves-ig2i/ig2i-le4-ii-2017-train_controller/unitelway"
	"github.com/eleves-ig2i/ig2i-le4-ii-2017-train_controller/xway"
)

var (
	secretSequence = []byte{0x00, 0x00, 0x00, 0x01, 0x00}

	ErrRequestFailed = fmt.Errorf("report says request failed")
)

func (t transmitter) writeVar(objectType, address uint16, v interface{}, mode bool) error {
	message, _, err := u.WriteObject(objectType, address, v)
	if err != nil {
		return err
	}
	t.requests <- mode
	t.message <- frame{
		b: message,
	}

	response := <-t.message
	if len(response.b) != 1 || response.b[0] != u.WRITE_OBJECT%256 {
		return ErrRequestFailed
	}

	return nil
}

func encodeMODBUS(request []byte) ([]byte, error) {
	if len(request) > 65535 {
		return nil, fmt.Errorf("too much data sent in the MODBUS request")
	}
	lg := len(request) + 1
	request = append([]byte{byte(lg % 256), byte(lg / 256)}, request...)
	return append(secretSequence, request...), nil
}

func newXWAY(station, network, gate byte) xway.XWAYRequest {
	return xway.NewXWAYRequest(station, network, gate, MY_STATION, MY_NETWORK, MY_GATE)
}
