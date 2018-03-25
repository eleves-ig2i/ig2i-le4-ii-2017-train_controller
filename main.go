package main

import "github.com/eleves-ig2i/ig2i-le4-ii-2017-train_controller/xway"

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
	train1, train2 := tracks()
	t := initCommunication()

	t.run(train1)
	t.run(train2)

	select {}
}
