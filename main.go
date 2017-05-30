package main

import (
	"time"

	u "github.com/kindermoumoute/schneider/unitelway"
)

const (
	automaton2    = "192.168.209.252:502"
	automaton3    = "192.168.209.253:502"
	mainAutomaton = "192.168.209.254:502"

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
	_, _, err := t.writeVar(u.InternalWord, 500, []uint16{0x0401, 0x3008, 0xFE00})
	if err != nil {
		panic(err)
	}

	time.Sleep(100 * time.Millisecond)
	//_, _, err = t.writeVar(u.InternalWord, 100, []uint16{36})
	//if err != nil {
	//	panic(err)
	//}
	//sender := make(chan frame)
	//receiver := make(chan frame)
	//request := transmitterChannels{
	//	sender:   sender,
	//	receiver: receiver,
	//	mode:     RECEIVE,
	//}
	//t <- request
	//<-request.receiver
	//request.sender <- frame{b:[]byte{0xfe}}
}
