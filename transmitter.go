package main

import (
	"fmt"
	"net"

	"github.com/kindermoumoute/schneider/util"
	"github.com/kindermoumoute/schneider/xway"
)

const (
	SEND    = false
	RECEIVE = true
)

type transmitter chan transmitterChannels

type transmitterChannels struct {
	sender, receiver chan frame
	mode             bool
}

type frame struct {
	b []byte
	x *xway.XWAYRequest
}

type uniteConn struct {
	x xway.XWAYRequest
	c net.Conn
}

func transmit(reqUnite transmitter) error {
	fmt.Printf("\nDialing machine on port %s", automaton2)
	conn, err := net.Dial("tcp", automaton2)
	if err != nil {
		return err
	}

	u := uniteConn{
		xway.NewXWAYRequest(automatonStation, automatonNetwork, automatonGate),
		conn,
	}
	// write XWAY header
	err = u.x.Encode()
	if err != nil {
		return err
	}

	for req := range reqUnite {
		if req.mode == SEND {
			message := <-req.sender
			if message.x != nil {
				u.x = *message.x
			}
			_, err := u.write(message.b)
			if err != nil {
				return err
			}

			x, b, err := u.read()
			if err != nil {
				return err
			}
			req.receiver <- frame{b, x}

		} else {
			x, b, err := u.read()
			if err != nil {
				return err
			}
			req.receiver <- frame{b, x}
			response := <-req.sender
			if response.x != nil {
				u.x = *response.x
			}
			_, err = u.write(response.b)
			if err != nil {
				return err
			}
		}
	}
	conn.Close()
	return nil
}

func (u uniteConn) read() (*xway.XWAYRequest, []byte, error) {
	fmt.Printf("\nMessage received")
	buffer := make([]byte, 7)
	n, err := u.c.Read(buffer)
	if err != nil || n != 7 {
		return nil, nil, err
	}
	lg := int(buffer[6])*256 + int(buffer[5])
	response := make([]byte, lg)
	_, err = u.c.Read(response)
	if err != nil {
		return nil, nil, err
	}
	x, b := xway.Decode(response)
	util.PrintHex(buffer, response[0:len(response)-len(b)], b)
	return x, b, nil
}
func (u uniteConn) write(message []byte) (int, error) {
	// encapsulate into XWAY and MODBUS
	request := append(u.x.Header, message...)
	request, err := encodeMODBUS(request)
	if err != nil {
		return 0, err
	}

	fmt.Printf("\nMessage sent")
	util.PrintHex(request[0:6], u.x.Header, message)
	// send request
	return u.c.Write(request)
}
