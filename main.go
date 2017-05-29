package main

import (
	"fmt"

	u "github.com/kindermoumoute/schneider/unitelway"
	"github.com/kindermoumoute/schneider/xway"
)

const (
	machineAddress = "192.168.209.254:502"

	// TODO: write constants
)

func requestExample() ([]byte, error) {
	request, _, err := u.WriteObject(u.InternalWord, 100, []uint16{1, 2, 3})
	return request, err
}

func main() {
	driver := make(chan requestDriver)
	go func() {
		err := runDriver(driver)
		if err != nil {
			panic(err)
		}
	}()

	message, err := requestExample()
	if err != nil {
		panic(err)
	}

	sender := make(chan []byte)
	receiver := make(chan []byte)
	request := requestDriver{
		sender,
		receiver,
	}
	driver <- request
	sender <- message
	response := <-request.receiver
	close(sender)
	close(receiver)

	myXWAY, unite := xway.Decode(response)
	fmt.Printf("\n\n\n")
	printHex(unite, 0)
	fmt.Printf("\n")
	fmt.Printf("XWAY object,%+v", myXWAY)
}

func printHex(b []byte, n int) {
	fmt.Printf("\n")
	fmt.Printf("%d bytes\n", n)
	for _, elt := range b {
		fmt.Printf("%x ", elt)
	}
}
