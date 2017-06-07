package main

import (
	"fmt"

	u "github.com/kindermoumoute/schneider/unitelway"
	"github.com/kindermoumoute/schneider/xway"
)

const (
	// Tronçons
	TI0 = iota
	TI1
	TI2
	TI3
	TI4
	TI5
	TI6
	TI7

	T8
	T9
	T10
	T11
	T12
	T13
	T14
	T15
	T16
	T17
	T18
	T19

	//Aiguillages
	A0b
	A0d
	A1b
	A1d
	A2b
	A2d
	A3b
	A3d
	A4b
	A4d
	A5b
	A5d
	A6b
	A6d
	A7b
	A7d
	A8b
	A8d

	PA0b
	PA0d
	PA1b
	PA1d

	TJ0b
	TJ0d
	TJ1b
	TJ1d
	TJ2b
	TJ2d

	// Inversion de sens
	I0
	I1
	I2
	I3
	I4
	I5
	I6
	I7
)

func (t transmitter) activate(section uint16) error {
	fmt.Println("Section demandée ", section)
	_, _, err := t.writeVar(u.InternalWord, 10, []uint16{T8})
	if err != nil {
		return err
	}
	t.requests <- RECEIVE
	r := <-t.message
	newXWAY := xway.NewXWAYRequest(r.x.Sender.Station, r.x.Sender.Network, r.x.Sender.Gate)
	t.message <- frame{
		b: []byte{0xfe},
		x: &newXWAY,
	}

	return nil
}
