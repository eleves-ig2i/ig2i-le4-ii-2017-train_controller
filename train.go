package main

import (
	"fmt"

	u "github.com/kindermoumoute/schneider/unitelway"
)

type train struct {
	track    []ressource
	position int
}

func (t transmitter) run(train train) {
	go func() {
		i := 0
		for {
			if i == 10 {
				err := t.writeVar(u.InternalBit, 61, []bool{true}, SEND)
				if err != nil {
					fmt.Printf("=== ERROR === %v", err)
				}
			}
			for _, ressource := range train.track {
				ressource.m.Lock()
				for _, section := range ressource.sections {
					err := t.activate(section)
					if err != nil {
						fmt.Printf("=== ERROR === %v", err)
					}
				}

				ressource.m.Lock()
			}
		}
	}()
}

func (t transmitter) activate(section uint16) error {
	fmt.Printf("\n\nSection %d", section)
	err := t.writeVar(u.InternalWord, 10, []uint16{section}, SEND_AND_RECEIVE)
	if err != nil {
		return err
	}
	r := <-t.message
	newXWAY := newXWAY(r.x.Sender.Station, r.x.Sender.Network, r.x.Sender.Gate)
	t.message <- frame{
		b: []byte{0xfe},
		x: &newXWAY,
	}
	return nil
}

func (t transmitter) setXWAYAddress() error {
	var xwayAddress uint16 = MY_STATION<<8 + MY_NETWORK
	return t.writeVar(u.InternalWord, 300, []uint16{0x0401, xwayAddress, 0xFE00}, SEND)

}
