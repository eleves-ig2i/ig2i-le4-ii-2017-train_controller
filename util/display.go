package util

import "fmt"

// u: MODBUS frame
// x: XWAY frame
// b: UNI-TE frame
func PrintHex(u, x, b []byte) {
	fmt.Printf("\n")
	for _, elt := range u {
		fmt.Printf("%x ", elt)
	}
	fmt.Printf("\t| ")
	for _, elt := range x {
		fmt.Printf("%x ", elt)
	}
	fmt.Printf("\t| ")
	for _, elt := range b {
		fmt.Printf("%x ", elt)
	}
	fmt.Printf("\nMODBUS\t\t| XWAY")
	for i := 0; i < len(x)/2; i++ {
		fmt.Printf("\t")
	}
	fmt.Printf("\t| UNI-TE\n")
}
