package main

import (
	"fmt"

	"github.com/kindermoumoute/schneider/xway"
)

const (
	automaton2    = "192.168.209.252:502"
	automaton3    = "192.168.209.253:502"
	mainAutomaton = "192.168.209.254:502"

	selectedAutomaton = mainAutomaton

	automaton2XWAY    = "8.1"
	automaton3XWAY    = "8.2"
	mainAutomatonXWAY = "8.3"

	automatonStation = 3
	automatonNetwork = 8
	automatonGate    = 0

	MY_STATION = 0x21
	MY_NETWORK = 0x08
	MY_GATE    = xway.GATE_SYSTEM
)

func main() {
	t := initCommunication()

	track := []uint16{TJ0d, A0d, T8, T12, TJ1b, A5b, T17, A6b, TJ2b, TI5, PA1d, T19, PA0d, T15, A2d, T9}
	for {
		for _, section := range track {
			err := t.activate(section)
			if err != nil {
				fmt.Printf("=== ERROR === %v", err)
			}
		}
	}
}
