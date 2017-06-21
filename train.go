package main

import (
	"fmt"

	u "github.com/kindermoumoute/schneider/unitelway"
)

type train struct {
	track        []ressource
	position, id int
}

func (s signaller) run(train train) {
	go func() {
		i := 0
		for {
			if i == 10 {
				err := s.t.writeVar(u.InternalBit, 61, []bool{true}, WRITE)
				if err != nil {
					fmt.Printf("=== ERROR === %v", err)
				}
			}
			for _, ressource := range train.track {
				ressource.m.Lock()
				fmt.Printf("\nRessource prise")
				for _, section := range ressource.sections {
					s.request[train.id] <- section
				}

				ressource.m.Unlock()
				fmt.Printf("\nRessource rendue")
			}
		}
	}()
}

func (t transmitter) activate(section uint16) error {
	fmt.Printf("\n\nSection %d", section)
	return t.writeVar(u.InternalWord, 10, []uint16{section}, WRITE)
}

func (t transmitter) setXWAYAddress() error {
	var xwayAddress uint16 = MY_STATION<<8 + MY_NETWORK
	return t.writeVar(u.InternalWord, 300, []uint16{0x0401, xwayAddress, 0xFE00}, WRITE)

}
