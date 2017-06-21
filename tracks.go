package main

import "sync"

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

type ressource struct {
	m        *sync.Mutex
	sections []uint16
}

func tracks() (train, train) {
	var ressource1 = &sync.Mutex{}
	return train{track: []ressource{
			ressource{
				sections: []uint16{TJ0b, A0d, T8},
				m:        ressource1,
			},
			ressource{
				sections: []uint16{T12, TJ1b, A5b, T17, A6b, TJ2b, TI5, PA1d, T19, PA0d, T15, A2d, T9},
				m:        &sync.Mutex{},
			},
		}},
		train{track: []ressource{
			ressource{
				sections: []uint16{TJ0d, A0b, TI1},
				m:        ressource1,
			},
			ressource{
				sections: []uint16{T11, A7d, T16, A8d, TI7, T14, A1d, TI2, I1, A1b, 101, I0, A1b, TI0},
				m:        &sync.Mutex{},
			},
		}}
}
