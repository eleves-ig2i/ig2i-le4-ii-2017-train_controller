package main

import "github.com/kindermoumoute/schneider/xway"

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

	MY_STATION = 0x20
	MY_NETWORK = 0x08
	MY_GATE    = xway.GATE_SYSTEM
)

func main() {
	t := initCommunication()

	for {
		t.activate(TJ0d)
		t.activate(A0d)
		t.activate(T8)
		t.activate(T12)
		t.activate(TJ1b)
		t.activate(A5b)
		t.activate(T17)
		t.activate(A6b)
		t.activate(TJ2b)
		t.activate(TI5)
		t.activate(PA1d)
		t.activate(T19)
		t.activate(PA0d)
		t.activate(T15)
		t.activate(A2d)
		t.activate(T9)
	}
}
