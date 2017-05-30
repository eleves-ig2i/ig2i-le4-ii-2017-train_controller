package main

import (
	u "github.com/kindermoumoute/schneider/unitelway"
	"github.com/kindermoumoute/schneider/xway"
)

const (
	automaton2    = "192.168.209.252:502"
	automaton3    = "192.168.209.253:502"
	mainAutomaton = "192.168.209.254:502"

	selectedAutomaton = automaton2

	automaton2XWAY    = "8.1"
	automaton3XWAY    = "8.2"
	mainAutomatonXWAY = "8.3"

	automatonStation = 1
	automatonNetwork = 8
	automatonGate    = 0
)

func main() {
	t := make(transmitter)
	go func() {
		err := transmit(t)
		if err != nil {
			panic(err)
		}
	}()
	_, _, err := t.writeVar(u.InternalWord, 500, []uint16{0x0401, 0x2008, 0xFE00})
	if err != nil {
		panic(err)
	}

	_, _, err = t.writeVar(u.InternalWord, 200, []uint16{36})
	if err != nil {
		panic(err)
	}
	sender := make(chan frame)
	receiver := make(chan frame)
	request := transmitterChannels{
		sender:   sender,
		receiver: receiver,
		mode:     RECEIVE,
	}
	t <- request
	r := <-request.receiver

	newXWAY := xway.NewXWAYRequest(r.x.Sender.Station, r.x.Sender.Network, r.x.Sender.Gate)
	request.sender <- frame{
		b: []byte{0xfe},
		x: &newXWAY,
	}
	//time.Sleep(1000 * time.Millisecond)
}
