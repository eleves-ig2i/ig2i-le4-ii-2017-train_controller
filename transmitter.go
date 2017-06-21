package main

import (
	"fmt"
	"net"

	"github.com/kindermoumoute/schneider/xway"
)

const (
	SEND             = false
	SEND_AND_RECEIVE = true
)

var (
	ErrBytesNotSent = fmt.Errorf("failed to write/read all bytes in the request")
)

func initCommunication() transmitter {
	t := newTransmitter()
	go func() {
		err := t.transmit()
		if err != nil {
			panic(err)
		}
	}()
	err := t.setXWAYAddress()
	if err != nil {
		panic(err)
	}
	return t
}

type transmitter struct {
	message  chan frame
	requests chan bool
}

type frame struct {
	b []byte
	x *xway.XWAYRequest
}

func newTransmitter() transmitter {
	return transmitter{
		message:  make(chan frame),
		requests: make(chan bool),
	}
}

func (t transmitter) transmit() error {
	u, err := NewUniteConn()
	if err != nil {
		return err
	}

	for req := range t.requests {
		err := u.write(<-t.message)
		if err != nil {
			return err
		}

		f, err := u.read()
		if err != nil {
			return err
		}
		t.message <- f

		if req == SEND_AND_RECEIVE {
			f, err = u.read()
			if err != nil {
				return err
			}
			t.message <- f
			err = u.write(<-t.message)
			if err != nil {
				return err
			}
		}
	}
	u.c.Close()
	return nil
}

type uniteConn struct {
	x xway.XWAYRequest
	c net.Conn
}

func NewUniteConn() (uniteConn, error) {
	// connect to the automaton
	fmt.Printf("\nDialing machine %s", selectedAutomaton)
	conn, err := net.Dial("tcp", selectedAutomaton)
	if err != nil {
		return uniteConn{}, err
	}

	// write uniteConn object
	u := uniteConn{
		newXWAY(automatonStation, automatonNetwork, automatonGate),
		conn,
	}
	err = u.x.Encode()

	return u, err
}

func (u uniteConn) read() (frame, error) {
	// read MODBUS frame
	buffer := make([]byte, 7)
	n, err := u.c.Read(buffer)
	if err != nil || n != 7 {
		return frame{}, err
	}

	// read XWAY and UNITE frame
	lg := int(buffer[6])*256 + int(buffer[5]) - 1
	response := make([]byte, lg)
	n, err = u.c.Read(response)
	if err != nil {
		return frame{}, err
	}
	if n != lg {
		return frame{}, ErrBytesNotSent
	}
	x, b := xway.Decode(response)

	// printf message
	//fmt.Printf("\nMessage received")
	//util.PrintHex(buffer, response[0:len(response)-len(b)], b)

	return frame{b, x}, nil
}

func (u uniteConn) write(f frame) error {
	var request []byte

	// encapsulate into XWAY and MODBUS
	x := u.x.Header
	if f.x != nil {
		f.x.Encode()
		x = f.x.Header

	}
	request = append(x, f.b...)
	request, err := encodeMODBUS(request)
	if err != nil {
		return err
	}

	// send request
	n, err := u.c.Write(request)
	if err != nil {
		return err
	}
	if n != len(request) {
		return ErrBytesNotSent
	}

	// printf message
	//fmt.Printf("\nMessage sent")
	//util.PrintHex(request[0:7], x, f.b)

	return nil
}
