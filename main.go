package main

import (
	"fmt"
	"net"

	"github.com/kindermoumoute/schneider/xway"
)

const (
	machineAddress = "192.168.209.254:502"

	// TODO: write constants
)

func requestExample() ([]byte, error) {
	request := []byte{0x68, 0x07, 0x64, 0x00, 0x03, 0x00, 0x01, 0x00, 0x02, 0x00, 0x03, 0x00}
	request = encodeUNITE(0x37, 0x06, request)

	myXWAY := xway.NewXWAYRequest(3, 8, 0)
	err := myXWAY.Encode()
	if err != nil {
		return nil, err
	}
	request = append(myXWAY.Header, request...)

	request, err = encodeMODBUS(request)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func main() {
	message, err := requestExample()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nDialing machine on port %s", machineAddress)
	conn, err := net.Dial("tcp", machineAddress)
	if err != nil {
		panic(err)
	}
	n, err := conn.Write(message)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nMessage sent")
	printHex(message, n)

	buffer := make([]byte, 7)
	n, err = conn.Read(buffer)
	if err != nil || n != 7 {
		panic(err)
	}
	lg := int(buffer[6])*256 + int(buffer[5])
	response := make([]byte, lg)
	n, err = conn.Read(response)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nMessage received")
	printHex(response, n)

	myXWAY, unite := xway.Decode(response)
	fmt.Println(myXWAY)
	printHex(unite, 0)
}

func printHex(b []byte, n int) {
	fmt.Printf("\n")
	fmt.Printf("%d bytes\n", n)
	for _, elt := range b {
		fmt.Printf("%x ", elt)
	}
}
