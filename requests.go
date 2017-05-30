package main

import (
	u "github.com/kindermoumoute/schneider/unitelway"
	"github.com/kindermoumoute/schneider/xway"
)

func (t transmitter) writeVar(objectType, address uint16, v interface{}) (*xway.XWAYRequest, []byte, error) {
	message, _, err := u.WriteObject(objectType, address, v)
	if err != nil {
		return nil, nil, err
	}
	request := transmitterChannels{
		sender:   make(chan frame),
		receiver: make(chan frame),
		mode:     SEND,
	}
	t <- request
	request.sender <- frame{b:message}
	response := <-request.receiver

	return response.x, response.b, nil
}

//func (t transmitter) readVar(objectType, address uint16)  {
//
//}
