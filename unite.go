package main

func encodeUNITE(req, cat byte, request []byte) []byte {
	return append([]byte{req, cat}, request...)
}