package main

import (
	"fmt"

	"sync"

	u "github.com/kindermoumoute/schneider/unitelway"
)

type signaller struct {
	sync.Mutex
	request chan sectionRequest
	t       transmitter

	trains map[int]uint16
}

type sectionRequest struct {
	section uint16
	trainID int
}

func newSignaller(t transmitter) signaller {
	return signaller{
		request: make(chan sectionRequest),
		t:       t,
		trains:  make(map[int]uint16),
	}
}

func startSignaller(t transmitter) signaller {
	s := newSignaller(t)
	go func() {
		req := <-s.request
		s.Lock()
		_, exist := s.trains[req.trainID]
		if !exist {
			s.trains[req.trainID] = req.section
		}
		if
		s.Unlock()
	}()
	go func() {
		sectionRequests := make(map[int]uint16)
		req1 := <-s.request[0]
		req2 := <-s.request[1]

		s.t.writeVar(u.InternalWord, 10, []uint16{req1}, WRITE)
		s.t.writeVar(u.InternalWord, 10, []uint16{req2}, WRITE)
		var tmp uint16

		for {
			t.requests <- READ_AND_WRITE
			r := <-t.message
			newXWAY := newXWAY(r.x.Sender.Station, r.x.Sender.Network, r.x.Sender.Gate)
			t.message <- frame{
				b: []byte{0xfe},
				x: &newXWAY,
			}

			cr := uint16(r.b[len(r.b)-2] + r.b[len(r.b)-1]*256)

			if cr&8 > 0 {
				fmt.Printf("=== WARNING === Train trop rapide")
			} else if cr&9 > 0 {
				fmt.Printf("=== WARNING === Train trop lent")
			}

			if cr&255 == req1 {
				req1 = <-s.request[0]
				tmp = req1
			} else if cr&255 == req2 {
				req2 = <-s.request[1]
				tmp = req2
			} else {
				panic("request is not supposed to be unknown")
			}
			message, _, err := u.WriteObject(u.InternalWord, 10, []uint16{tmp})
			if err != nil {
				panic(err)
			}
			t.message <- frame{
				b: message,
			}
			response := <-t.message
			if len(response.b) != 1 || response.b[0] != u.WRITE_OBJECT%256 {
				panic(ErrRequestFailed)
			}
		}

	}()
	return s
}
