package xway

import "fmt"

const (
	GATE_SYSTEM = iota
	GATE_TERMINAL1
	GATE_TERMINAL2
	GATE_TERMINAL3

	GATE_LEVEL5 = 5
	GATE_LEVEL6 = 8
)
const (
	sender   = 0
	receiver = 1
)

var (
	ErrWrongStationNumber = fmt.Errorf("Wrong station number (should be between 0-63, 255).")
)

type XWAYRequest struct {
	Sender, Receiver Address
	Refused          bool

	Header []byte

	npdu   byte
	params parameters
}

type parameters []parameter
type parameter struct {
	id, lg byte
	value  []byte
}

type Address struct {
	Station, Network, Gate byte
}

func NewXWAYRequest(station, network, gate, myStation, myNetwork, myGate byte) XWAYRequest {
	myXWAY := XWAYRequest{
		Sender: Address{
			myStation,
			myNetwork,
			myGate,
		},
		Receiver: Address{
			station,
			network,
			gate,
		},
	}
	return myXWAY
}
